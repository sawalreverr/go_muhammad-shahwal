package main

import (
	"prioritas1_orm_mvc/configs"
	"prioritas1_orm_mvc/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	configs.ConnectDatabase()

	e := echo.New()

	routes.InitRoutes(e)
	e.Logger.Fatal(e.Start(":8080"))
}
