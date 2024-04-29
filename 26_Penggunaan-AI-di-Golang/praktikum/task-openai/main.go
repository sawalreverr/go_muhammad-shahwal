package main

import (
	"log"
	"task-openai/controllers"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Cannot load env file. Err: %s", err)
	}
}

func main() {
	e := echo.New()

	e.POST("/rekomen", controllers.RecommendControllers)

	e.Logger.Fatal(e.Start(":4000"))
}
