package migrate

import (
	"IMfourm-go/pkg/console"
	"IMfourm-go/pkg/database"
	"IMfourm-go/pkg/file"
	"gorm.io/gorm"
	"io/ioutil"
)

// Migrator 数据迁移操作类
type Migrator struct {
	Folder   string
	DB       *gorm.DB
	Migrator gorm.Migrator
}

// Migration 对应数据的migrations表理的一条数据
type Migration struct {
	ID        uint64 `gorm:"primaryKey;autoIncrement;"`
	Migration string `gorm:"type:varchar(255);not null;unique;"`
	Batch     int
}

// NewMigrator Migrator实例，用以执行迁移操作
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

// Up 执行所有未迁移过的文件
func (migrator *Migrator) Up(){
	//读取所有迁移文件，确保按照时间排序
	migrateFiles := migrator.readAllMigrationFiles()
	//获取当前批次的值
	batch := migrator.getBatch()
	//获取所有迁移数据
	migrations := []Migration{}
	migrator.DB.Find(&migrations)
	//可以通过此值来判断数据库是否已是最新
	runed:= false
	//对比迁移文件进行遍历，如果没有执行过，就执行up回调
	for _,mfile := range migrateFiles{
		if mfile.isNotMigrated(migrations){
			migrator.runUpMigration(mfile,batch)
			runed = true
		}
	}
	if !runed{
		console.Success("database is up to date.")
	}
}

// Rollback 回滚上一个操作
func (migrator *Migrator) Rollback()  {
	//获取最后一批次的迁移数据
	lastMigration := Migration{}
	migrator.DB.Order("id DESC").First(&lastMigration)
	migrations := []Migration{}
	migrator.DB.Where("batch = ?",lastMigration.Batch).Order("id DESC").Find(&migrations)
	//回滚最后一批次的迁移
	if !migrator.rollbackMigrations(migrations){
		console.Success("[migrations] table is empty, nothing to rollback")
	}
}

func (migraotr *Migrator) rollbackMigrations(migrations []Migration) bool {
	//标记是否真的有执行
	runed := false
	for _, _migration := range migrations{
		console.Warning("rollback" + _migration.Migration)
		mfile := getMigrationFile(_migration.Migration)
		if mfile.Down != nil{
			mfile.Down(database.DB.Migrator(),database.SQLDB)
		}
		runed = true

		//回退成功就删除该记录
		migraotr.DB.Delete(&_migration)
		//打印运行状态
		console.Success("finish " + mfile.FileName)
	}
	return runed
}

//获取当前批次值
func (migrator *Migrator) getBatch() int {
	//默认1
	batch := 1
	lastMigration := Migration{}
	migrator.DB.Order("id DESC").First(&lastMigration)

	if lastMigration.ID > 0 {
		batch = lastMigration.Batch + 1
	}
	return batch
}

//从文件目录读取文件，保证正确的时间排序
func (migrator *Migrator) readAllMigrationFiles() []MigrationFile {
	//读取database/migrations/目录下的所有文件
	//默认是会按照文件名称进行排序
	files,err := ioutil.ReadDir(migrator.Folder)
	console.ExitIf(err)

	var migrateFiles []MigrationFile
	for _,f := range files{
		//去除文件后缀.go
		fileName := file.FileNameWithoutExtension(f.Name())

		//通过迁移文件的名称获取MigrationFile对象
		mfile := getMigrationFile(fileName)

		//加个判断，确保迁移文件可用，再放入migrateFiles数组中
		if len(mfile.FileName) > 0 {
			migrateFiles = append(migrationFiles,mfile)
		}
	}
	//返回排序好的MigrationFile数组
	return migrateFiles
}
//执行迁移，执行迁移的up方法
func (migrator *Migrator) runUpMigration(mfile MigrationFile,batch int){
	//执行 up区块的 SQL
	if mfile.Up!=nil{
		console.Warning("migrating " + mfile.FileName)
		//**
		mfile.Up(database.DB.Migrator(),database.SQLDB)

		console.Success("migrated " + mfile.FileName)
	}
	//入库
	err := migrator.DB.Create(&Migration{
		Migration: mfile.FileName,
		Batch:     batch,
	}).Error
	console.ExitIf(err)
}

//回滚所有迁移
func (migrator *Migrator) Reset(){
	migrations := []Migration{}
	//按照倒序读取所有迁移文件
	migrator.DB.Order("id DESC").Find(&migrations)
	//回滚所有迁移
	if !migrator.rollbackMigrations(migrations){
		console.Success("[migrations] table is empty, nothing to reset.")
	}
}

//回滚所有迁移，并运行所有迁移
func (migrator *Migrator) Refresh()  {
	//回滚所有迁移
	migrator.Reset()
	//再次执行所有迁移
	migrator.Up()
}