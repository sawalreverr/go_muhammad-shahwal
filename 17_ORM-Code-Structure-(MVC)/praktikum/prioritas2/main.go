package main

import (
	"prioritas2_orm_mvc/configs"
	"prioritas2_orm_mvc/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	configs.ConnectDatabase()

	e := echo.New()
	routes.InitRoute(e)

	e.Logger.Fatal(e.Start(":8080"))
}
