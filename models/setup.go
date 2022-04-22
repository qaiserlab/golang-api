package models

import (
	"crypto/sha1"
	"fmt"
	"os"
	"time"

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

	if err := db.AutoMigrate(
		&User{},
		&Role{},
	); err != nil {
		db.Where(
			Role{Name: "administrator"}).
			Assign(Role{
				ID:   1,
				Name: "Administrator",
			}).
			FirstOrCreate(&Role{})

		salt := fmt.Sprintf("%d", time.Now().UnixNano())

		sha := sha1.New()
		sha.Write([]byte(salt + "admin"))
		password := fmt.Sprintf("%x", sha.Sum(nil))

		db.Where(
			User{Username: "admin"}).
			Assign(User{
				RoleID:      1,
				Name:        "Fadlun Anaturdasa Wibawa",
				Gender:      0,
				Email:       "f.anaturdasa@gmail.com",
				PhoneNumber: "-",
				Username:    "admin",
				Password:    password,
				Salt:        salt,
			}).
			FirstOrCreate(&User{})
	}

	return db
}
