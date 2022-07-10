package cmd

import (
	"IMfourm-go/pkg/cache"
	"IMfourm-go/pkg/console"
	"github.com/spf13/cobra"
)

var CmdCache = &cobra.Command{
Use:    "cache",
Short:  "cache management",
}

var CmdCacheClear = &cobra.Command{
	Use: "clear",
	Short: "Clear cache",
	Run: runCacheClear,
}

func init()  {
	//注册cache命令的子命令
	CmdCache.AddCommand(CmdCacheClear)

}
func runCacheClear(cmd *cobra.Command, args []string) {
	//调用之前开发好的cache.Flush即可清空缓存
    cache.Flush()
	console.Success("cache cleared.")
}