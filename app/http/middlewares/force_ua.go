package middlewares

import (
	"IMfourm-go/pkg/response"
	"errors"
	"github.com/gin-gonic/gin"
)

// ForceUA 中间件，强制请求必须附带User-Agent标头
func ForceUA() gin.HandlerFunc  {
	return func(c *gin.Context) {
		//获取User-Agent标头信息
		if len(c.Request.Header["User-Agent"]) == 0 {
			response.BadRequest(c,errors.New("User-Agent标头未找到"),"请求必须附带User-Agent标头")
			return
			}
		c.Next()
	}
}
