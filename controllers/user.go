package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.api/models"
)

type UserRegister struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
	Username    string `json:"username"`
	Password    string `json:"password"`
}

func GetAll(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var users []models.User
	db.Find(&users)

	c.JSON(http.StatusOK, gin.H{"data": users})
}
