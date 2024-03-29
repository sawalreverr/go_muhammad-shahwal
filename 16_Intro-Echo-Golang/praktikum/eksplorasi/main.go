package main

import (
	"eksplorasi_echo_golang/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	routes.InitRoute(e)

	e.Logger.Fatal(e.Start(":8080"))
}
