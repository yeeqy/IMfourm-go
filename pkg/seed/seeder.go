// Package seed 处理数据库填充相关逻辑
package seed

import "gorm.io/gorm"

//按顺序执行的Seeder数组
var orderedSeederNames []string

//存放所有seeder
var seeders []Seeder

type SeederFunc func(db *gorm.DB)

// Seeder 对应每一个 database/seeders 目录下的 Seeder 文件
type Seeder struct {
	Func SeederFunc
	Name string
}

//Add 注册到seeders数组中
func Add(name string, fn SeederFunc)  {
	seeders = append(seeders,Seeder{
		Name: name,
		Func: fn,
	})
}
//SetRunOrder 设置 按顺序执行的Seeder数组
func SetRunOrder(names []string){
	orderedSeederNames = names
}