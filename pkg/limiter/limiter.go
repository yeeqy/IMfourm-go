// Package limiter 自建的 limiter 包对 ulule/limiter 包进行封装
package limiter

//限流就是控制用户访问接口的频率，例如未授权的接口 Github API 每小时最多 60 个请求（根据 IP），而授权以后的接口限流可以到 1000 个请求。
import (
	"IMfourm-go/pkg/config"
	"IMfourm-go/pkg/logger"
	"IMfourm-go/pkg/redis"
	"github.com/gin-gonic/gin"
	limiterlib "github.com/ulule/limiter/v3"
	sredis "github.com/ulule/limiter/v3/drivers/store/redis"
	"strings"
)

//获取Limitor的Key、IP
func GetKeyIP(c *gin.Context) string{
	return c.ClientIP()
}
//Limitor的Key，路由+IP，针对单个路由做限流
func GetKeyRouteWithIP(c *gin.Context) string{
	return routeToKeyString(c.FullPath()) + c.ClientIP()
}

//检测请求是否超频
func CheckRate(c *gin.Context, key string, formatted string)(limiterlib.Context,error){
	//实例化依赖的Limiter包的limiter.Rate对象
	var context limiterlib.Context
	rate, err := limiterlib.NewRateFromFormatted(formatted)
	if err != nil {
		logger.LogIf(err)
		return context,err
	}
	//初始化存储，使用我们程序里公用的redis.Redis对象
	store, err := sredis.NewStoreWithOptions(redis.Redis.Client,limiterlib.StoreOptions{
		//为limiter设置前缀，保持redis数据整洁
		Prefix: config.GetString("app.name")+":limiter",
	})
	if err != nil {
		logger.LogIf(err)
		return context,err
	}
	//使用上面的初始化的limiter.Rate对象和存储对象
	limiterObj := limiterlib.New(store,rate)
	//获取限流结果
	if c.GetBool("limiter-once"){
		//Peek()取结果，不增加访问次数
		return limiterObj.Peek(c,key)
	}else {
		//确保多个路由组里调用LimitIP进行限流时，只增加一次访问次数
		c.Set("limiter-once",true)
		//Get()取结果，并增加访问次数
		return limiterObj.Get(c,key)
	}
}
//辅助方法，将URL的/改为-
func routeToKeyString(routeName string) string{
	routeName = strings.ReplaceAll(routeName,"/","-")
	routeName = strings.ReplaceAll(routeName,":","-")
	return routeName
}
