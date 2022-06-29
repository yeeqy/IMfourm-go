package requests

import (
	"IMfourm-go/app/requests/validators"
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type ResetByPhoneRequest struct {
	Phone      string `json:"phone,omitempty" valid:"phone"`
	VerifyCode string `json:"verify_code,omitempty" valid:"verify_code"`
	Password   string `json:"password,omitempty" valid:"password"`
}

//验证表单
func ResetByPhone(data interface{}, c *gin.Context) map[string][]string {
	rules := govalidator.MapData{
		"phone":       []string{"required", "digits:11"},
		"verify_code": []string{"required", "digits:6"},
		"password":    []string{"required", "min:6"},
	}
	messages := govalidator.MapData{
		"phone": []string{
			"required:手机号为必填项，参数名称phone",
			"digits:手机号长度必须为11位的数字",
		},
		"verify_code": []string{
			"required:验证码答案必填",
			"digits:验证码长度必须为6位的数字",
		},
		"password": []string{
			"required:密码为必填项",
			"min:密码长度需大于6",
		},
	}
	errs := validate(data, rules, messages)
	//检查验证码
	_data := data.(*ResetByPhoneRequest)
	errs = validators.ValidateVerifyCode(_data.Phone, _data.VerifyCode, errs)

	return errs
}

type ResetByEmailRequest struct {
	Email      string `json:"email,omitempty" valid:"email"`
	VerifyCode string `json:"verify_code,omitempty" valid:"verify_code"`
	Password   string `json:"password,omitempty" valid:"password"`
}
func ResetByEmail(data interface{}, c *gin.Context) map[string][]string{
	rules := govalidator.MapData{
		"email":[]string{"required","min:4","max:30","email"},
		"verify_code":[]string{"required","digits:6"},
		"password":[]string{"required","min:6"},
	}
	messages := govalidator.MapData{
		"email":[]string{
			"required:Email为必填项",
			"min:Email长度需大于4",
			"max:Email长度需小于30",
			"email:Email格式不正确，请提供有效的邮箱地址",
		},
		"verify_code":[]string{
			"required:验证码答案必填",
			"digits:验证码长度必须为6位数字",
		},
		"password":[]string{
			"required:密码为必填项",
			"min:密码长度需大于6",
		},
	}
	errs := validate(data,rules,messages)

	_data := data.(*ResetByEmailRequest)
	errs = validators.ValidateVerifyCode(_data.Email,_data.VerifyCode,errs)
	return errs
}