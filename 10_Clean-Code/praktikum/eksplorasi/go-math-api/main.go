package main

import (
	"praktikum/eksplorasi/go-math-api/controllers"
	"praktikum/eksplorasi/go-math-api/models"
	"praktikum/eksplorasi/go-math-api/routers"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	// init model
	opModel := models.NewOperationModel()

	// init controller
	opController := controllers.NewOperationController(opModel)

	// init router
	routers := routers.NewRouters(e, opController)

	routers.InitRoutes()
	e.Start(":1323")
}
