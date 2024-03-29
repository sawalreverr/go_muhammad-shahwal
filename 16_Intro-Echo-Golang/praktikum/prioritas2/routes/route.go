package routes

import (
	"net/http"
	"prioritas2_echo_golang/controllers"

	"github.com/labstack/echo/v4"
)

func InitRoute(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, echo.Map{
			"status": "on",
		})
	})

	e.GET("/words", controllers.GetAllWords)
	e.POST("/words", controllers.AddWord)
}
