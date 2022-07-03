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
	FileName string
}
//新增一个迁移文件，所有的迁移文件都需要调用此方法来注册
func Add(name string,up migrationFunc, down migrationFunc) {
	migrationFiles = append(migrationFiles,MigrationFile{
		Up:       up,
		Down:     down,
		FileName: name,
	})
}
//通过迁移文件的名称来获取到MigrationFile对象
func getMigrationFile(name string) MigrationFile{
	for _,mfile := range migrationFiles{
		if name == mfile.FileName{
			return mfile
		}
	}
	return MigrationFile{}
}
//判断迁移是否已执行
func (mfile MigrationFile) isNotMigrated(migrations []Migration) bool {
	for _,migration := range migrations{
		if migration.Migration == mfile.FileName{
			return false
		}
	}
	return true
}