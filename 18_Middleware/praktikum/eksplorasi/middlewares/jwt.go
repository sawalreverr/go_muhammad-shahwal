package middlewares

import (
	"eksplorasi_middleware/constants"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

type jwtCustomClaims struct {
	Email string `json:"user_email"`
	Role  string `json:"user_role"`
	jwt.RegisteredClaims
}

func GenerateTokenJWT(userEmail string, userRole string) string {
	claims := &jwtCustomClaims{
		userEmail, userRole,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	resultJWT, _ := token.SignedString([]byte(constants.SECRET_JWT))
	return resultJWT
}

func NewJWTMiddleware() echo.MiddlewareFunc {
	configJWT := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(jwtCustomClaims)
		},
		SigningKey: []byte(constants.SECRET_JWT),
	}

	return echojwt.WithConfig(configJWT)
}

func UserValidate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		claims := c.Get("user").(*jwt.Token).Claims.(*jwtCustomClaims)
		fmt.Println(claims)
		if claims.Role != "user" && claims.Role != "admin" {
			return echo.ErrUnauthorized
		}
		return next(c)
	}
}

func AdminValidate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		claims := c.Get("user").(*jwt.Token).Claims.(*jwtCustomClaims)
		fmt.Println(claims)
		if claims.Role != "admin" {
			return echo.ErrUnauthorized
		}
		return next(c)
	}

}
