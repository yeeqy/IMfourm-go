package captcha

import (
	"IMfourm-go/pkg/app"
	"IMfourm-go/pkg/config"
	"IMfourm-go/pkg/redis"
	"github.com/mojocn/base64Captcha"
	"sync"
)

type Captcha struct {
	Base64Captcha *base64Captcha.Captcha
}

//确保internalCaptcha对象只初始化一次
var once sync.Once
//内部使用的Captcha对象
var internalCaptcha *Captcha

//单例模式获取
func NewCaptcha() *Captcha{
	once.Do(func() {
		//初始化captcha对象
		internalCaptcha = &Captcha{}

		//使用全局Redis对象， 并配置存储key的前缀
		store := RedisStore{
			RedisClient: redis.Redis,
			KeyPrefix: config.GetString("app.name")+":captcha",
		}
		//配置其驱动信息
		driver := base64Captcha.NewDriverDigit(
			config.GetInt("captcha.height"),
			config.GetInt("captcha.width"),
			config.GetInt("captcha.length"),
			config.GetFloat64("captcha.maxskew"),//数字的最大倾角度
			config.GetInt("captcha.dotcount"),//图片背景的混淆点数量
			)
		//实例化，并赋值给内部使用的对象
		internalCaptcha.Base64Captcha = base64Captcha.NewCaptcha(driver,&store)
	})
	return internalCaptcha
}

//生成图片验证码
func (c *Captcha) GenerateCaptcha()(id string,b64s string,err error){
	return c.Base64Captcha.Generate()
}
//验证验证码是否正确
func (c *Captcha) VerifyCaptcha(id string,answer string)(match bool)  {
	//方便本地和API自动测试
	if !app.IsProduction() && id == config.GetString("captcha.testing_key"){
		return true
	}
	//第三个参数是验证后是否删除，我们选择false
	//便于用户多次提交
	return c.Base64Captcha.Verify(id,answer,false)
}