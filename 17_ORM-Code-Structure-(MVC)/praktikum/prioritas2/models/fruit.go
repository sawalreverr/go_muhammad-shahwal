package models

import (
	"github.com/google/uuid"
)

type Fruit struct {
	Id         uuid.UUID  `json:"id" gorm:"primaryKey"`
	Name       string     `json:"name" form:"name"`
	Price      int        `json:"price" form:"price"`
	Nutritions Nutritions `json:"nutritions" form:"nutritions" gorm:"foreignKey:FruitID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

type Nutritions struct {
	FruitID       uuid.UUID `json:"-" gorm:"size:191"`
	Calories      float64   `json:"calories"`
	Fat           float64   `json:"fat"`
	Sugar         float64   `json:"sugar"`
	Carbohydrates float64   `json:"carbohydrates"`
	Protein       float64   `json:"protein"`
}
