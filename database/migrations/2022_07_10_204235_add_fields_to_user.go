package migrations

import (
	"IMfourm-go/pkg/migrate"
	"database/sql"

	"gorm.io/gorm"
)

func init() {

	type User struct {
		City         string `gorm:"type:varchar(10);"`
		Introduction string `gorm:"type:varchar(255)"`
		Avatar       string `gorm:"type:varchar(255);default:null"`
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.AutoMigrate(&User{})
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.DropTable(&User{},"City")
		migrator.DropTable(&User{},"Introduction")
		migrator.DropTable(&User{},"Avatar")
	}

	migrate.Add("2022_07_10_204235_add_fields_to_user", up, down)
}
