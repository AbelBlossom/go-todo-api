package db

import (
	"fmt"

	"github.com/abelblossom/todo/src/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectToDataBase() (*gorm.DB, error) {
	// dsn := "host=mel.db.elephantsql.com user=bbbrnfkp password=DgA2fZeVyt18e9l8hBs9UW0TxD9SmNr9 dbname=bbbrnfkp port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(sqlite.Open("data.db"), &gorm.Config{
		Logger: logger.Default,
	})
	if err == nil {
		DB = db
		err := db.AutoMigrate(&models.Todo{}, &models.User{})
		if err == nil {
			fmt.Println("Migrated")
		} else {
			fmt.Println("Fail to Migrate", err)
			// panic(err)
		}
	} else {
		fmt.Println("Fail to Connect")
	}

	return db, err
}
