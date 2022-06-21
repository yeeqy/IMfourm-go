package middlewares

import (
	"IMfourm-go/pkg/helpers"
	"IMfourm-go/pkg/logger"
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"go.uber.org/zap"
	"io/ioutil"
	"time"
)

type responseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (r responseBodyWriter) Writer(b []byte)(int,error){
	r.body.Write(b)
	return r.ResponseWriter.Write(b)
}

func Logger() gin.HandlerFunc{
	return func(c *gin.Context) {

		//获取response内容
		w := &responseBodyWriter{
			body: &bytes.Buffer{},
			ResponseWriter: c.Writer,
		}
		c.Writer = w

		//获取请求数据也不要“考试不考就不学”，多做一点“无用”的事，看一些“无用”的书，培养一些“无用”的爱好与能力。”谢谢Soren老师😭后天要考三笔原本很焦虑，看完就释怀了
		var requestBody []byte
		if c.Request.Body != nil {
			//c.Request.Body是一个buffer对象，只能读取一次
			requestBody,_ = ioutil.ReadAll(c.Request.Body)
			//读取后重新赋值
			c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(requestBody))
		}
		//设置开始时间
		start := time.Now()
		c.Next()

		//开始记录日志的逻辑
		cost := time.Since(start)
		responStatus := c.Writer.Status()
		logFileds := []zap.Field{
			zap.Int("status",responStatus),
			zap.String("request",c.Request.Method+""+c.Request.URL.String()),
			zap.String("query",c.Request.URL.RawQuery),
			zap.String("ip", c.ClientIP()),
			zap.String("user-agent", c.Request.UserAgent()),
			zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
			zap.String("time", helpers.MicrosecondsStr(cost)),
		}
		if c.Request.Method == "POST" || c.Request.Method == "PUT" || c.Request.Method == "DELETE"{
			logFileds = append(logFileds,zap.String("Request Body",string(requestBody)))
			logFileds = append(logFileds,zap.String("Response Body",w.body.String()))
		}
		if responStatus > 400 && responStatus <= 499 {
			logger.Warn("HTTP Warning" + cast.ToString(responStatus),logFileds...)
		}else{
			logger.Debug("HTTP Access Log",logFileds...)
		}
	}
}
