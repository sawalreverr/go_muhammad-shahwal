package models

import (
	"time"

	"gorm.io/gorm"
)

type Post struct {
	ID        int            `json:"id" gorm:"primaryKey"`
	Title     string         `json:"title"`
	Content   string         `json:"content"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}
