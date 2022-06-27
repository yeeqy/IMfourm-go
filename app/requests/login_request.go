package requests

import (
	"IMfourm-go/app/requests/validators"
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

//手机+短信验证码登录

type LoginByPhoneRequest struct {
	Phone      string `json:"phone,omitempty" valid:"phone"`
	VerifyCode string `json:"verify_code,omitempty" valid:"verify_code"`
}

//验证表单  返回长度=0即通过
func LoginByPhone(data interface{}, c *gin.Context) map[string][]string {
	rules := govalidator.MapData{
		"phone":       []string{"required", "digits:11"},
		"verify_code": []string{"required", "digits:6"},
	}
	messages := govalidator.MapData{
		"phone": []string{
			"required:手机号为必填项，参数名称phone",
			"digits:手机号长度必须为11位数字",
		},
		"verify_code": []string{
			"required:验证码答案必填",
			"digits:验证码长度必须为6位数字",
		},
	}
	errs := validate(data, rules, messages)

	//手机验证码
	_data := data.(*LoginByPhoneRequest)
	errs = validators.ValidateVerifyCode(_data.Phone, _data.VerifyCode, errs)
	return errs
}

//使用 手机/email/用户名 + 密码 登录

type LoginByPasswordRequest struct {
	CaptchaID     string `json:"captcha_id,omitempty" valid:"captcha_id"`
	CaptchaAnswer string `json:"captcha_answer,omitempty" valid:"captcha_answer"`
	LoginID       string `json:"login_id" valid:"login_id"`
	Password      string `json:"password,omitempty" valid:"password"`
}
//验证表单
func LoginByPassword(data interface{}, c *gin.Context) map[string][]string{
	rules := govalidator.MapData{
		"login_id": []string{"required","min:3"},
		"password":[]string{"required","min:6"},
		"captcha_id":[]string{"required"},
		"captcha_answer":[]string{"required","digits:6"},
	}
	messages := govalidator.MapData{
		"login_id":[]string{
			"required:登录ID为必填项，支持手机号、邮箱、用户名",
			"min:登录ID长度需大于3",
		},
		"password":[]string{
			"required:密码为必填项",
			"min:密码长度需大于6",
		},
		"captcha_id":[]string{
			"required:图片验证码ID为必填项",
		},
		"captcha_answer": []string{
			"required:图片验证码答案必填",
			"digits:图片验证码长度必须为6位数字",
		},
	}
	errs := validate(data,rules,messages)

	//图片验证码
	_data := data.(*LoginByPasswordRequest)
	errs = validators.ValidateCaptcha(_data.CaptchaID,_data.CaptchaAnswer,errs)
	return errs
}
