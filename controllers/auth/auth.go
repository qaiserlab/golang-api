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
// @Param        	user body FormLogin true "Form Data"
// @Success      	200 {object} LoginResponse
// @Router       	/v1/auth/login [post]
func Login(c *gin.Context) {
	var formData FormLogin
	var user models.User
	var loginResponse LoginResponse

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

	if password != user.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid user password."})
		return
	}

	JWT_KEY := []byte(os.Getenv("JWT_KEY"))

	// Session expire on 3 hour
	expirationTime := time.Now().Add(180 * time.Minute)

	claims := &Claims{
		Username: formData.Username,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(JWT_KEY)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// Finally, we set the client cookie for "token" as the JWT we just generated
	// we also set an expiry time which is the same as the token itself
	http.SetCookie(c.Writer, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})

	loginResponse.Username = user.Username
	loginResponse.Token = tokenString

	c.JSON(http.StatusOK, gin.H{"data": loginResponse})
}
