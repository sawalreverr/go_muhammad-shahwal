package middlewares

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type jwtCustomClaims struct {
	UserID    int    `json:"user_id"`
	UserEmail string `json:"user_email"`
	jwt.RegisteredClaims
}

func GenerateTokenJWT(userId int, userEmail string) string {
	userClaims := &jwtCustomClaims{
		userId, userEmail,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, userClaims)

	resultJWT, _ := token.SignedString([]byte("PRIVATE_123"))
	return resultJWT
}
