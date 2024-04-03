package models

type Category struct {
	ID    int    `json:"id" gorm:"primaryKey"`
	Name  string `json:"name"`
	Posts []Post `json:"posts" gorm:"foreignKey:CategoryID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
