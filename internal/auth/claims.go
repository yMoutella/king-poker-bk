package auth

import (
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	Username string `json:"user"`
	Email    string `json:"email"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}
