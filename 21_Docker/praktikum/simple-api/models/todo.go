package models

type Todo struct {
	ID          uint   `json:"id"`
	Task        string `json:"task"`
	IsCompleted bool   `json:"is_completed"`
}
