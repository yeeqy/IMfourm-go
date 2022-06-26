package config

import "IMfourm-go/pkg/config"

func init(){
	config.Add("jwt", func() map[string]interface{} {
		return map[string]interface{}{
			"expire_time":config.Env("JWT_EXPIRE_TIME",20),
			//允许刷新时间 单位分钟，86400为两个月 从Token签名时间算起
			"max_refresh_time":config.Env("JWT_MAX_REFRESH_TIME",86400),
			"debug_expire_time":86400,
		}
	})
}
