package auth

import "github.com/dgrijalva/jwt-go"

type FormLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	Username string `json:"username"`
	Token    string `json:"token"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}
