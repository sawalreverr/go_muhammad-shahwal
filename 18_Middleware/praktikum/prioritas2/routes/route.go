package routes

import (
	"net/http"
	"time"

	"prioritas2_middleware/controllers"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var (
	configLogger = middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}

	configRateLimiter = middleware.RateLimiterConfig{
		Skipper: middleware.DefaultSkipper,
		Store:   middleware.NewRateLimiterMemoryStoreWithConfig(middleware.RateLimiterMemoryStoreConfig{Rate: 10, Burst: 30, ExpiresIn: 3 * time.Minute}),
		IdentifierExtractor: func(ctx echo.Context) (string, error) {
			return ctx.RealIP(), nil
		},
		ErrorHandler: func(context echo.Context, err error) error {
			return context.JSON(http.StatusForbidden, nil)
		},
		DenyHandler: func(context echo.Context, identifier string, err error) error {
			return context.JSON(http.StatusTooManyRequests, nil)
		},
	}
)

func InitRoutes(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, echo.Map{
			"status": "ON",
		})
	})

	// endpoint
	v1 := e.Group("/api/v1")

	// auth
	v1.POST("/register", controllers.UserRegisterController)
	v1.POST("/login", controllers.UserLoginController)

	// logger
	v1.Use(middleware.LoggerWithConfig(configLogger))

	v1Auth := v1

	// rate limiter
	v1Auth.Use(middleware.RateLimiterWithConfig(configRateLimiter))

	// jwt
	v1Auth.Use(echojwt.JWT([]byte("PRIVATE_123")))

	// post
	v1Auth.GET("/posts", controllers.GetAllPostController)
	v1Auth.GET("/posts/:id", controllers.GetPostByIdController)
	v1Auth.POST("/posts", controllers.CreateNewPostController)
	v1Auth.PUT("/posts/:id", controllers.EditPostByIdController)
	v1Auth.DELETE("/posts/:id", controllers.DeletePostByIdController)
}
