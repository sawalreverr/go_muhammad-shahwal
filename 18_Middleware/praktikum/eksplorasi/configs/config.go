package configs

import (
	"eksplorasi_middleware/models"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "root:root@tcp(localhost:3306)/eksplorasi_middleware?charset=utf8mb4&parseTime=True&loc=Local"

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return
	}

	errMigrate := AutoMigration()
	if errMigrate != nil {
		return
	}

	log.Println("Database Connected")
}

func AutoMigration() error {
	err := DB.AutoMigrate(&models.User{}, &models.Category{}, &models.Post{})
	return err
}
