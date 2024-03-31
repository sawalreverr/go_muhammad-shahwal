package configs

import (
	"fmt"
	"log"
	"prioritas1_orm_mvc/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "root:root@tcp(localhost:3306)/prioritas1_orm_mvc?charset=utf8mb4&parseTime=True&loc=Local"

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	errMigrate := AutoMigration()
	if errMigrate != nil {
		log.Fatal(errMigrate)
	}

	fmt.Println("Database Connected")
}

func AutoMigration() error {
	err := DB.AutoMigrate(&models.Package{})
	return err
}
