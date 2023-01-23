package models

import (
	"os"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {

	database, err := gorm.Open(sqlserver.Open(os.Getenv("DB_CONNECTION_STRING")), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	err = database.AutoMigrate(&Hero{})

	if err != nil {
		return
	}

	DB = database
}
