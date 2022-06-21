package user

import "IMfourm-go/app/models"

type User struct {
	models.BaseModel

	Name string `json:"name,omitempty"`
	Email string `json:"-"`
	Phone string `json:"-"`
	Password string `json:"-"`
	//指示 JSON 解析器忽略字段
	//后面接口返回用户数据时候，这三个字段都会被隐藏
	models.CommonTimestampsField
}
