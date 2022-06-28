package middlewares

import (
	"IMfourm-go/pkg/logger"
	"IMfourm-go/pkg/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net"
	"net/http/httputil"
	"os"
	"strings"
	"time"
)

//我们希望当 recovery 时，使用 zap 来记录日志，所以需要创建自定的中间件。

func Recovery() gin.HandlerFunc{
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				//获取用户请求信息
				httpReq,_ := httputil.DumpRequest(c.Request, true)

				//连接终端，客户端中断连接为正常行为，不需要记录堆栈信息
				var brokenPipe bool
				if ne, ok := err.(*net.OpError);ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						errStr := strings.ToLower(se.Error())
						if strings.Contains(errStr,"broken pipie") || strings.Contains(errStr,"connection reset by peer"){
							brokenPipe = true
						}
					}
				}
				//连接终端的情况
				if brokenPipe{
					logger.Error(c.Request.URL.Path,
						zap.Time("time",time.Now()),
						zap.Any("error",err),
						zap.String("request",string(httpReq)),
						)
					c.Error(err.(error))
					c.Abort()
					//连接已断开，无法写状态码
					return
				}
				//如果不是连接断开，就开始记录堆栈信息
				logger.Error("recovery from panic",
					//记录时间、错误信息、请求信息、调用堆栈信息
					zap.Time("time",time.Now()),
					zap.Any("error",err),
					zap.String("request",string(httpReq)),
					zap.Stack("stacktrace"),
					)
				//返回状态码
				//c.AbortWithStatusJSON(http.StatusInternalServerError,gin.H{
				//	"message":"服务器内部错误，请稍后再试",
				//})

				//---正解
				//response.Abort500(c)
				response.CreatedJSON(c,gin.H{
					"message":"服务器内部错误",
					"error":err,
				})
			}
		}()
		c.Next()
	}

}
