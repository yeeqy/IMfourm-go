package category

//模型模板中放进去常用的方法，使用 FIXME() 这个不存在的函数，通知用户要记得修改这个地方；
import (
	"IMfourm-go/app/models"
	"IMfourm-go/pkg/database"
)

type Category struct {
	models.BaseModel
	//put fields in here
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`

	models.CommonTimestampsField
}

func (category *Category) Create() {
	database.DB.Create(&category)
}

func (category *Category) Save() (rowsAffected int64) {
	result := database.DB.Save(&category)
	return result.RowsAffected
}
func (category *Category) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&category)
	return result.RowsAffected
}
