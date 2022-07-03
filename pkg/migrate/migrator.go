package migrate

import (
	"IMfourm-go/pkg/database"
	"gorm.io/gorm"
)

//数据迁移操作类
type Migrator struct {
	Folder   string
	DB       *gorm.DB
	Migrator gorm.Migrator
}

//对应数据的migrations表理的一条数据
type Migration struct {
	ID        uint64 `gorm:"primaryKey;autoIncrement;"`
	Migration string `gorm:"type:varchar(255);not null;unique;"`
	Batch     int
}

//船舰Migrator实例，用以执行迁移操作
func NewMigrator() *Migrator{
	//初始化必要属性
	migrator := &Migrator{
		Folder:   "database/migrations/",
		DB:       database.DB,
		Migrator: database.DB.Migrator(),
	}
	//不存在的话就创建它
	migrator.createMigrationsTable()
	return migrator
}
//创建migrations表
func(migrator *Migrator) createMigrationsTable(){
	migration := Migration{}
	//不存在才创建
	if !migrator.Migrator.HasTable(&migration){
		migrator.Migrator.CreateTable(&migration)
	}
}
