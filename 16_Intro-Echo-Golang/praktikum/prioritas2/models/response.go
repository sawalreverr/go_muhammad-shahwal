package models

type ResponseSuccess struct {
	Data    any    `json:"data"`
	Message string `json:"message"`
}

type ResponseError struct {
	Message string `json:"message"`
}
