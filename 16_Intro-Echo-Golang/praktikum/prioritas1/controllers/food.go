package controllers

import (
	"net/http"
	"prioritas1_echo_golang/models"
	"slices"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

var Foods = []models.Food{}

func GetAllFoodsController(c echo.Context) error {
	resOk := models.ResponseSuccess{
		Message: "OK",
		Data:    &Foods,
	}
	return c.JSON(http.StatusOK, resOk)
}

func GetFoodByIdController(c echo.Context) error {
	id := c.Param("id")

	idx := slices.IndexFunc(Foods, func(food models.Food) bool {
		return food.Id.String() == id
	})

	if idx == -1 {
		resErr := models.ResponseError{
			Error: "Data Not Found",
		}

		return c.JSON(http.StatusBadRequest, resErr)
	}

	resOk := models.ResponseSuccess{
		Message: "OK",
		Data:    &Foods[idx],
	}

	return c.JSON(http.StatusFound, resOk)
}

func CreateFoodController(c echo.Context) error {
	food := models.Food{Id: uuid.New()}

	resOk := models.ResponseSuccess{
		Message: "Data inserted",
		Data:    &food,
	}

	resErr := models.ResponseError{
		Error: "Bad Request",
	}

	if err := c.Bind(&food); err != nil {
		return c.JSON(http.StatusBadRequest, resErr)
	}

	if err := c.Validate(food); err != nil {
		return c.JSON(http.StatusBadRequest, resErr)
	}

	Foods = append(Foods, food)

	return c.JSON(http.StatusCreated, resOk)
}

func EditFoodController(c echo.Context) error {
	id := c.Param("id")

	idx := slices.IndexFunc(Foods, func(food models.Food) bool {
		return food.Id.String() == id
	})

	if idx == -1 {
		resErr := models.ResponseError{
			Error: "Data Not Found",
		}

		return c.JSON(http.StatusBadRequest, resErr)
	}

	food := models.Food{Id: uuid.MustParse(id)}

	resErr := models.ResponseError{
		Error: "Bad Request",
	}

	if err := c.Bind(&food); err != nil {
		return c.JSON(http.StatusBadRequest, resErr)
	}

	if err := c.Validate(food); err != nil {
		return c.JSON(http.StatusBadRequest, resErr)
	}

	Foods[idx] = food

	resOk := models.ResponseSuccess{
		Message: "Data Edited",
		Data:    &food,
	}

	return c.JSON(http.StatusOK, resOk)
}

func DeleteFoodController(c echo.Context) error {
	id := c.Param("id")

	idx := slices.IndexFunc(Foods, func(food models.Food) bool {
		return food.Id.String() == id
	})

	if idx == -1 {
		resErr := models.ResponseError{
			Error: "Data Not Found",
		}

		return c.JSON(http.StatusBadRequest, resErr)
	}

	Foods = slices.DeleteFunc(Foods, func(food models.Food) bool {
		return food.Id.String() == id
	})

	resOk := models.ResponseSuccess{
		Message: "Data Deleted",
		Data:    id,
	}

	return c.JSON(http.StatusOK, resOk)
}
