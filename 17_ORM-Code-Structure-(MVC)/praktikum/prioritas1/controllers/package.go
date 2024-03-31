package controllers

import (
	"net/http"
	"prioritas1_orm_mvc/configs"
	"prioritas1_orm_mvc/models"
	"strconv"

	"github.com/labstack/echo/v4"
)

// GET /api/v1/packages
func GetAllPackagesController(c echo.Context) error {
	packages := []models.Package{}

	findAll := configs.DB.Find(&packages)

	if findAll.Error != nil {
		resErr := models.ResponseError{
			Message: "Failed Get Data Packages",
		}

		return c.JSON(http.StatusInternalServerError, resErr)
	}

	resOk := models.ResponseSuccess{
		Message: "OK",
		Data:    packages,
	}

	return c.JSON(http.StatusOK, resOk)
}

// GET /api/v1/packages/:id
func GetPackageByIdController(c echo.Context) error {
	packageModel := models.Package{}
	resErr := models.ResponseError{}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		resErr.Message = "Bad Request"

		return c.JSON(http.StatusBadRequest, resErr)
	}

	if findId := configs.DB.First(&packageModel, id); findId.Error != nil {
		resErr.Message = "Data Not Found"

		return c.JSON(http.StatusInternalServerError, resErr)
	}

	resOk := models.ResponseSuccess{
		Message: "OK",
		Data:    packageModel,
	}

	return c.JSON(http.StatusOK, resOk)
}

// POST /api/v1/packages
func CreatePackageController(c echo.Context) error {
	packageModel := models.Package{}
	resErr := models.ResponseError{}

	if err := c.Bind(&packageModel); err != nil {
		resErr.Message = "Bad Request"

		return c.JSON(http.StatusBadRequest, resErr)
	}

	if createData := configs.DB.Create(&packageModel); createData.Error != nil {
		resErr.Message = "Failed insert to database"

		return c.JSON(http.StatusInternalServerError, resErr)
	}

	resOk := models.ResponseSuccess{
		Message: "Data Inserted",
		Data:    packageModel,
	}

	return c.JSON(http.StatusOK, resOk)
}

// PUT /api/v1/packages/:id
func EditPackageByIdController(c echo.Context) error {
	packageModel := models.Package{}
	resErr := models.ResponseError{}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		resErr.Message = "Bad Request"
		return c.JSON(http.StatusBadRequest, resErr)
	}

	if findId := configs.DB.First(&packageModel, id); findId.Error != nil {
		resErr.Message = "Data Not Found"
		return c.JSON(http.StatusInternalServerError, resErr)
	}

	if err := c.Bind(&packageModel); err != nil {
		resErr.Message = "Bad Request"
		return c.JSON(http.StatusBadRequest, resErr)
	}

	if updateData := configs.DB.Save(&packageModel); updateData.Error != nil {
		resErr.Message = "Failed to update data"
		return c.JSON(http.StatusInternalServerError, resErr)
	}

	resOk := models.ResponseSuccess{
		Message: "Data Updated",
		Data:    packageModel,
	}

	return c.JSON(http.StatusOK, resOk)
}

// DELETE /api/v1/packages/:id
func DeletePackageByIdController(c echo.Context) error {
	packageModel := models.Package{}
	resErr := models.ResponseError{}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		resErr.Message = "Bad Request"
		return c.JSON(http.StatusBadRequest, resErr)
	}

	if findId := configs.DB.First(&packageModel, id); findId.Error != nil {
		resErr.Message = "Data Not Found"
		return c.JSON(http.StatusInternalServerError, resErr)
	}

	if deleteData := configs.DB.Delete(&packageModel, id); deleteData.Error != nil {
		resErr.Message = deleteData.Error.Error()
		return c.JSON(http.StatusInternalServerError, resErr)
	}

	resOk := models.ResponseSuccess{
		Message: "Data Deleted",
		Data:    "",
	}

	return c.JSON(http.StatusOK, resOk)
}
