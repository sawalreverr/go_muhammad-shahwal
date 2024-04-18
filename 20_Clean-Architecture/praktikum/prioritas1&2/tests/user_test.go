package tests

import (
	"encoding/json"
	"errors"
	"go-wishlist-api/controllers"
	"go-wishlist-api/user"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestRegister(t *testing.T) {
	mockUserService := &MockUserService{}
	mockAuthService := &MockAuthService{}

	mockUserService.On("GetUserByEmail", "test@example.com").Return(user.User{}, nil)
	mockUserService.On("Register", mock.Anything).Return(user.User{ID: 1, Name: "Test User", Email: "test@example.com"}, nil)
	mockAuthService.On("GenerateToken", 1).Return("token", nil)

	controller := controllers.NewUserController(mockUserService, mockAuthService)

	reqBody := strings.NewReader(`{"name": "Test User", "email": "test@example.com", "password": "password"}`)
	e := echo.New()

	req := httptest.NewRequest(http.MethodPost, "/register", reqBody)
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	err := controller.Register(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	var responseBody map[string]interface{}
	err = json.Unmarshal(rec.Body.Bytes(), &responseBody)
	assert.NoError(t, err)

	assert.Equal(t, "register success!", responseBody["status"])
	assert.Equal(t, "test@example.com", responseBody["email"])
}

func TestLogin(t *testing.T) {
	mockUserService := &MockUserService{}
	mockAuthService := &MockAuthService{}

	loginInput := user.UserLogin{
		Email:    "test@example.com",
		Password: "password",
	}

	mockUserService.On("Login", loginInput).Return(user.User{ID: 1, Name: "Test User", Email: "test@example.com"}, nil)
	mockAuthService.On("GenerateToken", 1).Return("token", nil)

	controller := controllers.NewUserController(mockUserService, mockAuthService)

	reqBody := strings.NewReader(`{"email": "test@example.com", "password": "password"}`)
	e := echo.New()

	req := httptest.NewRequest(http.MethodPost, "/login", reqBody)
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	err := controller.Login(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	var responseBody map[string]interface{}
	err = json.Unmarshal(rec.Body.Bytes(), &responseBody)
	assert.NoError(t, err)

	assert.Equal(t, "login success!", responseBody["status"])
	assert.Equal(t, "Test User", responseBody["data"].(map[string]interface{})["name"])
	assert.Equal(t, "test@example.com", responseBody["data"].(map[string]interface{})["email"])
	assert.Equal(t, "token", responseBody["data"].(map[string]interface{})["token"])
}

func TestLoginFailed(t *testing.T) {
	mockUserService := &MockUserService{}
	mockAuthService := &MockAuthService{}

	mockUserService.On("Login", user.UserLogin{Email: "invalid@example.com", Password: "wrongpassword"}).Return(user.User{}, errors.New("invalid credentials"))

	controller := controllers.NewUserController(mockUserService, mockAuthService)

	reqBody := strings.NewReader(`{"email": "invalid@example.com", "password": "wrongpassword"}`)
	e := echo.New()

	req := httptest.NewRequest(http.MethodPost, "/login", reqBody)
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	err := controller.Login(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, rec.Code)

	var responseBody map[string]interface{}
	err = json.Unmarshal(rec.Body.Bytes(), &responseBody)
	assert.NoError(t, err)

	assert.Equal(t, "login failed!", responseBody["error"])
}

func TestLoginInvalid(t *testing.T) {
	mockUserService := &MockUserService{}
	mockAuthService := &MockAuthService{}

	controller := controllers.NewUserController(mockUserService, mockAuthService)

	reqBody := strings.NewReader(`{"password": "password"}`)
	e := echo.New()

	req := httptest.NewRequest(http.MethodPost, "/login", reqBody)
	// req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	err := controller.Login(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, rec.Code)

	var responseBody map[string]interface{}
	err = json.Unmarshal(rec.Body.Bytes(), &responseBody)
	assert.NoError(t, err)

	assert.Equal(t, "request invalid", responseBody["error"])
}
