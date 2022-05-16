package auth

import (
	"crypto/sha1"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.api/models"
)

// Login         	godoc
// @Summary      	Login
// @Description  	Login as user
// @Tags         	auth
// @Accept       	json
// @Produce      	json
// @Param        	user body FormLogin true "Form Data"
// @Success      	200 {object} models.User
// @Router       	/v1/auth/login [post]
func Login(c *gin.Context) {
	var formData FormLogin
	var user models.User

	db := c.MustGet("db").(*gorm.DB)

	if err := c.ShouldBindJSON(&formData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Where("username = ?", formData.Username).First(&user).Error; err != nil {
		if err := db.Where("email = ?", formData.Username).First(&user).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found."})
			return
		}
	}

	salt := user.Salt

	sha := sha1.New()
	sha.Write([]byte(salt + formData.Password))
	password := fmt.Sprintf("%x", sha.Sum(nil))

	if formData.Password != password {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid user password."})
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}
