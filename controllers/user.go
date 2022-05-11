package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.api/models"
)

type FormUser struct {
	Name        string `json:"name" binding:"required"`
	Email       string `json:"email" binding:"required"`
	PhoneNumber string `json:"phoneNumber"`
	Username    string `json:"username" binding:"required"`
	Password    string `json:"password" binding:"required"`
}

// GetUsers 			godoc
// @Summary				List user
// @Description		Get list of users
// @Tags         	user
// @Accept       	json
// @Produce      	json
// @Success      	200 {array} models.User
// @Router       	/users [get]
func GetUsers(c *gin.Context) {
	var users []models.User

	db := c.MustGet("db").(*gorm.DB)
	db.Preload("Role").Find(&users)

	c.JSON(http.StatusOK, gin.H{"data": users})
}

// GetUserById		godoc
// @Summary     	Get user
// @Description  	Get one user data by ID
// @Tags         	user
// @Accept       	json
// @Produce      	json
// @Param       	id path int true "User ID"
// @Success     	200 {object} models.User
// @Router      	/users/{id} [get]
func GetUserById(c *gin.Context) {
	var user models.User

	db := c.MustGet("db").(*gorm.DB)
	db.Preload("Role").Find(&user, c.Param("id"))

	c.JSON(http.StatusOK, gin.H{"data": user})
}

// CreateUser			godoc
// @Summary     	Create user
// @Description  	Create new user
// @Tags         	user
// @Accept       	json
// @Produce      	json
// @Param       	user body FormUser true "Form Data"
// @Success     	201 {object} models.User
// @Router      	/users [post]
func CreateUser(c *gin.Context) {
	var formData FormUser
	var user models.User

	db := c.MustGet("db").(*gorm.DB)

	if err := c.ShouldBindJSON(&formData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Where("email = ?", formData.Email).First(&user).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email already registered"})
		return
	}

	if err := db.Where("username = ?", formData.Username).First(&user).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username already registered"})
		return
	}

	newUser := models.User{
		Name:        formData.Name,
		Email:       formData.Email,
		PhoneNumber: formData.PhoneNumber,
		Username:    formData.Username,
		Password:    formData.Password,
	}

	db.Create(&newUser)

	c.JSON(http.StatusCreated, gin.H{"data": newUser})
}

// UpdateUserById	godoc
// @Summary     	Update user
// @Description  	Update user data by ID
// @Tags         	user
// @Accept       	json
// @Produce      	json
// @Param       	id path int true "User ID"
// @Param       	user body FormUser true "Form Data"
// @Success     	200 {object} models.User
// @Router      	/users/{id} [put]
func UpdateUserById(c *gin.Context) {
	var formData FormUser
	var user models.User

	db := c.MustGet("db").(*gorm.DB)

	if err := db.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.ShouldBindJSON(&formData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Model(&user).Update(formData)

	c.JSON(http.StatusOK, gin.H{"data": user})
}

func DeleteUserById(c *gin.Context) {
	var user models.User

	db := c.MustGet("db").(*gorm.DB)

	if err := db.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Delete(&user)

	c.JSON(http.StatusOK, gin.H{"data": user})
}
