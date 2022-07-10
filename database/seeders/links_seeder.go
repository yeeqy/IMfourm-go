package seeders

import (
    "IMfourm-go/database/factories"
    "IMfourm-go/pkg/console"
    "IMfourm-go/pkg/logger"
    "IMfourm-go/pkg/seed"
    "fmt"

    "gorm.io/gorm"
)

func init() {

    seed.Add("SeedLinksTable", func(db *gorm.DB) {

        links  := factories.MakeLinks(5)

        result := db.Table("links").Create(&links)

        if err := result.Error; err != nil {
            logger.LogIf(err)
            return
        }

        console.Success(fmt.Sprintf("Table [%v] %v rows seeded", result.Statement.Table, result.RowsAffected))
    })
}