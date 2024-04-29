package controllers

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"task-openai/models"

	"github.com/labstack/echo/v4"
	"github.com/sashabaranov/go-openai"
)

func RecommendControllers(c echo.Context) error {
	request := models.Request{}
	response := models.Response{}

	if err := c.Bind(&request); err != nil {
		response.Status = "Bad Request"
		return c.JSON(http.StatusBadRequest, response)
	}

	recommendation, err := getLaptopRecommendation(request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"status":  "API Error",
			"message": err,
		})
	}

	response.Status = "Success"
	response.Data = recommendation

	return c.JSON(http.StatusOK, response)
}

func getLaptopRecommendation(r models.Request) (string, error) {
	client := openai.NewClient(os.Getenv("OPENAI_KEY"))
	messages := make([]openai.ChatCompletionMessage, 0)

	request := fmt.Sprintf(`Berdasarkan budget Rp %v juta dan untuk keperluan %v, rekomendasikan laptop terbaik dengan spesifikasi:
	- Merk: %v (opsional)
	- RAM: %v (opsional)
	- GPU: %v (opsional)
	`, r.Budget, r.Purpose, r.Merk, r.RAM, r.GPU)

	messages = append(messages, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: request,
	})

	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model:    openai.GPT3Dot5Turbo,
			Messages: messages,
		},
	)

	if err != nil {
		return "", err
	}

	return resp.Choices[0].Message.Content, nil
}
