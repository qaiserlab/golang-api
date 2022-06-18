package auth

import (
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.api/helpers"
	"golang.api/models"
	"golang.api/types"
)

// Login         	godoc
// @Summary      	Login
// @Description  	Login as user
// @Tags         	auth
// @Accept       	json
// @Produce      	json
// @Param        	user body LoginForm true "Form Data"
// @Success      	200 {object} AuthResponse
// @Router       	/v1/auth/login [post]
func Login(c *gin.Context) {
	var formData LoginForm
	var user models.User
	var authResponse AuthResponse

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

	password := helpers.GenHash(formData.Password, user.Salt)

	if password != user.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid user password."})
		return
	}

	expirationTime := time.Now().Add(JWT_EXP)

	claims := &types.AuthClaims{
		ID:          user.ID,
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

	// Generate & Save Access Token

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessTokenValue, err := accessToken.SignedString(JWT_KEY)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// Generate & Save Refresh Token

	refreshToken := jwt.New(jwt.SigningMethodHS256)
	refreshTokenClaims := refreshToken.Claims.(jwt.MapClaims)
	// Not sure about this sub & exp
	refreshTokenClaims["sub"] = 1
	// refreshTokenClaims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	refreshTokenClaims["exp"] = expirationTime.Add(time.Hour * 24).Unix()

	refreshTokenValue, err := refreshToken.SignedString(JWT_KEY)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	domain := os.Getenv("HOST") + ":" + os.Getenv("PORT")
	c.SetCookie("token", accessTokenValue, int(expirationTime.Unix()), "/", domain, true, true)

	authResponse.AccessToken = accessTokenValue
	authResponse.RefreshToken = refreshTokenValue

	c.JSON(http.StatusOK, gin.H{"data": authResponse})
}

// Refresh       	godoc
// @Summary      	Refresh
// @Description  	Refresh authorization token
// @Tags         	auth
// @Accept       	json
// @Produce      	json
// @Param       	token path string true "Refresh Token"
// @Success      	200 {object} AuthResponse
// @Router       	/v1/auth/refresh/{token} [get]
func Refresh(c *gin.Context) {
	refreshToken := c.Param("token")
	JWT_KEY := []byte(os.Getenv("JWT_KEY"))

	jwtToken, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		return JWT_KEY, nil
	})

	if err != nil || !jwtToken.Valid {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Access denied."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "Success."})
}
