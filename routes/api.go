package routes
//目录存放 我们所有项目的路由文件
import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterApiRoutes(r *gin.Engine){

	//测试一个v1 的路由组，我们所有的v1版本的路由都将存放到这里
	v1 := r.Group("/v1")
	{
		v1.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK,gin.H{
				"hello":"yeeqy",
			})
		})
	}
}
