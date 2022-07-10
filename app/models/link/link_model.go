package link

//模型模板中放进去常用的方法，使用 FIXME() 这个不存在的函数，通知用户要记得修改这个地方；
import (
	"IMfourm-go/app/models"
	"IMfourm-go/pkg/database"
)

type Link struct {
	models.BaseModel
	//put fields in here
	Name string `json:"name,omitempty"`
	URL  string `json:"url,omitempty"`

	models.CommonTimestampsField
}

func (link *Link) Create() {
	database.DB.Create(&link)
}

func (link *Link) Save() (rowsAffected int64) {
	result := database.DB.Save(&link)
	return result.RowsAffected
}
func (link *Link) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&link)
	return result.RowsAffected
}
