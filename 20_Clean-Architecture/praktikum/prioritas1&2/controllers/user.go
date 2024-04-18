package controllers

import (
	"go-wishlist-api/auth"
	"go-wishlist-api/user"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	userService user.Service
	authService auth.Service
}

func NewUserController(userService user.Service, authService auth.Service) *UserController {
	return &UserController{userService, authService}
}

func (uc *UserController) Register(c echo.Context) error {
	var userInput user.UserRegister

	if err := c.Bind(&userInput); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "request invalid",
		})
	}

	_, err := uc.userService.GetUserByEmail(userInput.Email)
	if err != nil {
		return c.JSON(http.StatusConflict, echo.Map{
			"error": "email already exists",
		})
	}

	_, err = uc.userService.Register(userInput)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": "register failed!",
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"status": "register success!",
		"email":  userInput.Email,
	})
}

func (uc *UserController) Login(c echo.Context) error {
	var userInput user.UserLogin

	if err := c.Bind(&userInput); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "request invalid",
		})
	}

	loginUser, err := uc.userService.Login(userInput)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "login failed!",
		})
	}

	token, err := uc.authService.GenerateToken(int(loginUser.ID))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "login failed!",
		})
	}

	responseData := map[string]interface{}{
		"name":  loginUser.Name,
		"email": loginUser.Email,
		"token": token,
	}

	return c.JSON(http.StatusOK, echo.Map{
		"status": "login success!",
		"data":   responseData,
	})
}
