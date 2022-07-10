package routes
//目录存放 我们所有项目的路由文件
import (
	controllers "IMfourm-go/app/http/controllers/api/v1"
	"IMfourm-go/app/http/controllers/api/v1/auth"
	"IMfourm-go/app/http/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterApiRoutes(r *gin.Engine){

	//测试一个v1 的路由组，我们所有的v1版本的路由都将存放到这里
	v1 := r.Group("/v1")

	// 全局限流中间件：每小时限流。这里是所有 API （根据 IP）请求加起来。
	// 作为参考 Github API 每小时最多 60 个请求（根据 IP）。
	// 测试时，可以调高一点。
	v1.Use(middlewares.LimitIP("200-H"))

	{
		authGroup := v1.Group("/auth")
		// 限流中间件：每小时限流，作为参考 Github API 每小时最多 60 个请求（根据 IP）
		// 测试时，可以调高一点
		authGroup.Use(middlewares.LimitIP("1000-H"))
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
			authGroup.POST( "/password-reset/using-phone",pwc.ResetByPhone)
			authGroup.POST("/password-reset/using-email",pwc.ResetByEmail)

		}

		uc := new(controllers.UsersController)
		//获取当前用户
		v1.GET("/user",middlewares.AuthJWT(),uc.CurrentUser)
		userGroup := v1.Group("/users")
		{
			userGroup.GET("",uc.Index)
		}

		cgc := new(controllers.CategoriesController)
		cgcGroup := v1.Group("/categories")
		{
			cgcGroup.GET("",cgc.Index)
			//登录用户才能创建分类，所以用了AuthJWT中间件
			cgcGroup.POST("",middlewares.AuthJWT(),cgc.Store)
			cgcGroup.PUT("/:id",middlewares.AuthJWT(),cgc.Update)
			cgcGroup.DELETE("/:id",middlewares.AuthJWT(),cgc.Delete)
		}

		tpc := new(controllers.TopicsController)
		tpcGroup := v1.Group("/topics")
		{
			tpcGroup.GET("",tpc.Index)
			tpcGroup.POST("",middlewares.AuthJWT(),tpc.Store)
			tpcGroup.PUT("/:id",middlewares.AuthJWT(),tpc.Update)
			tpcGroup.DELETE("/:id",middlewares.AuthJWT(),tpc.Delete)
			tpcGroup.GET("/:id",tpc.Show)
		}

		lsc := new(controllers.LinksController)
		linksGroup := v1.Group("/links")
		{
			linksGroup.GET("",lsc.Index)
		}
	}
}
