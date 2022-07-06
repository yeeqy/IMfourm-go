package seeders

import (
	"IMfourm-go/database/factories"
	"IMfourm-go/pkg/console"
	"IMfourm-go/pkg/logger"
	"IMfourm-go/pkg/seed"
	"fmt"
	"gorm.io/gorm"
)

func init(){

	//添加seeder
	seed.Add("SeedUsersTable", func(db *gorm.DB) {

		//创建10个用户对象
		users := factories.MakeUsers(10)
		//批量创建用户
		result := db.Table("users").Create(&users)

		if err := result.Error; err != nil{
			logger.LogIf(err)
			return
		}
		//打印运行情况
		console.Success(fmt.Sprintf("Table [%v] %v rows seeded",result.Statement.Table,result.RowsAffected))
	})
}
