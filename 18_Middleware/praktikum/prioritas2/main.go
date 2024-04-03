package main

import (
	"prioritas2_middleware/configs"
	"prioritas2_middleware/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	configs.ConnectDatabase()
	e := echo.New()

	routes.InitRoutes(e)
	e.Logger.Fatal(e.Start(":8080"))
}
