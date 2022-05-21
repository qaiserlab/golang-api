package auth

import (
	"crypto/sha1"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
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
// @Param        	user body LoginForm true "Form Data"
// @Success      	200 {object} LoginResponse
// @Router       	/v1/auth/login [post]
func Login(c *gin.Context) {
	var formData LoginForm
	var user models.User
	var loginResponse LoginResponse

	JWT_KEY := []byte(os.Getenv("JWT_KEY"))
	JWT_EXP, err := time.ParseDuration(os.Getenv("JWT_EXP"))

	db := c.MustGet("db").(*gorm.DB)

	if err := c.ShouldBindJSON(&formData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Preload("Role").Where("username = ?", formData.Username).First(&user).Error; err != nil {
		if err := db.Where("email = ?", formData.Username).First(&user).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found."})
			return
		}
	}

	salt := user.Salt

	sha := sha1.New()
	sha.Write([]byte(salt + formData.Password))
	password := fmt.Sprintf("%x", sha.Sum(nil))

	if password != user.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid user password."})
		return
	}

	// Session expire on 3 hour
	expirationTime := time.Now().Add(JWT_EXP)

	claims := &Claims{
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Gender:      user.Gender,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		Username:    user.Username,
		Role:        user.Role,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenValue, err := token.SignedString(JWT_KEY)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	domain := os.Getenv("HOST") + ":" + os.Getenv("PORT")
	c.SetCookie("token", tokenValue, 10, "/", domain, true, true)

	loginResponse.AccessToken = tokenValue
	loginResponse.RefreshToken = tokenValue

	c.JSON(http.StatusOK, gin.H{"data": loginResponse})
}
