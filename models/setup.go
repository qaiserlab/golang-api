package models

import (
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func SetupModels() *gorm.DB {
	DB_DRIVER := os.Getenv("DB_DRIVER")
	DB_USER := os.Getenv("DB_USER")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")
	DB_HOST := os.Getenv("DB_HOST")
	DB_NAME := os.Getenv("DB_NAME")

	db, err := gorm.Open(DB_DRIVER, DB_USER+":"+DB_PASSWORD+"@("+DB_HOST+")/"+DB_NAME)

	if err != nil {
		// panic("Database connection failed!")
		panic(err)
	}

	db.AutoMigrate(&User{}, &Role{})

	return db
}
