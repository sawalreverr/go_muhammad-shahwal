package routes

import (
	"net/http"
	"prioritas2_orm_mvc/controllers"

	"github.com/labstack/echo/v4"
)

func InitRoute(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, echo.Map{
			"status": "on",
		})
	})

	e.GET("/api/v1/fruits", controllers.GetAllFruitsController)
	e.GET("/api/v1/fruits/:id", controllers.GetFruitByIdController)
	e.POST("/api/v1/fruits", controllers.CreateFruitController)
	e.DELETE("/api/v1/fruits/:id", controllers.DeleteFruitController)
}
