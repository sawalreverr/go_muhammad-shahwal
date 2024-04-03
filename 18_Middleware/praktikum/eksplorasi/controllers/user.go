package controllers

import (
	"crypto/md5"
	"eksplorasi_middleware/configs"
	"eksplorasi_middleware/middlewares"
	"eksplorasi_middleware/models"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func UserRegisterController(c echo.Context) error {
	var user models.User
	var resErr models.ResponseError

	if err := c.Bind(&user); err != nil {
		resErr.Message = err.Error()
		return c.JSON(http.StatusBadRequest, resErr)
	}

	h := md5.New()
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

	h := md5.New()
	h.Write([]byte(user.Password))
	user.Password = fmt.Sprintf("%x", h.Sum(nil))

	if userExist := configs.DB.First(&user, "email = ? AND password = ?", user.Email, user.Password); userExist.Error != nil {
		return echo.ErrInternalServerError
	}

	userLogin := models.UserLogin{
		ID:    user.ID,
		Email: user.Email,
		Token: middlewares.GenerateTokenJWT(user.Email, user.Role),
	}

	resOk := models.ResponseSuccess{
		Message: "OK",
		Data:    userLogin,
	}

	return c.JSON(http.StatusOK, resOk)
}
