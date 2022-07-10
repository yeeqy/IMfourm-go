package link
import (
    "IMfourm-go/pkg/app"
    "IMfourm-go/pkg/cache"
    "IMfourm-go/pkg/database"
    "IMfourm-go/pkg/helpers"
    "IMfourm-go/pkg/paginator"
    "time"

    "github.com/gin-gonic/gin"
)

//model_util 模板文件我们放进去一些常用的方法：

func Get(idstr string)(link Link){
    database.DB.Where("id",idstr).First(&link)
    return
}
func GetBy(field,value string)(link Link){
    database.DB.Where("? = ?",field,value).First(&link)
    return
}
func All()(links []Link){
    database.DB.Find(&links)
    return
}
func IsExist(field, value string) bool {
    var count int64
    database.DB.Model(Link{}).Where(" = ?", field,value).Count(&count)
    return count > 0
}
func Paginate(c *gin.Context, perPage int) (links []Link, paging paginator.Paging) {
    paging = paginator.Paginate(
        c,
        database.DB.Model(Link{}),
        &links,
        app.V1URL(database.TableName(&Link{})),
        perPage,
    )
    return
}
func AllCache()(links []Link)  {
    //设置缓存key
    cacheKey := "links:all"
    //设置过期时间
    expireTime := 120 * time.Minute
    //取数据
    cache.GetObject(cacheKey,&links)
    //如果数据为空
    if helpers.Empty(links){
        //查询数据库
        links = All()
        if helpers.Empty(links){
            return links
        }
        //设置缓存
        cache.Set(cacheKey,links,expireTime)
    }
    return
}