package config

import "IMfourm-go/pkg/config"

func init(){
	config.Add("verifycode", func() map[string]interface{} {
		return map[string]interface{}{
			//1. 验证码长度
			"code_length": config.Env("VERIFY_CODE_LENGTH",6),
			//2. 过期时间，单位是分钟
			"expire_time": config.Env("VERIFY_CODE_EXPIRE",15),
			//3. debug模式下的过期时间，便于本地开发调试
			"debug_expire_time":10080,
			"debug_code":123456,
			//4. 方便本地和API自动测试
			"debug_phone_prefix":"000",
			"debug_email_suffix":"@testing.com",
		}
	})
}
