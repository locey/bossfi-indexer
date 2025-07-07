package core

import (
	appRouter "bossfi-indexer/src/app/router"
	"bossfi-indexer/src/app/sync"
	"bossfi-indexer/src/core/chainclient"
	"bossfi-indexer/src/core/config"
	"bossfi-indexer/src/core/ctx"
	"bossfi-indexer/src/core/db"
	"bossfi-indexer/src/core/gin/router"
	"bossfi-indexer/src/core/log"
	"bossfi-indexer/src/core/mq"
	"context"
	"fmt"
	"go.uber.org/zap"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Start(configFile string) {
	c, cancel := context.WithCancel(context.Background())
	defer cancel()

	// 初始化配置信息
	initConfig(configFile)
	// 初始化日志组件
	initLog()
	// 启用性能监控组件
	initPprof()
	// 初始化数据库/Redis
	initDB()
	// 初始化MQ
	initMq()
	// 初始化区块链客户端
	initChainClient()
	// 启动同步服务
	initSync(c)
	// 初始化Gin
	initGin()
	// 监听 kill 信号
	gracefulShutdown(cancel)
	// 阻塞主 goroutine
	select {}
}

func initMq() {
	mq.InitKafka()
	// 启动Kafka消费者
	//go consumer.StartKafkaConsumer()
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
		client, err := chainclient.New(chain)
		if err != nil {
			log.Logger.Error("init chain client error", zap.Error(err))
			panic(err)
		}

		chainMap[chain.ChainId] = &client
	}

	ctx.Ctx.ChainMap = chainMap
}

func initSync(c context.Context) {
	sync.StartSync(c)
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

func gracefulShutdown(cancel context.CancelFunc) {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	<-sigChan
	log.Logger.Info("Received shutdown signal, shutting down gracefully...")

	// 关闭Kafka连接
	mq.CloseKafka()

	// 触发 context cancel，通知所有依赖该 context 的后台协程退出
	cancel()

	// 这里可以加一些等待 DB/Redis 关闭的逻辑
	time.Sleep(2 * time.Second)
	log.Logger.Info("Shutdown complete.")
	os.Exit(0)
}
