package controllers

import (
	"encoding/json"
	"go-todos-api/database"
	"go-todos-api/models"
	"go-todos-api/services"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestGetAllTodo(t *testing.T) {
	database.InitDatabase()
	database.Migrate()

	service := services.InitTodoService()
	controller := TodoController{service}

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/api/v1/todos", nil)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	err := controller.GetAll(c)

	// pengecekan untuk no error dan status code harus sama dengan 200
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	var response models.BaseResponse[[]models.Todo]
	err = json.Unmarshal(rec.Body.Bytes(), &response)
	assert.NoError(t, err)

	expectedResponse := models.BaseResponse[[]models.Todo]{
		Status:  "success",
		Message: "all todos",
	}

	// pengecekan untuk status dan message saja, tanpa mengecek keseluruhan return data todos
	assert.Equal(t, expectedResponse.Status, response.Status)
	assert.Equal(t, expectedResponse.Message, response.Message)
}

func TestGetTodoByIDSuccess(t *testing.T) {
	database.InitDatabase()
	database.Migrate()

	service := services.InitTodoService()
	controller := TodoController{service}

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/api/v1/todos", nil)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	// mengecek id = 1
	c.SetParamNames("id")
	c.SetParamValues("1")

	err := controller.GetByID(c)

	// pengecekan untuk no error dan status code harus sama dengan 200
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	response := models.BaseResponse[models.Todo]{}
	err = json.Unmarshal(rec.Body.Bytes(), &response)
	assert.NoError(t, err)

	expectedResponse := models.BaseResponse[models.Todo]{
		Status:  "success",
		Message: "todo found",
	}

	// pengecekan untuk status dan message saja, tanpa mengecek isi data dari todo id = 1
	assert.Equal(t, expectedResponse.Status, response.Status)
	assert.Equal(t, expectedResponse.Message, response.Message)
}

func TestGetTodoByIDFailed(t *testing.T) {
	database.InitDatabase()
	database.Migrate()

	service := services.InitTodoService()
	controller := TodoController{service}

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/api/v1/todos", nil)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	// mengecek id = 100000000000000
	c.SetParamNames("id")
	c.SetParamValues("100000000000000")

	err := controller.GetByID(c)

	// pengecekan untuk no error dan status code harus sama dengan 404 / not found
	assert.NoError(t, err)
	assert.Equal(t, http.StatusNotFound, rec.Code)

	response := models.BaseResponse[string]{}
	err = json.Unmarshal(rec.Body.Bytes(), &response)
	assert.NoError(t, err)

	expectedResponse := models.BaseResponse[string]{
		Status:  "failed",
		Message: "todo not found",
	}

	// pengecekan untuk status dan message saja
	assert.Equal(t, expectedResponse.Status, response.Status)
	assert.Equal(t, expectedResponse.Message, response.Message)
}

func TestCreateTodoSuccess(t *testing.T) {
	database.InitDatabase()
	database.Migrate()

	service := services.InitTodoService()
	controller := TodoController{service}

	body := strings.NewReader(`{"title": "new todo","description": "todo for testing"}`)

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/api/v1/todos", body)
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	err := controller.Create(c)

	// pengecekan untuk no error dan status code harus sama dengan 201 / created
	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, rec.Code)

	response := models.BaseResponse[models.Todo]{}
	err = json.Unmarshal(rec.Body.Bytes(), &response)
	assert.NoError(t, err)

	expectedResponse := models.BaseResponse[models.Todo]{
		Status:  "success",
		Message: "todo created",
	}

	// pengecekan untuk status dan message saja, tanpa mengecek isi data dari todo yang baru dibuat
	assert.Equal(t, expectedResponse.Status, response.Status)
	assert.Equal(t, expectedResponse.Message, response.Message)
}

func TestCreateTodoBindFailed(t *testing.T) {
	database.InitDatabase()
	database.Migrate()

	service := services.InitTodoService()
	controller := TodoController{service}

	body := strings.NewReader(`{"title": "failed new todo","description": "todo for testing"}`)

	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/api/v1/todos", body)

	// content-type yang bukan json, akan menyebabkan bind failed
	// req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	err := controller.Create(c)

	// pengecekan untuk no error dan status code harus sama dengan 400 / bad request
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, rec.Code)

	response := models.BaseResponse[string]{}
	err = json.Unmarshal(rec.Body.Bytes(), &response)
	assert.NoError(t, err)

	expectedResponse := models.BaseResponse[string]{
		Status:  "failed",
		Message: "request invalid",
	}

	// pengecekan untuk status dan message saja
	assert.Equal(t, expectedResponse.Status, response.Status)
	assert.Equal(t, expectedResponse.Message, response.Message)
}

