package controllers

import (
	"crypto/sha512"
	"fmt"
	"net/http"
	"prioritas2_middleware/configs"
	"prioritas2_middleware/middlewares"
	"prioritas2_middleware/models"

	"github.com/labstack/echo/v4"
)

func UserRegisterController(c echo.Context) error {
	var user models.User
	var resErr models.ResponseError

	if err := c.Bind(&user); err != nil {
		resErr.Message = err.Error()
		return c.JSON(http.StatusBadRequest, resErr)
	}

	h := sha512.New()
	h.Write([]byte(user.Password))
	user.Password = fmt.Sprintf("%x", h.Sum(nil))

	if saveUser := configs.DB.Create(&user); saveUser.Error != nil {
		resErr.Message = saveUser.Error.Error()
		return c.JSON(http.StatusInternalServerError, resErr)
	}

	resOk := models.ResponseSuccess{
		Message: "OK",
		Data:    user,
	}

	return c.JSON(http.StatusOK, resOk)
}

func UserLoginController(c echo.Context) error {
	var user models.User
	var resErr models.ResponseError

	if err := c.Bind(&user); err != nil {
		resErr.Message = err.Error()
		return c.JSON(http.StatusBadRequest, resErr)
	}

	h := sha512.New()
	h.Write([]byte(user.Password))
	user.Password = fmt.Sprintf("%x", h.Sum(nil))

	if userExist := configs.DB.First(&user, "email = ? AND password = ?", user.Email, user.Password); userExist.Error != nil {
		resErr.Message = userExist.Error.Error()
		return c.JSON(http.StatusInternalServerError, resErr)
	}

	userLogin := models.UserLogin{
		ID:    user.ID,
		Email: user.Email,
		Token: middlewares.GenerateTokenJWT(user.ID, user.Email),
	}

	resOk := models.ResponseSuccess{
		Message: "OK",
		Data:    userLogin,
	}

	return c.JSON(http.StatusOK, resOk)
}
