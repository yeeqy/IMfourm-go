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
func ResetByPhone(data interface{},c *gin.Context) map[string][]string{
	rules := govalidator.MapData{
		"phone":[]string{"required","digits:11"},
		"verify_code":[]string{"required","digits:6"},
		"password":[]string{"required","min:6"},
	}
	messages := govalidator.MapData{
		"phone":[]string{
			"required:手机号为必填项，参数名称phone",
			"digits:手机号长度必须为11位的数字",
		},
		"verify_code":[]string{
			"required:验证码答案必填",
			"digits:验证码长度必须为6位的数字",
		},
		"password":[]string{
			"required:密码为必填项",
			"min:密码长度需大于6",
		},
	}
	errs := validate(data,rules,messages)
	//检查验证码
	_data := data.(*ResetByPhoneRequest)
	errs = validators.ValidateVerifyCode(_data.Phone,_data.VerifyCode,errs)

	return errs
}
