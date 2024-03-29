package main

import (
	"prioritas1_echo_golang/routes"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}
	routes.InitRoute(e)
	e.Logger.Fatal(e.Start(":8080"))
}

// Validating Data Bind, source from https://dasarpemrogramangolang.novalagung.com/C-http-request-payload-validation.html
type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}
