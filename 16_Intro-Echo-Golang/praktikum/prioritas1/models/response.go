package models

type ResponseSuccess struct {
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type ResponseError struct {
	Error string `json:"error"`
}
