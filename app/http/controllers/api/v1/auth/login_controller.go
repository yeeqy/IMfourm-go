package auth

import (
	v1 "IMfourm-go/app/http/controllers/api/v1"
	"IMfourm-go/app/requests"
	"IMfourm-go/pkg/auth"
	"IMfourm-go/pkg/jwt"
	"IMfourm-go/pkg/response"
	"github.com/gin-gonic/gin"
)

//用户控制器
type LoginController struct {
	v1.BaseAPIController
}
//手机登录
func (lc *LoginController) LoginByPhone(c * gin.Context){
	//1. 验证表单
	req := requests.LoginByPhoneRequest{}
	if ok := requests.Validate(c,&req,requests.LoginByPhone);!ok {
		return
	}
	//2. 尝试登陆
	user,err := auth.LoginByPhone(req.Phone)
	if err != nil {
		response.Error(c,err,"账号不存在或密码错误")
	}else{
		//登陆成功
		token := jwt.NewJWT().IssueToken(user.GetStringID(),user.Name)
		response.JSON(c,gin.H{
			"token":token,
		})
	}
}

//多种方法登录
func (lc *LoginController) LoginByPassword(c *gin.Context){
	//1. 验证表单
	req := requests.LoginByPasswordRequest{}
	if ok := requests.Validate(c,&req,requests.LoginByPassword);!ok{
		return
	}
	//2. 尝试登陆
	user ,err := auth.Attempt(req.LoginID,req.Password)
	if err != nil {
		//这样可以显示错误原因
		//response.JSON(c,gin.H{
		//	"data":"登陆失败",
		//	"error":err.Error(),
		//})
		response.Unauthorized(c,"登陆失败")
	} else {
		token := jwt.NewJWT().IssueToken(user.GetStringID(),user.Name)
		response.JSON(c,gin.H{
			"token":token,
		})
	}
}

// 刷新Access Token
func (lc *LoginController) RefreshToken(c *gin.Context){
	token,err := jwt.NewJWT().RefreshToken(c)
	if err != nil {
		response.Error(c,err,"令牌刷新失败")
	} else {
		response.JSON(c,gin.H{
			"token":token,
		})
	}
}