package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/yejingxuan/accumulate/application"
	"github.com/yejingxuan/accumulate/infrastructure/config"
	"github.com/yejingxuan/accumulate/interface/handler"
	"github.com/yejingxuan/accumulate/interface/middleware"
)

//启动http-server
func SetupHttpServer(stockApp application.StockAppInterface) error {
	engine, err := initHttpServer(stockApp)
	if err != nil {
		return err
	}
	if err := engine.Run(fmt.Sprintf(":%d", config.CoreConf.Server.Port)); err != nil {
		return err
	}
	return nil
}

//初始化http接口
func initHttpServer(stockApp application.StockAppInterface) (*gin.Engine, error) {
	// 初始化接口层实例
	stockHandler := handler.NewStockHandler(stockApp)

	//gin.SetMode(gin.ReleaseMode)
	gin.SetMode(gin.DebugMode)
	router := gin.New()
	//跨域支持
	router.Use(middleware.Cors())
	//swagger支持
	router.GET("/accumulate/v1/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//信息获取
	apiInfo := router.Group("/accumulate/v1/info")
	{
		apiInfo.GET("/stock/:code", stockHandler.GetStockInfo) //获取详情
	}

	//任务执行
	apiExec := router.Group("/accumulate/v1/exec")
	{
		apiExec.POST("/update/all", stockHandler.UpdateAll) //更新全部数据
	}

	return router, nil
}
