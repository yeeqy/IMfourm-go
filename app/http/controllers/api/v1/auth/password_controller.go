// Package auth 处理用户注册、登录、密码重置
package auth

import (
	v1 "IMfourm-go/app/http/controllers/api/v1"
	"IMfourm-go/app/models/user"
	"IMfourm-go/app/requests"
	"IMfourm-go/pkg/response"
	"github.com/gin-gonic/gin"
)

//用户控制器
type PasswordController struct {
	v1.BaseAPIController
}

//使用手机和验证码重置密码
func (pc *PasswordController) ResetByPhone(c *gin.Context){
	//1. 验证表单
	req := requests.ResetByPhoneRequest{}
	if ok := requests.Validate(c,&req,requests.ResetByPhone);!ok{
		return
	}
	//2. 更新密码
	userModel := user.GetByPhone(req.Phone)
	if userModel.ID == 0 {
		response.Abort404(c)
	} else {
		userModel.Password = req.Password
		userModel.Save()

		response.Success(c)
	}
}

//使用Email和验证码重置验证码
func(pc *PasswordController) ResetByEmail(c *gin.Context){
	//1. 验证表单
	req := requests.ResetByEmailRequest{}
	if ok := requests.Validate(c,&req,requests.ResetByEmail); !ok {
		return
	}
	//2. 更新密码
	userModel := user.GetByEmail(req.Email)
	if userModel.ID == 0 {
		response.Abort404(c)
	} else {
		userModel.Password = req.Password
		userModel.Save()
		response.Success(c)
	}
}