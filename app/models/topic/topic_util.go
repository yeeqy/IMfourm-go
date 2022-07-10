package topic
import (
    "IMfourm-go/pkg/app"
    "IMfourm-go/pkg/database"
    "IMfourm-go/pkg/paginator"
    "gorm.io/gorm/clause"

    "github.com/gin-gonic/gin"
)

//model_util 模板文件我们放进去一些常用的方法：

func Get(idstr string)(topic Topic){
    database.DB.Preload(clause.Associations).Where("id",idstr).First(&topic)
    return
}
func GetBy(field,value string)(topic Topic){
    database.DB.Where("? = ?",field,value).First(&topic)
    return
}
func All()(topics []Topic){
    database.DB.Find(&topics)
    return
}
func IsExist(field, value string) bool {
    var count int64
    database.DB.Model(Topic{}).Where(" = ?", field,value).Count(&count)
    return count > 0
}
func Paginate(c *gin.Context, perPage int) (topics []Topic, paging paginator.Paging) {
    paging = paginator.Paginate(
        c,
        database.DB.Model(Topic{}),
        &topics,
        app.V1URL(database.TableName(&Topic{})),
        perPage,
    )
    return
}