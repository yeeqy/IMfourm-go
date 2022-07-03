package migrate

import (
	"database/sql"
	"gorm.io/gorm"
)

//定义up和down回调方法的类型
type migrationFunc func(migrator gorm.Migrator, db *sql.DB)

//所有的迁移文件数组
var migrationFiles []MigrationFile

//单个迁移文件
type MigrationFile struct {
	Up       migrationFunc
	Down     migrationFunc
	fileName string
}
//新增一个迁移文件，所有的迁移文件都需要调用此方法来注册
func Add(name string,up migrationFunc, down migrationFunc) {
	migrationFiles = append(migrationFiles,MigrationFile{
		Up:       up,
		Down:     down,
		fileName: name,
	})
}