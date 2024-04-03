package routes

import (
	"net/http"
	"prioritas1_middleware/controllers"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitRoutes(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, echo.Map{
			"status": "ON",
		})
	})

	v1 := e.Group("/api/v1")
	v1.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	// AUTH
	v1.POST("/register", controllers.UserRegisterController)
	v1.POST("/login", controllers.UserLoginController)

	v1Auth := v1
	v1Auth.Use(echojwt.JWT([]byte("PRIVATE_123")))

	// PACKAGES
	v1Auth.GET("/packages", controllers.GetAllPackagesController)
	v1Auth.GET("/packages/:id", controllers.GetPackageByIdController)
	v1Auth.POST("/packages", controllers.CreatePackageController)
	v1Auth.PUT("/packages/:id", controllers.EditPackageByIdController)
	v1Auth.DELETE("/packages/:id", controllers.DeletePackageByIdController)
}
