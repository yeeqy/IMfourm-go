package response

import (
	"IMfourm-go/pkg/logger"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

//方便后续统一输出格式

//内用的辅助函数，用以支持默认参数默认值
//Go不支持参数默认值，只能使用多变参数来实现类似效果
func defaultMessage(defaultMsg string, msg ...string)(message string){
	if len(msg) > 0 {
		message = msg[0]
	}else {
		message = defaultMsg
	}
	return
}

// JSON 响应200和JSON数据
func JSON(c *gin.Context, data interface{}){
	c.JSON(http.StatusOK, data)
}
//Success 响应200和 预设操作成功的JSON数据
func Success(c *gin.Context){
	JSON(c, gin.H{
		"success":true,
		"message":"操作成功！",
	})
}

//Data 响应 200 和 带data键的 JSON数据
func Data(c *gin.Context,data interface{}){
	JSON(c,gin.H{
		"success":true,
		"data": data,
	})
}
//响应201 和 带data键的JSON数据
func Created(c *gin.Context, data interface{}){
	c.JSON(http.StatusCreated,gin.H{
		"success":true,
		"data":data,
	})
}
func CreatedJSON(c *gin.Context,data interface{})  {
	c.JSON(http.StatusCreated,data)
}
//响应404，未传参msg时使用默认消息
func Abort404(c *gin.Context,msg ...string){
	c.AbortWithStatusJSON(http.StatusNotFound,gin.H{
		"message":defaultMessage("数据不存在，请确定请求正确", msg...),
	})
}
//响应403，未传参msg时使用默认消息
func Abort403(c *gin.Context,msg ...string){
	c.AbortWithStatusJSON(http.StatusForbidden,gin.H{
		"message":defaultMessage("权限不足，请确定您有对应的权限", msg...),
	})
}
//响应500，未传参msg时使用默认消息
func Abort500(c *gin.Context,msg ...string){
	c.AbortWithStatusJSON(http.StatusInternalServerError,gin.H{
		"message":defaultMessage("服务器内部错误，请稍后再试", msg...),
	})
}
//响应400，传参err对象，未传参msg时使用默认消息
func BadRequest(c *gin.Context,err error,msg ...string){
	logger.LogIf(err)
	c.AbortWithStatusJSON(http.StatusBadRequest,gin.H{
		"message":defaultMessage("请求解析错误，请确认请求格式是否正确。上传文件请使用multipart标头，参数请使用JSON格式", msg...),
		"error":err.Error(),
	})
}
//响应404或422，未传参msg时使用默认消息
func Error(c *gin.Context,err error,msg ...string){
	logger.LogIf(err)
	if err == gorm.ErrRecordNotFound{
		Abort404(c)
		return
	}
	c.AbortWithStatusJSON(http.StatusUnprocessableEntity,gin.H{
		"message":defaultMessage("请求处理失败，请查看error的值",msg...),
		"error":err.Error(),
	})
}
//处理表单验证不通过的错误
func ValidationError(c *gin.Context,errors map[string][]string){
	c.AbortWithStatusJSON(http.StatusUnprocessableEntity,gin.H{
		"message":"请求验证不通过，具体请查看errors",
		"errors":errors,
	})
}
//响应401，登陆失败、jwt解析失败时调用
func Unauthorized(c *gin.Context,msg ...string){
	c.AbortWithStatusJSON(http.StatusUnauthorized,gin.H{
		"message":defaultMessage("请求解析错误，请确认请求格式是否正确",msg...),
	})
}

