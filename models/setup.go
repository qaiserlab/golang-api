package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func SetupModels() *gorm.DB {
	db, err := gorm.Open("mysql", "amd:m30ng@(localhost)/db_golang")

	if err != nil {
		// panic("Database connection failed!")
		panic(err)
	}

	db.AutoMigrate(&User{}, &Role{})

	return db
}
