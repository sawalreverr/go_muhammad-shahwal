package tests

import (
	"go-wishlist-api/auth"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
)

var (
	SECRET_KEY = []byte("PRIVATE_123")
	userID     = 1
)

func TestGenerateToken(t *testing.T) {
	service := auth.NewService()

	tokenString, err := service.GenerateToken(userID)

	assert.NoError(t, err, "failed to generate token")
	assert.NotEmpty(t, tokenString, "token string is empty")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return SECRET_KEY, nil
	})

	assert.NoError(t, err, "failed to parse generated token")
	assert.NotNil(t, token, "parsed token is nil")

	claims, ok := token.Claims.(jwt.MapClaims)
	assert.True(t, ok, "claims are not of type MapClaims")

	assert.Equal(t, userID, int(claims["user_id"].(float64)), "user ID in claims does not match")
}

func TestValidateTokenSuccess(t *testing.T) {
	service := auth.NewService()

	tokenString, err := service.GenerateToken(userID)
	assert.NoError(t, err)

	validToken, err := service.ValidateToken(tokenString)

	assert.NoError(t, err, "failed to validate valid token")
	assert.NotNil(t, validToken, "validated token is nil")
	assert.True(t, validToken.Valid, "token is not valid")
}

func TestValidateTokenInvalid(t *testing.T) {
	service := auth.NewService()

	invalidToken := "invalid_token_string"
	_, err := service.ValidateToken(invalidToken)

	assert.Error(t, err, "expected error for invalid token")
	assert.EqualError(t, err, "token is malformed: token contains an invalid number of segments")
}

func TestValidateTokenExpired(t *testing.T) {
	service := auth.NewService()

	expiredToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Second * -10).Unix(),
	})

	tokenString, err := expiredToken.SignedString(SECRET_KEY)
	assert.NoError(t, err)

	_, err = service.ValidateToken(tokenString)

	assert.Error(t, err, "expected error for expired token")
	assert.EqualError(t, err, "token has invalid claims: token is expired")
}