func TestUpdateTodoSuccess(t *testing.T) {
	database.InitDatabase()
	database.Migrate()

	service := services.InitTodoService()
	controller := TodoController{service}

	body := strings.NewReader(`{"title": "update todo", "description": "update todo for testing"}`)

	e := echo.New()
	req := httptest.NewRequest(http.MethodPut, "/api/v1/todos", body)
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	// update/edit todo id = 1
	c.SetParamNames("id")
	c.SetParamValues("1")

	err := controller.Update(c)

	// pengecekan untuk no error dan status code harus sama dengan 200 / ok
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	response := models.BaseResponse[models.Todo]{}
	err = json.Unmarshal(rec.Body.Bytes(), &response)
	assert.NoError(t, err)

	expectedResponse := models.BaseResponse[models.Todo]{
		Status:  "success",
		Message: "todo updated",
	}

	// pengecekan untuk status dan message saja, tanpa mengecek isi data dari todo yang baru di update
	assert.Equal(t, expectedResponse.Status, response.Status)
	assert.Equal(t, expectedResponse.Message, response.Message)
}

func TestUpdateTodoFailed(t *testing.T) {
	database.InitDatabase()
	database.Migrate()

	service := services.InitTodoService()
	controller := TodoController{service}

	body := strings.NewReader(`{"title": "update todo", "description": "update todo for testing"}`)

	e := echo.New()
	req := httptest.NewRequest(http.MethodPut, "/api/v1/todos", body)
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	// update/edit todo id = 10000000 >> id not found
	c.SetParamNames("id")
	c.SetParamValues("100000")

	err := controller.Update(c)

	// pengecekan untuk no error dan status code harus sama dengan 500 / internal server error
	assert.NoError(t, err)
	assert.Equal(t, http.StatusInternalServerError, rec.Code)

	response := models.BaseResponse[string]{}
	err = json.Unmarshal(rec.Body.Bytes(), &response)
	assert.NoError(t, err)

	expectedResponse := models.BaseResponse[string]{
		Status:  "failed",
		Message: "failed to update todo",
	}

	// pengecekan untuk status dan message saja
	assert.Equal(t, expectedResponse.Status, response.Status)
	assert.Equal(t, expectedResponse.Message, response.Message)
}

func TestDeleteTodoSuccess(t *testing.T) {
	database.InitDatabase()
	database.Migrate()

	service := services.InitTodoService()
	controller := TodoController{service}

	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/api/v1/todos", nil)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	// delete todo id = 3++
	c.SetParamNames("id")
	c.SetParamValues("4")

	err := controller.Delete(c)

	// pengecekan untuk no error dan status code harus sama dengan 200 / ok
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	response := models.BaseResponse[string]{}
	err = json.Unmarshal(rec.Body.Bytes(), &response)
	assert.NoError(t, err)

	expectedResponse := models.BaseResponse[string]{
		Status:  "success",
		Message: "todo deleted",
	}

	// pengecekan untuk status dan message saja
	assert.Equal(t, expectedResponse.Status, response.Status)
	assert.Equal(t, expectedResponse.Message, response.Message)
}

func TestDeleteTodoFailed(t *testing.T) {
	database.InitDatabase()
	database.Migrate()

	service := services.InitTodoService()
	controller := TodoController{service}

	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/api/v1/todos", nil)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	// delete todo id = 1000000000
	c.SetParamNames("id")
	c.SetParamValues("1000000000")

	err := controller.Delete(c)

	// pengecekan untuk no error dan status code harus sama dengan 500 / internal server error
	assert.NoError(t, err)
	assert.Equal(t, http.StatusInternalServerError, rec.Code)

	response := models.BaseResponse[string]{}
	err = json.Unmarshal(rec.Body.Bytes(), &response)
	assert.NoError(t, err)

	expectedResponse := models.BaseResponse[string]{
		Status:  "failed",
		Message: "failed to delete todo",
	}

	// pengecekan untuk status dan message saja
	assert.Equal(t, expectedResponse.Status, response.Status)
	assert.Equal(t, expectedResponse.Message, response.Message)
}
