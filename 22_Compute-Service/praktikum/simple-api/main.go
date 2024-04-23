package main

import (
	"net/http"
	"simple-api/database"

	"github.com/labstack/echo/v4"
)

func main() {
	database.InitialDatabase()

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.Logger.Fatal(e.Start(":8080"))
}
