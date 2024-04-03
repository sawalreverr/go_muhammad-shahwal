package configs

import (
	"log"
	"prioritas1_middleware/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "root:root@tcp(localhost:3306)/prioritas1_middleware?charset=utf8mb4&parseTime=True&loc=Local"

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		// log.Fatal(err)
		return
	}

	errMigrate := AutoMigration()
	if errMigrate != nil {
		// log.Fatal(errMigrate)
		return
	}

	log.Println("Database Connected")
}

func AutoMigration() error {
	err := DB.AutoMigrate(&models.Package{}, &models.User{})
	return err
}
