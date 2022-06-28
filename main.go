package main

import (
	"IMfourm-go/app/http/middlewares"
	"IMfourm-go/bootstrap"
	btsConfig "IMfourm-go/config"
	"IMfourm-go/pkg/auth"
	"IMfourm-go/pkg/config"
	"IMfourm-go/pkg/response"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
)

func init(){

	//加载config目录一下的配置信息
	btsConfig.Initialize()
}

func main(){

	//配置初始化，以来命令行 --env参数
	var env string
	flag.StringVar(&env,"env","","加载 .env文件，如 --env=testing 加载是 .env.testing文件")
	flag.Parse()
	config.InitConfig(env)

	//初始化 Logger
	bootstrap.SetupLogger()

	// 设置 gin 的运行模式，支持 debug, release, test
	// release 会屏蔽调试信息，官方建议生产环境中使用
	// 非 release 模式 gin 终端打印太多信息，干扰到我们程序中的 Log
	// 故此设置为 release，有特殊情况手动改为 debug 即可
	gin.SetMode(gin.ReleaseMode)

	router := gin.New()

	//初始化DB
	bootstrap.SetupDB()
	//初始化redis
	bootstrap.SetupRedis()

	//初始化路由绑定
	bootstrap.SetupRoute(router)

	router.GET("/test_auth",middlewares.AuthJWT(), func(c *gin.Context) {
		userModel := auth.CurrentUser(c)
		response.Data(c,userModel)
	})

	//运行服务
	err := router.Run(":"+config.Get("app.port"))
	if err!=nil{
		//错误处理
		fmt.Println(err.Error())
	}
}
