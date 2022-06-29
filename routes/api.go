package routes
//目录存放 我们所有项目的路由文件
import (
	"IMfourm-go/app/http/controllers/api/v1/auth"
	"github.com/gin-gonic/gin"
)

func RegisterApiRoutes(r *gin.Engine){

	//测试一个v1 的路由组，我们所有的v1版本的路由都将存放到这里
	v1 := r.Group("/v1")
	{
		authGroup := v1.Group("/auth")
		{
			suc := new(auth.SignupController)
			authGroup.POST("/signup/phone/exist",suc.IsPhoneExist)
			authGroup.POST("/signup/email/exist",suc.IsEmailExist)

			//发送验证码
			vcc := new(auth.VerifyCodeController)
			//图片验证码，需要加限流
			authGroup.POST("/verify-codes/captcha",vcc.ShowCaptcha)
			authGroup.POST("/verify-codes/phone",vcc.SendUsingPhone)
			//注册
			authGroup.POST("/signup/using-phone",suc.SignupUsingPhone)
			authGroup.POST("/signup/using-email",suc.SignupUsingEmail)

			//登录
			lgc := new(auth.LoginController)
			//使用手机号，短信验证码进行登录
			authGroup.POST("/login/using-phone",lgc.LoginByPhone)
			//支持手机号，Email和用户名
			authGroup.POST("/login/using-password",lgc.LoginByPassword)
			authGroup.POST("/login/refresh-token",lgc.RefreshToken)

			//重置密码
			pwc := new(auth.PasswordController)
			authGroup.POST("/password-reset/using-phone",pwc.ResetByPhone)
			authGroup.POST("/password-reset/using-email",pwc.ResetByEmail)
		}
	}
}
