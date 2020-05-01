package domain

import (
	"github.com/dgrijalva/jwt-go"
)

//Token auth token
type Token struct {
	UserID uint
	jwt.StandardClaims
}

//User account info
type User struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Token    string `json:"token"`
}

//Ward ward info
type Ward struct {
	Name string
	Code string
}
