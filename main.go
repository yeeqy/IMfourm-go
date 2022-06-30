package main

import (
	"IMfourm-go/app/cmd"
	"IMfourm-go/bootstrap"
	btsConfig "IMfourm-go/config"
	"IMfourm-go/pkg/config"
	"IMfourm-go/pkg/console"
	//"flag"
	"fmt"
	//"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"os"
)

func init(){

	//加载config目录一下的配置信息
	btsConfig.Initialize()
}

func main(){

	//2. 创建root命令

	//程序主入口，默认调用cmd.CmdServe命令
	var rootCmd = &cobra.Command{
		Use: config.Get("app.name"),
		Short: "A simple forum project",
		Long: `Default will run "serve" command,you can use "-h" to see all subcommands`,
		//rootCmd所有子命令都会执行一下代码
		PersistentPreRun: func(command *cobra.Command, args []string) {
			//配置初始化，以来命令函 --env参数
			config.InitConfig(cmd.Env)
			//初始化Logger
			bootstrap.SetupLogger()
			//初始化DB
			bootstrap.SetupDB()
			//初始化redis
			bootstrap.SetupRedis()

		},
	}
	//注册子命令
	//后续我们的命令都需要到 main 里的 root 命令里注册。
	rootCmd.AddCommand(
		cmd.CmdServe,
		cmd.CmdKey,
		cmd.CmdPlay,
		)
	//配置默认运行web服务
	cmd.RegisterDefaultCmd(rootCmd,cmd.CmdServe)
	//注册全局参数，--env
	cmd.RegisterGlobalFlags(rootCmd)
	//执行主命令
	if err := rootCmd.Execute();err !=nil{
		console.Exit(fmt.Sprintf("Failed to run app with %v: %s",os.Args,err.Error()))
	}


	////配置初始化，以来命令行 --env参数
	//var env string
	//flag.StringVar(&env,"env","","加载 .env文件，如 --env=testing 加载是 .env.testing文件")
	//flag.Parse()
	//config.InitConfig(env)
	//
	////初始化 Logger
	//bootstrap.SetupLogger()
	//
	//// 设置 gin 的运行模式，支持 debug, release, test
	//// release 会屏蔽调试信息，官方建议生产环境中使用
	//// 非 release 模式 gin 终端打印太多信息，干扰到我们程序中的 Log
	//// 故此设置为 release，有特殊情况手动改为 debug 即可
	//gin.SetMode(gin.ReleaseMode)
	//
	//router := gin.New()
	//
	////初始化DB
	//bootstrap.SetupDB()
	////初始化redis
	//bootstrap.SetupRedis()
	//
	////初始化路由绑定
	//bootstrap.SetupRoute(router)



	//运行服务
	//err := router.Run(":"+config.Get("app.port"))
	//if err!=nil{
		//错误处理
	//	fmt.Println(err.Error())
	//}
}
