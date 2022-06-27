package auth

import (
	"IMfourm-go/app/models/user"
	"errors"
)

//授权相关逻辑


// 尝试登录
func Attempt(email string, password string)(user.User, error){
	userModel := user.GetByMulti(email)
	if userModel.ID == 0 {
		return user.User{},errors.New("账号不存在")
	}
	if !userModel.ComparePassword(password){
		return user.User{}, errors.New("密码错误")
	}
	return userModel, nil
}
//登录指定用户
func LoginByPhone(phone string)(user.User,error){
	userModel := user.GetByPhone(phone)
	if userModel.ID == 0 {
		return user.User{},errors.New("手机号未注册")
	}
	return userModel,nil
}
