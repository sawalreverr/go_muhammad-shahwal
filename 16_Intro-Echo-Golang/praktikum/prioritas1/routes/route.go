package routes

import (
	"net/http"
	"prioritas1_echo_golang/controllers"

	"github.com/labstack/echo/v4"
)

func InitRoute(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Status: ON")
	})

	e.GET("/api/v1/foods", controllers.GetAllFoodsController)
	e.GET("/api/v1/foods/:id", controllers.GetFoodByIdController)
	e.POST("/api/v1/foods", controllers.CreateFoodController)
	e.PUT("/api/v1/foods/:id", controllers.EditFoodController)
	e.DELETE("/api/v1/foods/:id", controllers.DeleteFoodController)
}
