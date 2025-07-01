package ctx

import (
	"bossfi-indexer/src/core/chainclient"
	"bossfi-indexer/src/core/config"
	"bossfi-indexer/src/core/db"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var Ctx = Context{}

type Context struct {
	Config   *config.Config
	DB       *gorm.DB
	Redis    *db.RedisClient
	Log      *zap.Logger
	ChainMap map[int]*chainclient.ChainClient
	Gin      *gin.Engine
}

func GetClient(chainId int) chainclient.ChainClient {
	return *Ctx.ChainMap[chainId]
}

func GetEvmClient(chainId int) *ethclient.Client {
	client := *Ctx.ChainMap[chainId]
	return client.Client().(*ethclient.Client)
}
