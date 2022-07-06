// Package factories 存放工厂方法
package factories

import (
	"IMfourm-go/app/models/user"
	"IMfourm-go/pkg/helpers"
	"github.com/bxcodec/faker/v3"
)

//Seeder 功能是往数据库里填充假数据，方便列表接口有数据可用。
//factory —— 数据工厂，用来批量生成模型对象，且使用假数据对这些模型对象的属性进行赋值；
//seeder —— 负责将数据工厂里生成的对象插入到数据库中。

// MakeUsers 创建用户对象
func MakeUsers(times int) []user.User{
	var obj []user.User

	//设置唯一值
	faker.SetGenerateUniqueValues(true)

	for i:= 0; i < times; i++{
		model := user.User{
			Name: faker.Username(),
			Email: faker.Email(),
			Phone: helpers.RandomNumber(11),
			Password:"a/test/password",
		}
		obj = append(obj,model)
	}
	return obj
}