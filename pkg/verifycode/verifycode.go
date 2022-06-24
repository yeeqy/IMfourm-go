package verifycode

import (
	"IMfourm-go/pkg/app"
	"IMfourm-go/pkg/config"
	"IMfourm-go/pkg/helpers"
	"IMfourm-go/pkg/logger"
	"IMfourm-go/pkg/redis"
	"IMfourm-go/pkg/sms"
	"strings"
	"sync"
)

type VerifyCode struct {
	Store Store
}

var once sync.Once
var internalVerifyCode *VerifyCode

//单例模式获取
func NewVerifyCode() *VerifyCode{
	once.Do(func() {
		internalVerifyCode = &VerifyCode{
			Store: &RedisStore{
				RedisClient: redis.Redis,
				KeyPrefix: config.GetString("app.name")+":verifycode",
			},
		}
	})
	return internalVerifyCode
}
//生成验证码，并放置于Redis中
func (vc *VerifyCode) generateVerifyCode(key string) string{
	//生成随机码(生成六伪随机码）
	code := helpers.RandomNumber(config.GetInt("verifycode.code_length"))
	//便于开发因此本地环境使用固定验证码
	if app.IsLocal(){
		code = config.GetString("verifycode.debug_code")
	}
	logger.DebugJSON("验证码","生成验证码",map[string]string{key:code})

	//将验证码以及key存放到redis，并设置过期时间
	vc.Store.Set(key,code)
	return code

}


//发送短信验证码
func (vc *VerifyCode) SendSMS(phone string) bool {
	//1. 生成验证码
	code := vc.generateVerifyCode(phone)
	//2. 便于本地与API自动测试
	if !app.IsProduction() && strings.HasPrefix(phone,config.GetString("verifycode.debug_phone_prefix")){
		return true
	}
	//3. 发送短信
	return sms.NewSMS().Send(phone,sms.Message{
		Template: config.GetString("sms.aliyun.template_code"),
		Data: map[string]string{"code":code},
	})
}

func(vc *VerifyCode) CheckAnswer(key string, answer string) bool{
	logger.DebugJSON("验证码","检查验证码",map[string]string{key:answer})
	//便于开发，非生产环境，具有特殊前缀的手机号 || 特殊后缀的Email，会直接验证成功
	if !app.IsProduction() &&
		(strings.HasSuffix(key,config.GetString("verifycode.debug_email_suffix")) ||
			strings.HasPrefix(key,config.GetString("verifycode.debug_phone_prefix"))){
		return true
	}
	return vc.Store.Verify(key,answer,false)
}