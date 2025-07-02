package service

import (
	"bossfi-indexer/src/app/model"
	"bossfi-indexer/src/core/chainclient/evm"
	"bossfi-indexer/src/core/ctx"
	"bossfi-indexer/src/core/log"
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	zazap "go.uber.org/zap"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"math/big"
	"time"
)

type TokenEventService struct {
	dao *model.TokenEventModel
}

// TransferEvent 事件结构
type TransferEvent struct {
	From  common.Address
	To    common.Address
	Value *big.Int
}

// MinterChangedEvent 事件结构
type MinterChangedEvent struct {
	PreviousMinter common.Address
	NewMinter      common.Address
}

func NewTokenEventService() *TokenEventService {
	return &TokenEventService{
		dao: &model.TokenEventModel{},
	}
}

// GetEarlyUnConfirmBlock 查询未确认区块
func (s *TokenEventService) GetEarlyUnConfirmBlock(finalizedNumber uint64) ([]*model.TokenEvent, error) {
	return s.dao.GetEarlyUnConfirmBlock(finalizedNumber)
}

// GetLastBlockNumber 获取最后一条区块高度
func (s *TokenEventService) GetLastBlockNumber() int64 {
	return s.dao.GetLastBlockNumber()
}

// Create 创建记录
func (s *TokenEventService) Create(event *model.TokenEvent) error {
	return s.dao.Create(event)
}

// GetByTxHashAndIndex 查询单条记录
func (s *TokenEventService) GetByTxHashAndIndex(txHash string, index int) (*model.TokenEvent, error) {
	return s.dao.GetByTxHashAndIndex(txHash, index)
}

// ListByBlockNumber 查询区块中的事件
func (s *TokenEventService) ListByBlockNumber(blockNumber int64) ([]*model.TokenEvent, error) {
	return s.dao.ListByBlockNumber(blockNumber)
}

// ListAll 查询所有事件
func (s *TokenEventService) ListAll() ([]*model.TokenEvent, error) {
	return s.dao.ListAll()
}

// Page 分页查询
func (s *TokenEventService) Page(page, pageSize int) ([]*model.TokenEvent, int64, error) {
	return s.dao.Page(page, pageSize)
}

// SyncTokenEvent 同步Token事件
func (s *TokenEventService) SyncTokenEvent(chainID int) {
	evmAdapter := ctx.GetClient(chainID).(*evm.Evm)

	lastNumber := s.GetLastBlockNumber()
	if lastNumber == 0 {
		lastNumber = evmAdapter.ChainConfig().StartBlockNumber
	} else {
		lastNumber += 1
	}

	// 根据最早和最后区块高度拉取所有该合约的事件
	logs, err := evmAdapter.GetLogs(big.NewInt(lastNumber), nil, evmAdapter.ChainConfig().BossTokenAddress)
	if err != nil {
		log.Logger.Error("GetLogs failed!!", zazap.Error(err))
		return
	}
	if len(logs) == 0 {
		log.Logger.Info("GetLogs is empty")
		return
	}

	// 解析日志数据并转成事件列表
	events := ParseTokenEvents(logs)

	// 计算出事件数组中的所有地址的余额变更数量
	balanceMap, balanceLogsMap := CalcAddressBalance(events)

	// 启用事务
	tx := ctx.Ctx.DB.Begin()

	// 批量插入事件表
	if err := tx.Clauses(clause.OnConflict{DoNothing: true}).Create(&events).Error; err != nil {
		log.Logger.Error("Create batch failed!", zazap.Error(err))
		tx.Rollback()
		return
	}

	// 循环更新地址余额
	for addr, value := range balanceMap {
		// 构造一条更新或插入记录
		record := model.UserBalance{
			Address:    addr,
			Balance:    value.String(), // 会传入 excluded.balance
			ModifyTime: time.Now(),
		}

		err := tx.Clauses(clause.OnConflict{
			Columns: []clause.Column{{Name: "address"}}, // 冲突字段
			DoUpdates: clause.Assignments(map[string]interface{}{
				"balance":     gorm.Expr("bii_user_balance.balance + EXCLUDED.balance"),
				"modify_time": time.Now(),
			}),
		}).Create(&record).Error

		if err != nil {
			log.Logger.Error("Create failed!", zazap.Error(err))
			tx.Rollback()
			return
		}

		// 重新获取最新的地址余额记录
		var newRecord model.UserBalance
		tx.Model(&model.UserBalance{}).Where("address = ?", addr).Find(&newRecord)

		if len(balanceLogsMap[addr]) == 0 {
			continue
		}

		// 根据最新余额往回推算历史更新前后余额，作为余额变更日志记录
		currentBalance := new(big.Int)
		if _, ok := currentBalance.SetString(newRecord.Balance, 10); !ok {
			log.Logger.Error("invalid number", zazap.String("balance", newRecord.Balance))
			tx.Rollback()
			return
		}

		for i := 0; i < len(balanceLogsMap[addr]); i++ {
			balanceLogsMap[addr][i].AfterBalance = currentBalance.String()

			changeBalance := new(big.Int)
			if _, ok := changeBalance.SetString(balanceLogsMap[addr][i].ChangeBalance, 10); !ok {
				log.Logger.Error("invalid number", zazap.String("changeBalance", balanceLogsMap[addr][i].ChangeBalance))
				tx.Rollback()
				return
			}
			currentBalance = new(big.Int).Sub(currentBalance, changeBalance)
			balanceLogsMap[addr][i].BeforeBalance = currentBalance.String()
		}
	}

	// 批量插入余额变更日志表
	var allBalanceLogs []*model.UserBalanceLog
	for _, balanceLogs := range balanceLogsMap {
		allBalanceLogs = append(allBalanceLogs, balanceLogs...)
	}
	if err := tx.Create(allBalanceLogs).Error; err != nil {
		log.Logger.Error("Create balance log failed!", zazap.Error(err))
		tx.Rollback()
		return
	}

	tx.Commit()
}

