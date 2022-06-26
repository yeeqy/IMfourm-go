package config
//站点配置信息
import "IMfourm-go/pkg/config"

func init(){
	config.Add("app", func() map[string]interface{}{
		return map[string]interface{}{
			"name": config.Env("APP_NAME","IMfourm"),
			"env": config.Env("APP_ENV","production"),
			"debug":config.Env("APP_DEBUG",false),
			//应用服务端口
			"port":config.Env("APP_PORT","3000"),
			"key":config.Env("APP_KEY","33446a9dcf9ea060a0a6532b166da32f304af0de"),
			"ufl":config.Env("APP_URL","http://localhost:3000"),
			//尝试写Asia/Hangzhou然后报错500了
			"timezone":config.Env("TIMEZONE","Asia/Shanghai"),

		}
	})
}
