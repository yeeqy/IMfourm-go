package cmd

import (
	"github.com/spf13/cobra"
)

//创建 play ，方便我们临时调试代码
//之前我们都在 main.go 里调试，有时候代码会忘记删除
//在 play 命令里测试代码，忘记删除了也不用担心影响到主程序。

var CmdPlay = &cobra.Command{
	Use: "play",
	Short: "Like the Go playground, but running at our application context",
	Run: runPlay,
}
//调试完成后请记得清楚测试代码
func runPlay(cmd *cobra.Command, args []string){
	//存进redis中
	//redis.Redis.Set("hello","hi from redis",10*time.Second)
	//从redis里取出
	//console.Success(redis.Redis.Get("hello"))
}