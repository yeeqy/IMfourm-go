package middlewares

import (
	"IMfourm-go/pkg/jwt"
	"IMfourm-go/pkg/response"
	"github.com/gin-gonic/gin"
)

//Guest 中间件的用法与 Auth 中间件相反。
//用在一些游客身份才能操作的接口上，例如说用户注册、登录接口等。
//使用方法同 Auth 中间件一样，在路由中做绑定，哪些路由只有游客才能访问，直接为其添加中间件即可。

//强制使用游客身份访问
func GuestJWT() gin.HandlerFunc{
	return func(c *gin.Context) {
		if len(c.GetHeader("Authorization")) > 0 {
			//解析token成功，说明登录成功
			_,err := jwt.NewJWT().ParserToken(c)
			if err == nil {
				response.Unauthorized(c,"请使用游客身份访问")
				c.Abort()
				return
			}
		}
		c.Next()
	}
}