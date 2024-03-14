package controllers

import (
	"net/http"
	"praktikum/eksplorasi/go-math-api/models"

	"github.com/labstack/echo/v4"
)

type OperationController struct {
	OperationModel *models.OperationModel
}

func NewOperationController(opMod *models.OperationModel) *OperationController {
	return &OperationController{OperationModel: opMod}
}

func (oc *OperationController) Add(c echo.Context) error {
	oq := new(models.Operation)
	if err := c.Bind(oq); err != nil {
		return err
	}

	result := oq.A + oq.B

	return c.JSON(http.StatusOK, echo.Map{
		"result": result,
	})
}

func (oc *OperationController) Subtract(c echo.Context) error {
	oq := new(models.Operation)
	if err := c.Bind(oq); err != nil {
		return err
	}

	result := oq.A - oq.B

	return c.JSON(http.StatusOK, echo.Map{
		"result": result,
	})
}

func (oc *OperationController) Multiply(c echo.Context) error {
	oq := new(models.Operation)
	if err := c.Bind(oq); err != nil {
		return err
	}

	result := oq.A * oq.B

	return c.JSON(http.StatusOK, echo.Map{
		"result": result,
	})
}

func (oc *OperationController) Divide(c echo.Context) error {
	oq := new(models.Operation)
	if err := c.Bind(oq); err != nil {
		return err
	}

	if oq.B == 0 {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "division by zero",
		})
	}

	result := oq.A / oq.B

	return c.JSON(http.StatusOK, echo.Map{
		"result": result,
	})
}
