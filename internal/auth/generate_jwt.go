package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(user string, email string) (string, error) {
	claims := Claims{
		Username: user,
		Email:    email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte("your_secret_key"))
}
