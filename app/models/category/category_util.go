package category
import (
    "IMfourm-go/pkg/app"
    "IMfourm-go/pkg/database"
    "IMfourm-go/pkg/paginator"

     "github.com/gin-gonic/gin"
)

//model_util 模板文件我们放进去一些常用的方法：

func Get(idstr string)(category Category){
    database.DB.Where("id",idstr).First(&category)
    return
}
func GetBy(field,value string)(category Category){
    database.DB.Where("? = ?",field,value).First(&category)
    return
}
func All()(categories []Category){
    database.DB.Find(&categories)
    return
}
func IsExist(field, value string) bool {
    var count int64
    database.DB.Model(Category{}).Where(" = ?", field,value).Count(&count)
}
func Paginate(c *gin.Context, perPage int) (categories []Category, paging paginator.Paging) {
    paging = paginator.Paginate(
        c,
        database.DB.Model(Category{}),
        &categories,
        app.V1URL(database.TableName(&Category{})),
        perPage,
    )
    return
}