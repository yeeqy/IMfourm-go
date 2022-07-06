// Package seed 处理数据库填充相关逻辑
package seed

import (
	"IMfourm-go/pkg/console"
	"IMfourm-go/pkg/database"
	"gorm.io/gorm"
)

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

// GetSeeder 通过名称来获取Seeder对象
func GetSeeder(name string) Seeder{
	for _, sdr := range seeders{
		if name == sdr.Name{
			return sdr
		}
	}
	return Seeder{}
}
//运行所有Seeder
func RunAll(){
	//先运行ordered的
	executed := make(map[string]string)
	for _,name := range orderedSeederNames{
		sdr := GetSeeder(name)
		if len(sdr.Name) > 0 {
			console.Warning("running order seeder: " + sdr.Name)
			sdr.Func(database.DB)
			executed[name] = name
		}
	}
	//再运行剩下的
	for _,sdr := range seeders{
		//过滤已运行
		if _,ok := executed[sdr.Name];!ok{
			console.Warning("running seeder: " + sdr.Name)
			sdr.Func(database.DB)
		}
	}
}

// RunSeeder 运行单个seeder
func RunSeeder(name string) {
	for _,sdr := range seeders{
		if name == sdr.Name{
			sdr.Func(database.DB)
			break
		}
	}
}