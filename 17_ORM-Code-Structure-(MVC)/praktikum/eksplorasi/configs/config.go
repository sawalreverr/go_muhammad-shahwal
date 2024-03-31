package configs

import (
	"eksplorasi_orm_mvc/models"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "root:root@tcp(localhost:3306)/eksplorasi_orm_mvc?charset=utf8mb4&parseTime=True&loc=Local"

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	if err := AutoMigrate(); err != nil {
		log.Fatal(err)
	}

	log.Println("DATABASE CONNECTED")
}

func AutoMigrate() error {
	err := DB.AutoMigrate(&models.Post{}, &models.Comment{})

	return err
}
