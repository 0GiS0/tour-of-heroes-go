package models

import (
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {

	database, err := gorm.Open(sqlserver.Open("sqlserver://sa:Password1!@localhost:1433?database=heroes"), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	err = database.AutoMigrate(&Hero{})

	if err != nil {
		return
	}

	DB = database
}
