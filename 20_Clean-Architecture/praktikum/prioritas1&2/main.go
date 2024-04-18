package main

import (
	"go-wishlist-api/routes"
	"go-wishlist-api/utils/database"
	"go-wishlist-api/utils/migration"
	"net/http"

	"github.com/labstack/echo/v4"
)

func init() {
	// connect database
	database.InitDB()

	// migrate models
	migration.AutoMigrate()
}

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	routes.InitRoute(e)
	e.Logger.Fatal(e.Start(":1323"))
}
