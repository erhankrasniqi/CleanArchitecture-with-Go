package db

import (
	"fmt"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var database *gorm.DB

func ConnectToDB() {
	var err error

	database, err = gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Cannot connect with database")
	}
	fmt.Println("Connected with database success")
}

func GetDB() *gorm.DB {
	return database
}
