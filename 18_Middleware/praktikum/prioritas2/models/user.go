package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        int            `json:"id" gorm:"primaryKey"`
	Email     string         `json:"email"`
	Password  string         `json:"password"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

type UserLogin struct {
	ID    int    `json:"id" gorm:"primaryKey"`
	Email string `json:"email"`
	Token string `json:"token"`
}
