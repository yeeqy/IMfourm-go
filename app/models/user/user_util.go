package user

import "IMfourm-go/pkg/database"

//存放模型相关的数据库操作

//邮箱是否已注册
func IsEmailExist(email string)bool  {
	var count int64
	database.DB.Model(User{}).Where("email=?",email).Count(&count)
	return count > 0
}
//手机号是否已注册
func IsPhoneExist(phone string) bool{
	var count int64
	database.DB.Model(User{}).Where("phone=?",phone).Count(&count)
	return count > 0
}
