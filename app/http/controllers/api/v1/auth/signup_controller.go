package auth

import (
	v1 "IMfourm-go/app/http/controllers/api/v1"
	"IMfourm-go/app/models/user"
	"IMfourm-go/app/requests"
	"github.com/gin-gonic/gin"
	"net/http"
)

//注册控制器
type SignupController struct {
	v1.BaseAPIController
}

//检测手机号是否被注册
func (sc *SignupController) IsPhoneExist(c *gin.Context){

	//请求对象
	//获取请求数据，并作表单验证
	request := requests.SignupPhoneExistRequest{}
	if ok := requests.Validate(c,&request,requests.SignupPhoneExist);!ok{
		return
	}

	//解析JSON请求
	//if err := c.ShouldBindJSON(&request); err!=nil{
	//	//解析失败，返回422状态码和错误信息
	//	c.AbortWithStatusJSON(http.StatusUnprocessableEntity,gin.H{
	//		"error":err.Error(),
	//	})
	//	fmt.Println(err.Error())
	//	return
	//}
	////表单验证
	//errs := requests.ValidateSignupPhoneExist(&request,c)
	////errs 返回长度=0即通过，>0则表示有错误发生
	//if len(errs) > 0 {
	//	//验证失败，返回422状态码和错误状态
	//	c.AbortWithStatusJSON(http.StatusUnprocessableEntity,gin.H{
	//		"errors":errs,
	//	})
	//}

	//检查数据库并返回响应
	c.JSON(http.StatusOK,gin.H{
		"exist":user.IsPhoneExist(request.Phone),
	})
}

func (sc *SignupController) IsEmailExist(c *gin.Context){
	//初始化请求对象
	request := requests.SignupEmailExistRequest{}

	//if err := c.ShouldBindJSON(&request); err != nil {
	//	c.AbortWithStatusJSON(http.StatusUnprocessableEntity,gin.H{
	//		"error":err.Error(),
	//	})
	//	fmt.Println(err.Error())
	//	return
	//}
	////表单验证
	//errs := requests.ValidateSignupEmailExist(&request,c)
	//if len(errs) > 0{
	//	c.AbortWithStatusJSON(http.StatusUnprocessableEntity,gin.H{
	//		"error":errs,
	//	})
	//	return
	//}

	if ok := requests.Validate(c,&request,requests.SignupEmailExist);!ok{
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"exist": user.IsEmailExist(request.Email),
	})
}