package middlewares
//Auth中间件
//Auth 中间件用在一些需要用户授权才能操作的接口，例如说创建话题、更新个人资料等。
//使用方法是在路由中做绑定，哪些路由需要授权才能访问，直接为其添加中间件即可。

import (
	"IMfourm-go/app/models/user"
	"IMfourm-go/pkg/config"
	"IMfourm-go/pkg/jwt"
	"IMfourm-go/pkg/response"
	"fmt"
	"github.com/gin-gonic/gin"
)

func AuthJWT() gin.HandlerFunc{
	return func(c *gin.Context) {
		//从标头 Authorization:Bearer xxxx中获取信息，并验证JWT的准确性
		claims,err := jwt.NewJWT().ParserToken(c)
		//JWT解析失败，有错误发生
		if err!=nil {
			response.Unauthorized(c,fmt.Sprintf("请查看%v相关的接口认证文档",config.GetString("app.name")))
			return
		}
		//JWT解析成功，设置用户信息
		userModel := user.Get(claims.UserID)
		if userModel.ID == 0 {
			response.Unauthorized(c,"找不到对应用户，用户可能已删除")
			return
		}
		//将用户信息存入gin.context里，后续auth包将从这里拿到当前用户数据
		c.Set("current_user_id",userModel.GetStringID())
		c.Set("current_user_name",userModel.Name)
		c.Set("current_user",userModel)

		c.Next()
	}
}
