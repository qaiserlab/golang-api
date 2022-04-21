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
	var formData UserRegister

	db := c.MustGet("db").(*gorm.DB)

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

func Update(c *gin.Context) {
	var modelData models.User
	var formData UserRegister

	db := c.MustGet("db").(*gorm.DB)

	if err := db.Where("id = ?", c.Param("id")).First(&modelData).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.ShouldBindJSON(&formData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Model(&modelData).Update(formData)

	c.JSON(http.StatusOK, gin.H{"data": modelData})
}

func GetAll(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var users []models.User
	db.Find(&users)

	c.JSON(http.StatusOK, gin.H{"data": users})
}
