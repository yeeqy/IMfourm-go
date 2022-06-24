package auth

import (
	v1 "IMfourm-go/app/http/controllers/api/v1"
	"IMfourm-go/app/requests"
	"IMfourm-go/pkg/captcha"
	"IMfourm-go/pkg/logger"
	"IMfourm-go/pkg/response"
	"IMfourm-go/pkg/verifycode"
	"github.com/gin-gonic/gin"
)

//用户控制器
type VerifyCodeController struct {
	v1.BaseAPIController
}
//显示图片验证码
func (vc *VerifyCodeController) ShowCaptcha(c *gin.Context){
	//生成验证码
	id,b64s, err := captcha.NewCaptcha().GenerateCaptcha()
	//记录错误日志
	//因为验证码是用户入口，出错时应记error等级的日志
	logger.LogIf(err)

	//返回给用户
	//c.JSON(http.StatusOK,gin.H{
	//	"captcha_id":id,
	//	"captcha)image":b64s,
	//})
	response.JSON(c,gin.H{
		"captcha_id":id,
		"captcha)image":b64s,
	})
}
//发送手机验证码
func (vc *VerifyCodeController) SendUsingPhone(c *gin.Context){
	//1. 验证表单
	//2. 发送SMS

	req := requests.VerifyCodePhoneRequest{}
	if ok := requests.Validate(c,&req,requests.VerifyCodePhone); !ok {
		return
	}

	if ok := verifycode.NewVerifyCode().SendSMS(req.Phone); !ok {
		response.Abort500(c,"发送短信失败")
	}else {
		response.Success(c)
	}
}