package router

import (
	"bossfi-indexer/src/app/api"
	"bossfi-indexer/src/core/config"
	"bossfi-indexer/src/core/ctx"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Bind(r *gin.Engine, ctx *ctx.Context) {
	// 注册 swagger 路由
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v := r.Group("/api/" + config.Conf.App.Version)

	{
		userBalanceApi := api.NewUserBalanceApi()
		v.GET("/user_balance/:address", userBalanceApi.GetBalanceByAddress)
	}

}
