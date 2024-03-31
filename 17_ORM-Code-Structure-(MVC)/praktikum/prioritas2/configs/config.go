package configs

import (
	"log"
	"prioritas2_orm_mvc/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "root:root@tcp(localhost:3306)/prioritas2_orm_mvc?charset=utf8mb4&parseTime=True&loc=Local"

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	if err := AutoMigrate(); err != nil {
		log.Fatal(err)
	}
}

func AutoMigrate() error {
	err := DB.AutoMigrate(&models.Fruit{}, &models.Nutritions{})

	return err
}
