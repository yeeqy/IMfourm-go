package user

import (
	"IMfourm-go/app/models"
	"IMfourm-go/pkg/database"
	"IMfourm-go/pkg/hash"
)

type User struct {
	models.BaseModel

	Name string `json:"name,omitempty"`

	City         string `json:"city,omitempty"`
	Introduction string `json:"introduction,omitempty"`
	Avatar       string `json:"avatar,omitempty"`

	Email    string `json:"-"`
	Phone    string `json:"-"`
	Password string `json:"-"`
	//指示 JSON 解析器忽略字段
	//后面接口返回用户数据时候，这三个字段都会被隐藏
	models.CommonTimestampsField
}

//创建用户
func (userModel *User) Create() {
	database.DB.Create(&userModel)
}

//匹配密码
func (userModel *User) ComparePassword(_password string) bool {
	return hash.BcryptCheck(_password, userModel.Password)
}

func (userModel *User) Save() (rowsAffected int64) {
	result := database.DB.Save(&userModel)
	return result.RowsAffected
}
