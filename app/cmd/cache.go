package cmd

import (
	"IMfourm-go/pkg/cache"
	"IMfourm-go/pkg/console"
	"fmt"
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

// cache forget 命令，清空特定 key 的缓存。
var CmdCacheForget = &cobra.Command{
	Use: "forget",
	Short: "delete redis key, example: cache forget cache-key",
	Run: runCacheForget,
}
var cacheKey string

func init()  {
	//注册cache命令的子命令
	CmdCache.AddCommand(CmdCacheClear,CmdCacheForget)
	//设置cache forget命令的选项
	CmdCacheForget.Flags().StringVarP(&cacheKey,"key","k","","KEY of the cache")
	CmdCacheForget.MarkFlagRequired("key")
}

func runCacheForget(cmd *cobra.Command, args []string)  {
	cache.Forget(cacheKey)
	console.Success(fmt.Sprintf("cache key [%s] deleted.",cacheKey))

}