package cmd

import (
	"IMfourm-go/pkg/console"
	"IMfourm-go/pkg/helpers"
	"github.com/spf13/cobra"
)

//开发 key 命令，生成随机字符串
//可用来设置我们的 APP_KEY 环境变量的值

var CmdKey = &cobra.Command{
	Use: "key",
	Short: "Generate App Key, will print the generated key",
	Run: runKeyGenerate,
	Args: cobra.NoArgs,//不允许传参
}

func runKeyGenerate(cmd *cobra.Command, args []string){
	console.Success("---")
	console.Success("App Key:")
	console.Success(helpers.RandomString(32))
	console.Success("---")
	console.Warning("please go to .env file to change the APP_KEY option")
}