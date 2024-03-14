package routers

import (
	"praktikum/eksplorasi/go-math-api/controllers"

	"github.com/labstack/echo/v4"
)

type OperationRouters struct {
	E                   *echo.Echo
	OperationController *controllers.OperationController
}

func NewRouters(e *echo.Echo, opController *controllers.OperationController) *OperationRouters {
	return &OperationRouters{
		E:                   e,
		OperationController: opController,
	}
}

func (r *OperationRouters) InitRoutes() {
	e := r.E
	e.POST("/api/add", r.OperationController.Add)
	e.POST("/api/subtract", r.OperationController.Subtract)
	e.POST("/api/multiply", r.OperationController.Multiply)
	e.POST("/api/divide", r.OperationController.Divide)
}
