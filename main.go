package main

import (
	"IMfourm-go/bootstrap"
	btsConfig "IMfourm-go/config"
	"IMfourm-go/pkg/config"
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

	router := gin.New()

	//初始化DB
	bootstrap.SetupDB()

	//初始化路由绑定
	bootstrap.SetupRoute(router)


	//运行服务
	err := router.Run(":"+config.Get("app.port"))
	if err!=nil{
		//错误处理
		fmt.Println(err.Error())
	}
}
