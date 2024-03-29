package controllers

import (
	"eksplorasi_echo_golang/models"
	"encoding/json"
	"net/http"
	"slices"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

var Fruits = []models.Fruit{}

func GetAllFruitsController(c echo.Context) error {
	resOk := models.ResponseSuccess{
		Message: "all fruits",
		Data:    Fruits,
	}

	return c.JSON(http.StatusOK, resOk)
}

func GetFruitByIdController(c echo.Context) error {
	id := c.Param("id")

	idx := slices.IndexFunc(Fruits, func(fruit models.Fruit) bool {
		return fruit.Id.String() == id
	})

	if idx == -1 {
		resErr := models.ResponseError{
			Message: "data not found",
		}

		return c.JSON(http.StatusBadRequest, resErr)
	}

	resOk := models.ResponseSuccess{
		Message: "ok",
		Data:    &Fruits[idx],
	}

	return c.JSON(http.StatusOK, resOk)
}

func CreateFruitController(c echo.Context) error {
	fruit := models.Fruit{Id: uuid.New()}

	c.Bind(&fruit)

	nutritions := getNutritions(fruit.Name)
	fruit.Nutritions = nutritions

	// if err := c.Bind(&fruit); err != nil {
	// 	resErr := models.ResponseError{
	// 		Message: "bad request",
	// 	}

	// 	return c.JSON(http.StatusBadRequest, resErr)
	// }

	Fruits = append(Fruits, fruit)

	resOk := models.ResponseSuccess{
		Message: "fruit added",
		Data:    fruit,
	}

	return c.JSON(http.StatusOK, resOk)
}

func DeleteFruitController(c echo.Context) error {
	id := c.Param("id")
	idx := slices.IndexFunc(Fruits, func(fruit models.Fruit) bool {
		return fruit.Id.String() == id
	})

	if idx == -1 {
		resErr := models.ResponseError{
			Message: "data not found",
		}

		return c.JSON(http.StatusBadRequest, resErr)
	}

	Fruits = slices.DeleteFunc(Fruits, func(fruit models.Fruit) bool {
		return fruit.Id.String() == id
	})

	resOk := models.ResponseSuccess{
		Message: "fruit deleted",
		Data:    id,
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
