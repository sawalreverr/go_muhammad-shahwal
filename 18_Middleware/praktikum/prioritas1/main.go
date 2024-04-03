package main

import (
	"prioritas1_middleware/configs"
	"prioritas1_middleware/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	configs.ConnectDatabase()

	routes.InitRoutes(e)

	e.Logger.Fatal(e.Start(":8080"))
}
