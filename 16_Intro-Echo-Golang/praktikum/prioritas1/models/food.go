package models

import "github.com/google/uuid"

type Food struct {
	Id          uuid.UUID `json:"id"`
	Name        string    `json:"name" form:"name" validate:"required"`
	Price       int       `json:"price" form:"price" validate:"required"`
	Description string    `json:"description" form:"description"`
}
