//将整个程序改造为命令行模式
package cmd

import (
	"IMfourm-go/bootstrap"
	"IMfourm-go/pkg/config"
	"IMfourm-go/pkg/console"
	"IMfourm-go/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

//serve命令运行我们的web服务

//1. 创建serve命令

var CmdServe = &cobra.Command{
	Use:   "serve",
	Short: "Start web server",
	Run:   runWeb,
	Args:  cobra.NoArgs,
}

func runWeb(cmd *cobra.Command, args []string) {
	// 设置 gin 的运行模式，支持 debug, release, test
	// release 会屏蔽调试信息，官方建议生产环境中使用
	// 非 release 模式 gin 终端打印太多信息，干扰到我们程序中的 Log
	// 故此设置为 release，有特殊情况手动改为 debug 即可
	gin.SetMode(gin.ReleaseMode)

	//gin实例
	router := gin.New()
	//初始化路由绑定
	bootstrap.SetupRoute(router)
	//运行服务器
	err := router.Run(":" + config.Get("app.port"))
	if err != nil {
		logger.ErrorString("CMD","serve",err.Error())
		console.Exit("Unable to start server,error:" + err.Error())
	}

}
