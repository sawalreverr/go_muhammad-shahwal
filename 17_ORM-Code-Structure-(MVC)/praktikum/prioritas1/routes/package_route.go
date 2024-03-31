package routes

import (
	"net/http"
	"prioritas1_orm_mvc/controllers"

	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, echo.Map{
			"status": "on",
		})
	})

	e.GET("/api/v1/packages", controllers.GetAllPackagesController)
	e.GET("/api/v1/packages/:id", controllers.GetPackageByIdController)
	e.POST("/api/v1/packages", controllers.CreatePackageController)
	e.PUT("/api/v1/packages/:id", controllers.EditPackageByIdController)
	e.DELETE("/api/v1/packages/:id", controllers.DeletePackageByIdController)
}
