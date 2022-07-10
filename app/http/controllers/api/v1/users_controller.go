package v1

import (
	"IMfourm-go/app/models/user"
	"IMfourm-go/app/requests"
	"IMfourm-go/pkg/auth"
	"IMfourm-go/pkg/config"
	"IMfourm-go/pkg/file"
	"IMfourm-go/pkg/response"
	"github.com/gin-gonic/gin"
)

type UsersController struct {
	BaseAPIController
}

func (ctrl *UsersController) CurrentUser(c *gin.Context){
	userModel := auth.CurrentUser(c)
	response.Data(c,userModel)
}

// Index 所有用户
func (ctrl *UsersController) Index(c *gin.Context)  {
	req := requests.PaginationRequest{}
	if ok := requests.Validate(c,&req,requests.Pagination); !ok{
		return
	}

	data,pager := user.Paginate(c,10)
	response.JSON(c,gin.H{
		"data":data,
		"pager":pager,
	})
}

func (ctrl *UsersController) UpdateProfile(c *gin.Context)  {
	req := requests.UserUpdateProfileRequest{}
	if ok := requests.Validate(c,&req,requests.UserUpdateProfile);!ok{
		return
	}
	currentUser := auth.CurrentUser(c)
	currentUser.Name = req.Name
	currentUser.City = req.City
	currentUser.Introduction = req.Introduction

	rowsAffected := currentUser.Save()

	if rowsAffected > 0 {
		response.Data(c,currentUser)
	} else {
		response.Abort500(c,"更新失败，请稍后尝试")
	}

}

func (ctrl *UsersController) UpdateEmail(c *gin.Context)  {
	req := requests.UserUpdateEmailRequest{}
	if ok:= requests.Validate(c,&req,requests.UserUpdateEmail);!ok{
		return
	}
	currentUser := auth.CurrentUser(c)
	currentUser.Email = req.Email

	rowsAffected := currentUser.Save()

	if rowsAffected > 0 {
		response.Success(c)
	} else {
		response.Abort500(c,"更新失败，请稍后尝试")
	}

}

func (ctrl *UsersController) UpdatePhone(c *gin.Context) {

	request := requests.UserUpdatePhoneRequest{}
	if ok := requests.Validate(c, &request, requests.UserUpdatePhone); !ok {
		return
	}

	currentUser := auth.CurrentUser(c)
	currentUser.Phone = request.Phone
	rowsAffected := currentUser.Save()

	if rowsAffected > 0 {
		response.Success(c)
	} else {
		response.Abort500(c, "更新失败，请稍后尝试~")
	}
}

func (ctrl *UsersController) UpdatePassword(c *gin.Context) {
	request := requests.UserUpdatePasswordRequest{}
	if ok := requests.Validate(c, &request, requests.UserUpdatePassword); !ok {
		return
	}
	currentUser := auth.CurrentUser(c)
	// 验证原始密码是否正确
	_, err := auth.Attempt(currentUser.Name, request.Password)
	if err != nil {
		// 失败，显示错误提示
		response.Unauthorized(c, "原密码不正确")
	} else {
		// 更新密码为新密码
		currentUser.Password = request.NewPassword
		currentUser.Save()
	}
}

func (ctrl *UsersController) UpdateAvatar(c *gin.Context)  {
	req := requests.UserUpdateAvatarRequest{}
	if ok := requests.Validate(c,&req,requests.UserUpdateAvatar);!ok{
		return
	}
	//处理文件上传的逻辑封装在file.SaveUploadAvatar方法里
	avatar,err := file.SaveUploadAvatar(c,req.Avatar)
	if err!=nil{
		response.Abort500(c,"上传头像失败，请稍后尝试")
		return
	}

	currentUser := auth.CurrentUser(c)
	currentUser.Avatar = config.GetString("app.url") + avatar
	currentUser.Save()

	response.Data(c,currentUser)

}