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
		demoApi := api.NewDemoApi()
		v.GET("/demo/page", demoApi.Page)
		v.POST("/demo/create", demoApi.Create)
		v.GET("/demo/:id", demoApi.GetById)
		v.PUT("/demo/:id", demoApi.Update)
		v.DELETE("/demo/:id", demoApi.Delete)
		v.GET("/demo/list", demoApi.List)
	}

	{
		evmApi := api.NewEvmApi()
		v.GET("/evm/get_block_by_num/:block_num", evmApi.GetBlockByNum)
	}

}