// ConfirmTokenEvent 确认区块事件（包括清理重组区块和回滚余额）
func (s *TokenEventService) ConfirmTokenEvent(chainID int) {
	log.Logger.Info("ConfirmTokenEvent start...")
	evmAdapter := ctx.GetClient(chainID).(*evm.Evm)
	// 获取最终确认区块高度
	block, err := evmAdapter.GetFinalizedBlock()
	if err != nil {
		log.Logger.Error("GetFinalizedBlock failed!", zazap.Error(err))
		return
	}
	finalizedNumber := block.Number
	// 获取数据库中未确认的区块事件
	confirmBlockList, err := s.GetEarlyUnConfirmBlock(finalizedNumber)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Logger.Info("GetEarlyUnConfirmBlock ErrRecordNotFound!")
			return
		}
		log.Logger.Error("GetEarlyUnConfirmBlock failed!", zazap.Error(err))
		return
	}
	if len(confirmBlockList) == 0 {
		log.Logger.Info("No unconfirmed blocks found.")
		return
	}

	// 数据库里最早未确认的区块高度
	beginNumber := big.NewInt(confirmBlockList[0].BlockNumber)
	// 数据库里最后一个未确认的区块高度
	endNumber := big.NewInt(confirmBlockList[len(confirmBlockList)-1].BlockNumber)

	// 根据最早和最后区块高度拉取所有该合约的事件
	logs, err := evmAdapter.GetLogs(beginNumber, endNumber, evmAdapter.ChainConfig().BossTokenAddress)
	if err != nil {
		log.Logger.Error("GetLogs failed!", zazap.Error(err))
		return
	}
	//解析日志数据并转成map key=交易hash
	eventOnChainMap := EventToMap(ParseTokenEvents(logs))

	// 发生重组及重组之后的所有区块事件 及对应的 主键id
	var rorgBlockEvents []*model.TokenEvent
	var deleteEventIds []int64
	// 可以确认的区块事件
	var confirmedEventIds []int64
	// 是否重组
	var rorg bool

	for _, event := range confirmBlockList {
		if rorg || eventOnChainMap[event.TxHash] == nil {
			// 发现区块重组后，后面的所有区块事件都加入待回滚和待删除数组 confirmBlockList是升序的
			rorg = true
			rorgBlockEvents = append(rorgBlockEvents, event)
			deleteEventIds = append(deleteEventIds, event.ID)
		} else {
			confirmedEventIds = append(confirmedEventIds, event.ID)
		}
	}

	// 在开启事务前先把已经确认的区块直接更新掉
	if len(confirmedEventIds) > 0 {
		err = s.dao.ConfirmedByIds(confirmedEventIds)
		if err != nil {
			log.Logger.Error("confirmedByIds failed!", zazap.Error(err))
			return
		}
	}

	// 计算出要回滚区块的 账户>金额数量
	rollbackBalanceMap, _ := CalcAddressBalance(rorgBlockEvents)

	// 启用事务
	tx := ctx.Ctx.DB.Begin()

	// 批量删除（务必只删除当前指定的这些id，防止删除该时刻别的同步线程新同步的id）
	if len(deleteEventIds) > 0 {
		if err := tx.Model(&model.TokenEvent{}).Where("id in (?)", deleteEventIds[:]).UpdateColumn("deleted", true).Error; err != nil {
			log.Logger.Error("deleteByIds failed!", zazap.Error(err))
			tx.Rollback()
			return
		}
	}

	// 回滚余额
	if len(rollbackBalanceMap) > 0 {
		var balanceLogs []*model.UserBalanceLog
		for address, balance := range rollbackBalanceMap {
			if err := tx.Model(&model.UserBalance{}).Where("address = ?", address).UpdateColumn("balance", gorm.Expr("balance - ?", balance)).Error; err != nil {
				log.Logger.Error("rollback balance failed!address:"+address, zazap.Error(err))
				tx.Rollback()
				return
			}
			newUserBalance := model.UserBalance{}
			tx.Where("address = ?", address).First(&newUserBalance)
			currentBalance := new(big.Int)
			if _, ok := currentBalance.SetString(newUserBalance.Balance, 10); !ok {
				log.Logger.Error("rollback balance invalid number", zazap.String("balance", newUserBalance.Balance))
				tx.Rollback()
				return
			}

			var fromBalanceLog = model.UserBalanceLog{
				Address:       address,
				LogType:       3,                                                  // 区块重组回滚
				BeforeBalance: new(big.Int).Add(currentBalance, balance).String(), // 后面更新余额后根据那个时刻都余额往前推算得出
				ChangeBalance: new(big.Int).Mul(balance, big.NewInt(-1)).String(),
				AfterBalance:  newUserBalance.Balance, // 后面更新余额后根据那个时刻都余额往前推算得出
				TxHash:        "",
			}

			balanceLogs = append(balanceLogs, &fromBalanceLog)
		}

		if err := tx.Create(&balanceLogs).Error; err != nil {
			log.Logger.Error("Create rollback balance log failed!", zazap.Error(err))
			tx.Rollback()
		}
	}

	// 提交事务
	tx.Commit()
	log.Logger.Info("ConfirmTokenEvent done")
}

