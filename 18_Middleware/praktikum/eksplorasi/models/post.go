package models

import (
	"time"

	"gorm.io/gorm"
)

type Post struct {
	ID         int            `json:"id" gorm:"primaryKey"`
	CategoryID int            `json:"category_id"`
	Title      string         `json:"title"`
	Content    string         `json:"content"`
	CreatedAt  time.Time      `json:"-"`
	UpdatedAt  time.Time      `json:"-"`
	DeletedAt  gorm.DeletedAt `json:"-" gorm:"index"`
}
