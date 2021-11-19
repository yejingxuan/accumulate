package restful

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"op-register/application"
	"op-register/infrastructure/persistence"
	"op-register/interfaces/handler/http_handler"
	"op-register/interfaces/middleware"
)

//初始化http接口
func InitHttpServer(persisDB *persistence.Repositories) (*gin.Engine, error) {
	// 初始化应用层实例
	operatorApp := application.NewOperatorApp(persisDB.OperatorRepo, persisDB.OperatorTempRepo, persisDB.EnvRepo,
		persisDB.ServiceRepo, persisDB.AdjunctRepo)
	categoryApp := application.NewCategoryApp(persisDB.CategoryRepo)
	tagsApp := application.NewTagsApp(persisDB.TagRepo)
	imageApp := application.NewImageApp(persisDB.ImageRepo)
	logInfoApp := application.NewLogApp(persisDB.LogInfoRepo)

	// 初始化接口层实例
	operatorHandler := http_handler.NewOperatorHandler(operatorApp, categoryApp, tagsApp, logInfoApp)
	uploadHandler := http_handler.NewUploadHandler(operatorApp)
	categoryHandler := http_handler.NewCategoryHandler(categoryApp)
	tagsHandler := http_handler.NewTagsHandler(tagsApp)
	imageHandler := http_handler.NewImageHandler(imageApp, logInfoApp)

	//gin.SetMode(gin.ReleaseMode)
	gin.SetMode(gin.DebugMode)
	router := gin.New()
	//跨域支持
	router.Use(middleware.Cors())
	//swagger支持
	router.GET("/operator/v1/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//文件上传接口
	apiUpload := router.Group("/operator/v1/upload")
	{
		uploadHandler.InitTusServer(apiUpload)
	}

	//算子注册接口
	apiOp := router.Group("/operator/v1/register")
	{
		apiOp.POST("/analysis/:id", operatorHandler.Analysis)               //算子解析
		apiOp.POST("/matchImage", operatorHandler.MatchImage)               //镜像匹配
		apiOp.POST("/makeImage", operatorHandler.MakeImage)                 //算子镜像制作
		apiOp.GET("/makeImageLog/:id", operatorHandler.GetMakeImageLog)     //获取算子镜像制作的日志
		apiOp.GET("/publishLog/:id", operatorHandler.GetPublishLog)         //查看算子发布记录
		apiOp.POST("/checkMd5/:md5", operatorHandler.CheckMd5)              //校验算子是否已经上传过
		apiOp.POST("/adjunct", operatorHandler.Adjunct)                     //上传附件
		apiOp.GET("/download/opDepends", operatorHandler.DownloadOpDepends) //下载算子的依赖
		apiOp.POST("/svr/run", operatorHandler.RunSvr)                      //启动服务
		apiOp.POST("/clear/package", operatorHandler.ClearPackage)          //清理多余算子包
	}
	//算子信息接口
	infoOp := router.Group("/operator/v1/info")
	{
		infoOp.GET("/op/:id", operatorHandler.OperatorInfo)                 //获取算子信息
		infoOp.POST("/op/listByName", operatorHandler.OpListByName)         //通过算子名称获取算子的历史版本信息
		infoOp.DELETE("/op/:id", operatorHandler.OpDelete)                  //删除算子
		infoOp.POST("/op/page", operatorHandler.QueryOp)                    //算子分页查询
		infoOp.POST("/op/tree", operatorHandler.QueryOpTree)                //查询算子树状结构
		infoOp.POST("/op/setDefaultValue", operatorHandler.SetDefaultValue) //设置算子参数的默认值
		infoOp.POST("/op/byName", operatorHandler.OpInfoByName)             //根据算子名称查询最新算子

		infoOp.PUT("/op/checkout", operatorHandler.OpCheckout) //算子测试

		infoOp.PUT("/op/online/:id", operatorHandler.OpOnline)   //算子上线
		infoOp.PUT("/op/offline/:id", operatorHandler.OpOffline) //算子下线

		infoOp.POST("/category/create", categoryHandler.CreateCategory) //创建算子分类
		infoOp.POST("/category/list", categoryHandler.GetAllCategory)   //获取算子全部分类
		infoOp.POST("/tags/list", tagsHandler.GetAllTags)               //获取算子全部标签

		infoOp.POST("/adjunct/list", operatorHandler.GetAdjunctList) //获取算子的附件列表
		infoOp.GET("/adjunct/:id", operatorHandler.GetAdjunct)       //下载附件

		infoOp.POST("/image/page", imageHandler.QueryImage) //镜像分页查询
	}

	return router, nil
}
