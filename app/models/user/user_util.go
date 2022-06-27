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

//通过手机号获取用户
func GetByPhone(phone string)(userModel User){
	database.DB.Where("phone = ?",phone).First(&userModel)
	return
}
//通过 手机号/Email/用户名 获取用户
func GetByMulti(loginID string)(userModel User){
	database.DB.Where("phone = ?",loginID).Or("email = ?",loginID).Or("name = ?",loginID).First(&userModel)
	return
}