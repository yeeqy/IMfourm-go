package requests

import (
	"IMfourm-go/app/requests/validators"
	"IMfourm-go/pkg/auth"
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
	"mime/multipart"
)

type UserUpdateProfileRequest struct {
	Name         string `valid:"name" json:"name"`
	City         string `valid:"city" json:"city"`
	Introduction string `valid:"introduction" json:"introduction"`
}

func UserUpdateProfile(data interface{}, c *gin.Context) map[string][]string {

	//查询用户名重复时，过滤掉当前用户id
	uid := auth.CurrentUID(c)

	rules := govalidator.MapData{
		"name":         []string{"required", "alpha_num", "between:3,20", "not_exists:users,name," + uid},
		"introduction": []string{"min_cn:4", "max_cn:240"},
		"city":         []string{"min_cn:2", "max_cn:20"},
	}
	messages := govalidator.MapData{
		"name": []string{
			"required:用户名为必填项",
			"alpha_num:用户名格式错误，只允许数字和英文",
			"between:用户名长度需要在3~20之间",
			"not_exists:用户名已被占用",
		},
		"introduction": []string{
			"min_cn:描述长度需至少 4 个字",
			"max_cn:描述长度不能超过 240 个字",
		},
		"city": []string{
			"min_cn:城市需至少2个字",
			"max_cn:城市不能超过20个字",
		},
	}
	return validate(data, rules, messages)
}

type UserUpdateEmailRequest struct {
	Email      string `json:"email,omitempty" valid:"email"`
	VerifyCode string `json:"verify_code,omitempty" valid:"verify_code"`
}

func UserUpdateEmail(data interface{}, c *gin.Context) map[string][]string {
	currentUser := auth.CurrentUser(c)
	rules := govalidator.MapData{
		"email": []string{"required", "min:4", "max:30", "email",
			"not_exists:user,email," + currentUser.GetStringID(),
			"not_in:" + currentUser.Email,
		},
		//这个digits写成了digit，报了服务器内部错误说
		"verify_code": []string{"required", "digits:6"},
	}
	messages := govalidator.MapData{
		"email": []string{
			"required:Email为必填项",
			"min:Email长度需大于4",
			"max:Email长度需小于30",
			"not_exists:Email已被占用",
			"not_in:新的Email与旧Email一致",
		},
		"verify_code": []string{
			"required:验证码为必填项",
			"digits:验证码长度必须为6位数字",
		},
	}
	errs := validate(data, rules, messages)
	_data := data.(*UserUpdateEmailRequest)
	errs = validators.ValidateVerifyCode(_data.Email, _data.VerifyCode, errs)
	return errs
}

type UserUpdatePhoneRequest struct {
	Phone      string `json:"phone,omitempty" valid:"phone"`
	VerifyCode string `json:"verify_code,omitempty" valid:"verify_code"`
}

func UserUpdatePhone(data interface{}, c *gin.Context) map[string][]string {
	currentUser := auth.CurrentUser(c)

	rules := govalidator.MapData{
		"phone": []string{
			"required",
			"digits:11",
			"not_exists:users,phone," + currentUser.GetStringID(),
			"not_in:" + currentUser.Phone,
		},
		"verify_code": []string{"required", "digits:6"},
	}
	messages := govalidator.MapData{
		"phone": []string{
			"required:手机号为必填项，参数名称 phone",
			"digits:手机号长度必须为 11 位的数字",
			"not_exists:手机号已被占用",
			"not_in:新的手机与老手机号一致",
		},
		"verify_code": []string{
			"required:验证码答案必填",
			"digits:验证码长度必须为 6 位的数字",
		},
	}

	errs := validate(data, rules, messages)
	_data := data.(*UserUpdatePhoneRequest)
	errs = validators.ValidateVerifyCode(_data.Phone, _data.VerifyCode, errs)

	return errs
}

type UserUpdatePasswordRequest struct {
	Password           string `json:"password,omitempty" valid:"password"`
	NewPassword        string `json:"new_password,omitempty" valid:"new_password"`
	NewPasswordConfirm string `json:"new_password_confirm,omitempty" valid:"new_password_confirm"`
}

func UserUpdatePassword(data interface{}, c *gin.Context) map[string][]string {
	rules := govalidator.MapData{
		"password":             []string{"required", "min:6"},
		"new_password":         []string{"required", "min:6"},
		"new_password_confirm": []string{"required", "min:6"},
	}
	messages := govalidator.MapData{
		"password": []string{
			"required:密码为必填项",
			"min:密码长度需大于 6",
		},
		"new_password": []string{
			"required:密码为必填项",
			"min:密码长度需大于 6",
		},
		"new_password_confirm": []string{
			"required:确认密码框为必填项",
			"min:确认密码长度需大于 6",
		},
	}
	//确保confirm密码正确
	errs := validate(data,rules,messages)
	_data := data.(*UserUpdatePasswordRequest)
	errs = validators.ValidatePasswordConfirm(_data.NewPassword,_data.NewPasswordConfirm,errs)
	return errs
}

//上传用户头像

type UserUpdateAvatarRequest struct {
	//*multipart.FileHeader 是 Go validator 验证文件必须使用的类型。
	Avatar *multipart.FileHeader `form:"avatar" valid:"avatar"`
}

func UserUpdateAvatar(data interface{},c *gin.Context) map[string][]string  {
	rules := govalidator.MapData{
		//Go validator 规定了，验证文件必须在字段前方加 file: 前缀
		//size单位bytes，20971520--20mb
		"file:avatar":[]string{"ext:png,jpg,jpeg","size:20971520","required"},
	}
	messages := govalidator.MapData{
		"file:avatar":[]string{
			"ext:ext头像只能上传png,jpg,jpeg 任意一种图片",
			"size:头像文件最大不能超过20MB",
			"required:必须上传图片",
		},
	}
	return validateFile(c,data,rules,messages)
}