package user

import (
	"IMfourm-go/pkg/hash"
	"gorm.io/gorm"
)

//gorm提供beforeSave的模型钩子，会在模型创新和更新前被调用
//利用此机制在入库前对密码做加密

func (userModel *User) BeforeSave(tx *gorm.DB) (err error) {
	if !hash.BcryptHashed(userModel.Password){
		userModel.Password = hash.BcryptHash(userModel.Password)
	}
	return
}
