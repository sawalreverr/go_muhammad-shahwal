package models

import "github.com/google/uuid"

type Fruit struct {
	Id         uuid.UUID  `json:"id"`
	Name       string     `json:"name" form:"name"`
	Price      int        `json:"price" form:"price"`
	Nutritions Nutritions `json:"nutritions" form:"nutritions"`
}

type Nutritions struct {
	Calories      float64 `json:"calories"`
	Fat           float64 `json:"fat"`
	Sugar         float64 `json:"sugar"`
	Carbohydrates float64 `json:"carbohydrates"`
	Protein       float64 `json:"protein"`
}
