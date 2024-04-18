package middlewares

import (
	"errors"
	"go-wishlist-api/auth"
	"go-wishlist-api/user"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func AuthMiddleware(authService auth.Service, userService user.Service) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")

			if !strings.Contains(authHeader, "Bearer") {
				return echo.ErrUnauthorized
			}

			tokenString := ""
			parts := strings.Split(authHeader, " ")
			if len(parts) == 2 {
				tokenString = parts[1]
			}

			token, err := authService.ValidateToken(tokenString)
			if err != nil {
				if errors.Is(err, jwt.ErrSignatureInvalid) {
					return echo.NewHTTPError(http.StatusUnauthorized, "invalid token signature")
				}
				return echo.ErrUnauthorized
			}

			claims, ok := token.Claims.(jwt.MapClaims)
			if !ok || !token.Valid {
				return echo.NewHTTPError(http.StatusUnauthorized, "invalid token")
			}

			userID := int(claims["user_id"].(float64))

			user, err := userService.GetUserByID(userID)
			if err != nil {
				return echo.NewHTTPError(http.StatusInternalServerError, "failed to get user")
			}

			c.Set("CurrentUser", user)
			return next(c)
		}
	}
}