func ParseTokenEvents(logs []types.Log) []*model.TokenEvent {
	var (
		transferSig      = crypto.Keccak256Hash([]byte("Transfer(address,address,uint256)")).Hex()
		minterChangedSig = crypto.Keccak256Hash([]byte("MinterChanged(address,address)")).Hex()
	)
	var events []*model.TokenEvent

	for _, l := range logs {
		topic0 := l.Topics[0].Hex()

		switch topic0 {
		case transferSig:
			var value string
			// 假设 Value 是最后一个数据字段，32 字节 big.Int
			if len(l.Data) >= 32 {
				value = new(big.Int).SetBytes(l.Data[:32]).String()
			}

			var event = model.TokenEvent{
				BlockNumber: int64(l.BlockNumber),
				BlockHash:   l.BlockHash.Hex(),
				TxHash:      l.TxHash.Hex(),
				LogIndex:    int(l.Index),
				EventType:   model.EventTypeTransfer,
				FromAddress: common.HexToAddress(l.Topics[1].Hex()).Hex(),
				ToAddress:   common.HexToAddress(l.Topics[2].Hex()).Hex(),
				Amount:      value,
				Confirmed:   false,
			}

			events = append(events, &event)

		case minterChangedSig:
			//var event MinterChangedEvent
			//event.PreviousMinter = common.HexToAddress(l.Topics[1].Hex())
			//event.NewMinter = common.HexToAddress(l.Topics[2].Hex())
			//events = append(events, event)

		default:
			// 可选：记录未知事件
			continue
		}
	}

	return events
}

func EventToMap(events []*model.TokenEvent) map[string]*model.TokenEvent {
	m := make(map[string]*model.TokenEvent)
	for _, event := range events {
		m[event.TxHash] = event
	}
	return m
}

// CalcAddressBalance 遍历事件列表，计算每个地址本次操作金额数量
func CalcAddressBalance(events []*model.TokenEvent) (map[string]*big.Int, map[string][]*model.UserBalanceLog) {
	addressBalance := make(map[string]*big.Int)
	balanceLogMap := make(map[string][]*model.UserBalanceLog)
	for _, event := range events {
		if event.EventType == model.EventTypeTransfer {
			// 如果 FromAddress 尚未初始化，则设为 0
			fromBalance := addressBalance[event.FromAddress]
			if fromBalance == nil {
				fromBalance = big.NewInt(0)
			}
			toBalance := addressBalance[event.ToAddress]
			if toBalance == nil {
				toBalance = big.NewInt(0)
			}
			amount := new(big.Int)
			amount.SetString(event.Amount, 10)

			// 减去转出的金额
			addressBalance[event.FromAddress] = new(big.Int).Sub(fromBalance, amount)
			// 添加转入的金额
			addressBalance[event.ToAddress] = new(big.Int).Add(toBalance, amount)

			var fromBalanceLog = model.UserBalanceLog{
				Address:       event.FromAddress,
				LogType:       2,                      // 支出
				BeforeBalance: big.NewInt(0).String(), // 后面更新余额后根据那个时刻都余额往前推算得出
				ChangeBalance: event.Amount,
				AfterBalance:  big.NewInt(0).String(), // 后面更新余额后根据那个时刻都余额往前推算得出
				TxHash:        event.TxHash,
			}
			var toBalanceLog = model.UserBalanceLog{
				Address:       event.ToAddress,
				LogType:       1,                      // 收入
				BeforeBalance: big.NewInt(0).String(), // 后面更新余额后根据那个时刻都余额往前推算得出
				ChangeBalance: event.Amount,
				AfterBalance:  big.NewInt(0).String(), // 后面更新余额后根据那个时刻都余额往前推算得出
				TxHash:        event.TxHash,
			}

			balanceLogMap[event.FromAddress] = append(balanceLogMap[event.FromAddress], &fromBalanceLog)
			balanceLogMap[event.ToAddress] = append(balanceLogMap[event.ToAddress], &toBalanceLog)
		}
	}

	return addressBalance, balanceLogMap
}
