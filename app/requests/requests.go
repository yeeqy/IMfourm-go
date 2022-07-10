package requests
//处理请求数据和表单验证

import (
	"IMfourm-go/pkg/response"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
	"net/http"
)

// ValidatorFunc 验证函数类型
type ValidatorFunc func(interface{},*gin.Context)map[string][]string

func Validate(c *gin.Context, obj interface{}, handler ValidatorFunc) bool {
	//1. 解析请求
	if err := c.ShouldBind(obj); err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity,gin.H{
			"message":"请求解析错误，请确认请求格式是否正确。上传文件multipart标头，参数JSON格式",
			"error": err.Error(),
		})
		fmt.Println(err.Error())
		return false
	}
	//2. 表单验证
	errs := handler(obj,c)

	//3. 判断验证是否通过
	if len(errs) > 0 {
		//c.AbortWithStatusJSON(http.StatusUnprocessableEntity,gin.H{
		//	"message":"请求验证不通过，具体请看errors",
		//	"errors":errs,
		//})
		response.ValidationError(c,errs)
		return false
	}
	return true
}

func validate(data interface{},rules govalidator.MapData,messages govalidator.MapData) map[string][]string{
	opts := govalidator.Options{
		Data: data,
		Rules: rules,
		TagIdentifier: "valid",
		Messages: messages,
	}
	return govalidator.New(opts).ValidateStruct()
}

func validateFile(c *gin.Context, data interface{}, rules govalidator.MapData, messages govalidator.MapData) map[string][]string  {
	opts := govalidator.Options{
		Request: c.Request,
		Rules: rules,
		Messages: messages,
		TagIdentifier: "valid",
	}
	//调用go validator的validate方法来验证
	return govalidator.New(opts).Validate()
}