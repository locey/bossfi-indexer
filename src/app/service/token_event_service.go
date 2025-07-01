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
	}

	// 根据最早和最后区块高度拉取所有该合约的事件
	logs, err := evmAdapter.GetLogs(big.NewInt(lastNumber), nil, evmAdapter.ChainConfig().BossTokenAddress)
	if err != nil {
		log.Logger.Error("GetLogs failed!", zazap.Error(err))
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
			Balance:    *value, // 会传入 excluded.balance
			ModifyTime: time.Now(),
		}

		err := tx.Clauses(clause.OnConflict{
			Columns: []clause.Column{{Name: "address"}}, // 冲突字段
			DoUpdates: clause.Assignments(map[string]interface{}{
				"balance":     gorm.Expr("balance + EXCLUDED.balance"),
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
		currentBalance := &newRecord.Balance
		for i := 0; i < len(balanceLogsMap[addr]); i++ {
			balanceLogsMap[addr][i].AfterBalance = currentBalance
			currentBalance = new(big.Int).Sub(currentBalance, balanceLogsMap[addr][i].ChangeBalance)
			balanceLogsMap[addr][i].BeforeBalance = currentBalance
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

	for _, event := range confirmBlockList {
		if eventOnChainMap[event.TxHash] == nil {
			rorgBlockEvents = append(rorgBlockEvents, event)
			deleteEventIds = append(deleteEventIds, event.ID)
		} else {
			confirmedEventIds = append(confirmedEventIds, event.ID)
		}
	}

	// 在事务前直接更新掉
	err = s.dao.ConfirmedByIds(confirmedEventIds)
	if err != nil {
		log.Logger.Error("confirmedByIds failed!", zazap.Error(err))
		return
	}

	// 计算出要回滚区块的 账户>金额数量
	rollbackBalanceMap, _ := CalcAddressBalance(rorgBlockEvents)

	// 启用事务
	tx := ctx.Ctx.DB.Begin()

	// 批量删除
	if err := tx.Model(&model.TokenEvent{}).Where("id in (?)", deleteEventIds[:]).UpdateColumn("deleted", true).Error; err != nil {
		log.Logger.Error("deleteByIds failed!", zazap.Error(err))
		tx.Rollback()
		return
	}

	// 回滚余额
	for address, balance := range rollbackBalanceMap {
		if err := tx.Model(&model.UserBalance{}).Where("address = ?", address).UpdateColumn("balance", gorm.Expr("balance - ?", balance)).Error; err != nil {
			log.Logger.Error("rollback balance failed!address:"+address, zazap.Error(err))
			tx.Rollback()
			return
		}
	}

	// 提交事务
	tx.Commit()
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
			var value *big.Int
			// 假设 Value 是最后一个数据字段，32 字节 big.Int
			if len(l.Data) >= 32 {
				value = new(big.Int).SetBytes(l.Data[:32])
			}

			var event = model.TokenEvent{
				BlockNumber: int64(l.BlockNumber),
				BlockHash:   l.BlockHash.Hex(),
				TxHash:      l.TxHash.Hex(),
				LogIndex:    int(l.Index),
				EventType:   model.EventTypeTransfer,
				FromAddress: l.Topics[1].Hex(),
				ToAddress:   l.Topics[2].Hex(),
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
			// 减去转出的金额
			addressBalance[event.FromAddress] = addressBalance[event.FromAddress].Sub(addressBalance[event.FromAddress], event.Amount)
			// 添加转入的金额
			addressBalance[event.ToAddress] = addressBalance[event.ToAddress].Add(addressBalance[event.ToAddress], event.Amount)

			var fromBalanceLog = model.UserBalanceLog{
				Address:       event.FromAddress,
				LogType:       2,             // 支出
				BeforeBalance: big.NewInt(0), // 后面更新余额后根据那个时刻都余额往前推算得出
				ChangeBalance: event.Amount,
				AfterBalance:  big.NewInt(0), // 后面更新余额后根据那个时刻都余额往前推算得出
				TxHash:        event.TxHash,
			}
			var toBalanceLog = model.UserBalanceLog{
				Address:       event.ToAddress,
				LogType:       1,             // 收入
				BeforeBalance: big.NewInt(0), // 后面更新余额后根据那个时刻都余额往前推算得出
				ChangeBalance: event.Amount,
				AfterBalance:  big.NewInt(0), // 后面更新余额后根据那个时刻都余额往前推算得出
				TxHash:        event.TxHash,
			}

			balanceLogMap[event.FromAddress] = append(balanceLogMap[event.FromAddress], &fromBalanceLog)
			balanceLogMap[event.ToAddress] = append(balanceLogMap[event.ToAddress], &toBalanceLog)
		}
	}

	return addressBalance, balanceLogMap
}
