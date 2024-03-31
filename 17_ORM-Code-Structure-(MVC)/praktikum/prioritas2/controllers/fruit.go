package controllers

import (
	"encoding/json"
	"net/http"
	"prioritas2_orm_mvc/configs"
	"prioritas2_orm_mvc/models"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func GetAllFruitsController(c echo.Context) error {
	fruits := []models.Fruit{}
	resErr := models.ResponseError{}

	if findAll := configs.DB.Preload("Nutritions").Find(&fruits); findAll.Error != nil {
		resErr.Message = "Failed get all data"
		return c.JSON(http.StatusInternalServerError, resErr)
	}

	resOk := models.ResponseSuccess{
		Message: "all fruits",
		Data:    fruits,
	}

	return c.JSON(http.StatusOK, resOk)
}

func GetFruitByIdController(c echo.Context) error {
	id := c.Param("id")
	fruit := models.Fruit{}
	resErr := models.ResponseError{}

	if findById := configs.DB.Preload("Nutritions").First(&fruit, uuid.MustParse(id)); findById.Error != nil {
		resErr.Message = "Failed get data"
		return c.JSON(http.StatusInternalServerError, resErr)
	}

	resOk := models.ResponseSuccess{
		Message: "ok",
		Data:    fruit,
	}

	return c.JSON(http.StatusOK, resOk)
}

func CreateFruitController(c echo.Context) error {
	fruit := models.Fruit{Id: uuid.New()}
	c.Bind(&fruit)

	nutritions := getNutritions(fruit.Name)
	fruit.Nutritions = nutritions

	if createData := configs.DB.Create(&fruit); createData.Error != nil {
		resErr := models.ResponseError{
			Message: "Failed create data",
		}
		return c.JSON(http.StatusInternalServerError, resErr)
	}

	resOk := models.ResponseSuccess{
		Message: "Fruit added",
		Data:    fruit,
	}

	return c.JSON(http.StatusOK, resOk)
}

func DeleteFruitController(c echo.Context) error {
	id := c.Param("id")
	fruit := models.Fruit{}
	resErr := models.ResponseError{}

	if deleteById := configs.DB.Delete(&fruit, uuid.MustParse(id)); deleteById.Error != nil {
		resErr.Message = deleteById.Error.Error()
		return c.JSON(http.StatusInternalServerError, resErr)
	}

	resOk := models.ResponseSuccess{
		Message: "Fruit deleted",
		Data:    "",
	}

	return c.JSON(http.StatusOK, resOk)
}

func getNutritions(fruitName string) models.Nutritions {
	baseURL := "https://www.fruityvice.com/api/fruit/"

	nutrition := models.Nutritions{}
	client := &http.Client{}

	req, err := http.NewRequest("GET", baseURL+fruitName, nil)
	if err != nil {
		return nutrition
	}

	res, err := client.Do(req)
	if err != nil {
		return nutrition
	}

	if res.StatusCode == 404 {
		return nutrition
	}

	defer res.Body.Close()

	var tempData map[string]interface{}
	json.NewDecoder(res.Body).Decode(&tempData)

	nutritionMap := tempData["nutritions"].(map[string]interface{})
	nutrition.Calories = nutritionMap["calories"].(float64)
	nutrition.Fat = nutritionMap["fat"].(float64)
	nutrition.Sugar = nutritionMap["sugar"].(float64)
	nutrition.Carbohydrates = nutritionMap["carbohydrates"].(float64)
	nutrition.Protein = nutritionMap["protein"].(float64)

	return nutrition
}
