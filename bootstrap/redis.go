package bootstrap

import (
	"IMfourm-go/pkg/config"
	"IMfourm-go/pkg/redis"
	"fmt"
)

//redis使用配置信息建立连接

func SetupRedis(){
	//建立redis连接
	redis.ConnectRedis(
		fmt.Sprintf("%v:%v", config.GetString("redis.host"), config.GetString("redis.port")),
		config.GetString("redis.username"),
		config.GetString("redis.password"),
		config.GetInt("redis.database"),
		)
}
