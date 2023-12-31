package models

import (
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)

var DB *gorm.DB

func ConnectDatabase () {
	
	database, err := gorm.Open(mysql.Open("root:root@tcp(localhost:8889)/gormauth"))

	if err != nil {
		panic(err)
	}

	database.AutoMigrate(&User{})

	DB = database
}