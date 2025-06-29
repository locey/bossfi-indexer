package core

import (
	appRouter "bossfi-indexer/src/app/router"
	"bossfi-indexer/src/core/chainclient"
	"bossfi-indexer/src/core/config"
	"bossfi-indexer/src/core/ctx"
	"bossfi-indexer/src/core/db"
	"bossfi-indexer/src/core/gin/router"
	"bossfi-indexer/src/core/log"
	"fmt"
	"go.uber.org/zap"
	"net/http"
	_ "net/http/pprof"
)

func Start(configFile string) {
	// 初始化配置信息
	initConfig(configFile)
	// 初始化日志组件
	initLog()
	// 启用性能监控组件
	initPprof()
	// 初始化数据库/Redis
	initDB()
	// 初始化区块链客户端
	initChainClient()
	// 初始化Gin
	initGin()
}

func initConfig(configFile string) {
	ctx.Ctx.Config = config.InitConfig(configFile)
}

func initPprof() {
	if !config.Conf.Monitor.PprofEnable {
		return
	}
	log.Logger.Info("init pprof")
	go func() {
		err := http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", config.Conf.Monitor.PprofPort), nil)
		if err != nil {
			log.Logger.Error("init pprof error", zap.Error(err))
			return
		}
	}()

}

func initLog() {
	ctx.Ctx.Log = log.InitLog()
}

func initDB() {
	ctx.Ctx.DB = db.InitPgsql()
	ctx.Ctx.Redis = db.InitRedis()
}

func initChainClient() {
	chainMap := make(map[int]*chainclient.ChainClient)
	for _, chain := range config.Conf.Chains {
		client, err := chainclient.New(chain.ChainId, chain.Endpoint)
		if err != nil {
			log.Logger.Error("init chain client error", zap.Error(err))
			panic(err)
		}

		chainMap[chain.ChainId] = &client
	}

	ctx.Ctx.ChainMap = chainMap
}

func initGin() {
	r := router.InitRouter()
	ctx.Ctx.Gin = r
	appRouter.Bind(r, &ctx.Ctx)
	err := r.Run(":" + ctx.Ctx.Config.App.Port)
	if err != nil {
		panic(err)
	}
}
