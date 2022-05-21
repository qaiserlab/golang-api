package auth

import (
	"github.com/dgrijalva/jwt-go"
	"golang.api/models"
)

type Claims struct {
	ID          int `json:"id"`
	RoleID      int
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Gender      int    `json:"gender"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
	Username    string `json:"username"`
	Role        models.Role
	jwt.StandardClaims
}
