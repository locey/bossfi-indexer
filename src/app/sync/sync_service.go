package sync

import (
	"bossfi-indexer/src/app/service"
	"bossfi-indexer/src/core/chainclient/evm"
	"bossfi-indexer/src/core/ctx"
	"bossfi-indexer/src/core/log"
	"context"
	zazap "go.uber.org/zap"
	"strconv"
	"time"
)

type Sync struct{}

// StartSync 监听并同步链上事件到本地数据库
func StartSync(c context.Context) {
	mainChainId := ctx.Ctx.Config.App.ChainId
	// 定时同步最终区块
	go FinalizedBlock(c, mainChainId)
	// 定时获取区块事件
	go BlockEvent(c, mainChainId)
	// 定时确认区块事件
	go ConfirmBlock(c, mainChainId)

	//q := ethereum.FilterQuery{
	//	FromBlock: nil,
	//	ToBlock:   nil,
	//	Addresses: nil,
	//	Topics:    nil,
	//}
	//client.FilterLogs(context.Context(), q)
	//return nil
}

// FinalizedBlock 定时同步最终区块
func FinalizedBlock(c context.Context, chainID int) {
	evmAdapter := ctx.GetClient(chainID).(*evm.Evm)
	redis := ctx.Ctx.Redis

	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			block, err := evmAdapter.GetFinalizedBlock()
			if err != nil {
				log.Logger.Error("GetFinalizedBlock failed!", zazap.Error(err))
				return
			}

			err = redis.RedisSet("EVM_BLOCK_FINALIZED", block, 600)
			if err != nil {
				log.Logger.Error("Unmarshal json failed:" + err.Error())
			}
			log.Logger.Info("FinalizedBlock: " + strconv.FormatUint(block.Number, 10))
		case <-c.Done():
			log.Logger.Info("Context canceled, FinalizedBlock exiting...")
			return
		}
	}
}

func BlockEvent(c context.Context, chainID int) {
	tokenEventService := service.NewTokenEventService()

	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			tokenEventService.SyncTokenEvent(chainID)
		case <-c.Done():
			log.Logger.Info("Context canceled, ConfirmBlock exiting...")
			return
		}
	}
}

// ConfirmBlock 定时确认区块
func ConfirmBlock(c context.Context, chainID int) {
	tokenEventService := service.NewTokenEventService()

	ticker := time.NewTicker(3 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			tokenEventService.ConfirmTokenEvent(chainID)
		case <-c.Done():
			log.Logger.Info("Context canceled, ConfirmBlock exiting...")
			return
		}
	}
}
