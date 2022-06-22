package auth

import (
	v1 "IMfourm-go/app/http/controllers/api/v1"
	"IMfourm-go/pkg/captcha"
	"IMfourm-go/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
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
	c.JSON(http.StatusOK,gin.H{
		"captcha_id":id,
		"captcha)image":b64s,
	})

}
