package auth

import (
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Cookie("token")
		errorResponse := gin.H{"error": "Access denied."}

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse)
			return
		}

		JWT_KEY := []byte(os.Getenv("JWT_KEY"))
		claims := &Claims{}

		jwtToken, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
			return JWT_KEY, nil
		})

		if err != nil || !jwtToken.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, errorResponse)
			return
		}

		c.Next()
	}
}
