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

func Register(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var formData UserRegister

	if err := c.ShouldBindJSON(&formData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	modelData := models.User{
		Name:        formData.Name,
		Email:       formData.Email,
		PhoneNumber: formData.PhoneNumber,
		Username:    formData.Username,
		Password:    formData.Password,
	}

	db.Create(&modelData)

	c.JSON(http.StatusOK, gin.H{"data": modelData})
}

func GetAll(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var users []models.User
	db.Find(&users)

	c.JSON(http.StatusOK, gin.H{"data": users})
}
