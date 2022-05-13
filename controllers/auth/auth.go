package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.api/models"
)

// Login					godoc
// @Summary     	Login
// @Description  	Login as user
// @Tags         	auth
// @Accept       	json
// @Produce      	json
// @Param       	user body FormLogin true "Form Data"
// @Success     	200 {object} models.User
// @Router      	/auth/login [post]
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

	c.JSON(http.StatusOK, gin.H{"data": user})
}
