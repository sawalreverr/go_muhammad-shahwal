package controllers

import (
	"eksplorasi_middleware/configs"
	"eksplorasi_middleware/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetAllCategoryController(c echo.Context) error {
	var categories []models.Category
	var resErr models.ResponseError

	if findAll := configs.DB.Preload("Posts").Find(&categories); findAll.Error != nil {
		resErr.Message = findAll.Error.Error()
		return c.JSON(http.StatusInternalServerError, resErr)
	}

	resOk := models.ResponseSuccess{
		Message: "OK",
		Data:    categories,
	}

	return c.JSON(http.StatusOK, resOk)
}

func CreateNewCategoryController(c echo.Context) error {
	var category models.Category
	var resErr models.ResponseError

	if bindCategory := c.Bind(&category); bindCategory != nil {
		resErr.Message = "Bad Request"
		return c.JSON(http.StatusBadRequest, resErr)
	}

	if createCategory := configs.DB.Create(&category); createCategory.Error != nil {
		resErr.Message = createCategory.Error.Error()
		return c.JSON(http.StatusInternalServerError, resErr)
	}

	resOk := models.ResponseSuccess{
		Message: "OK",
		Data:    category,
	}

	return c.JSON(http.StatusOK, resOk)
}

func DeleteCategoryByIdController(c echo.Context) error {
	var category models.Category
	var resErr models.ResponseError

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		resErr.Message = "Bad Request"
		return c.JSON(http.StatusBadRequest, resErr)
	}

	if categoryById := configs.DB.First(&category, id); categoryById.Error != nil {
		resErr.Message = categoryById.Error.Error()
		return c.JSON(http.StatusInternalServerError, resErr)
	}

	if deleteCategory := configs.DB.Delete(&category, id); deleteCategory != nil {
		resErr.Message = deleteCategory.Error.Error()
		return c.JSON(http.StatusInternalServerError, resErr)
	}

	resOk := models.ResponseSuccess{
		Message: "OK",
		Data:    id,
	}

	return c.JSON(http.StatusOK, resOk)
}
