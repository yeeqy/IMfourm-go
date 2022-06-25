package hash

import (
	"IMfourm-go/pkg/logger"
	"golang.org/x/crypto/bcrypt"
)

//用户密码是明文存储于数据库的，我们需要对其进行哈希后再存入，以保证用户密码的安全

//使用bcrypt对密码进行加密
func BcryptHash(password string) string{
	bytes,err := bcrypt.GenerateFromPassword([]byte(password),14)
	logger.LogIf(err)
	return string(bytes)
}
//对比明文密码和数据库的哈希值
func BcryptCheck(password,hash string)bool{
	err := bcrypt.CompareHashAndPassword([]byte(hash),[]byte(password))
	return err == nil
}
//判断字符串是否哈希过
func BcryptHashed(str string)bool{
	// bcrypt加密后的长度=60
	return len(str) == 60
}
