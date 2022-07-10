package topic

//模型模板中放进去常用的方法，使用 FIXME() 这个不存在的函数，通知用户要记得修改这个地方；
import (
	"IMfourm-go/app/models"
	"IMfourm-go/app/models/category"
	"IMfourm-go/app/models/user"
	"IMfourm-go/pkg/database"
)

type Topic struct {
	models.BaseModel
	//put fields in here
	Title      string `json:"title,omitempty"`
	Body       string `json:"body,omitempty"`
	UserID     string `json:"user_id,omitempty"`
	CategoryID string `json:"category_id,omitempty"`

	//通过user_id关联用户
	User user.User `json:"user"`
	//通过category_id关联用户
	Category category.Category `json:"category"`

	models.CommonTimestampsField
}

func (topic *Topic) Create() {
	database.DB.Create(&topic)
}

func (topic *Topic) Save() (rowsAffected int64) {
	result := database.DB.Save(&topic)
	return result.RowsAffected
}
func (topic *Topic) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&topic)
	return result.RowsAffected
}
