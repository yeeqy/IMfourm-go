package user

import (
	"IMfourm-go/pkg/app"
	"IMfourm-go/pkg/database"
	"IMfourm-go/pkg/paginator"
	"github.com/gin-gonic/gin"
)

//存放模型相关的数据库操作

// IsEmailExist 邮箱是否已注册
func IsEmailExist(email string)bool  {
	var count int64
	database.DB.Model(User{}).Where("email=?",email).Count(&count)
	return count > 0
}

// IsPhoneExist 手机号是否已注册
func IsPhoneExist(phone string) bool{
	var count int64
	database.DB.Model(User{}).Where("phone=?",phone).Count(&count)
	return count > 0
}

// GetByPhone 通过手机号获取用户
func GetByPhone(phone string)(userModel User){
	database.DB.Where("phone = ?",phone).First(&userModel)
	return
}

// GetByMulti 通过 手机号/Email/用户名 获取用户
func GetByMulti(loginID string)(userModel User){
	database.DB.Where("phone = ?",loginID).Or("email = ?",loginID).Or("name = ?",loginID).First(&userModel)
	return
}

func Get(idstr string)(userModel User){
	database.DB.Where("id",idstr).First(&userModel)
	return
}

func GetByEmail(email string)(userModel User){
	database.DB.Where("email = ?",email).First(&userModel)
	return
}

func All()(user []User)  {
	database.DB.Find(&user)
	return
}

// Paginate 分页内容
func Paginate(c *gin.Context,perPage int)(users []User,paging paginator.Paging){
	paging = paginator.Paginate(
		c, database.DB.Model(User{}), &users,
		app.V1URL(database.TableName(&User{})),
		perPage,
		)
	return
}