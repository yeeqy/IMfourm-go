package bootstrap

import (
	"IMfourm-go/pkg/cache"
	"IMfourm-go/pkg/config"
	"fmt"
)

func SetupCache()  {
	//初始化缓存专用的redis client，使用专属缓存DB
	rds := cache.NewRedisStore(
		fmt.Sprintf("%v:%v",config.GetString("redis.host"),config.GetString("redis.port")),
		config.GetString("redis.username"),
		config.GetString("redis.password"),
		config.GetInt("redis.database_cache"),
		)
	cache.InitWithCacheStore(rds)
}