package database

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitialDatabase() {
	var err error

	var dsn string = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		GetConfig("DB_USERNAME"),
		GetConfig("DB_PASSWORD"),
		GetConfig("DB_HOST"),
		GetConfig("DB_PORT"),
		GetConfig("DB_NAME"),
	)

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("error when creating a connection to the database: %s\n", err)
		panic(err.Error())
	}

	log.Println("connected to the database")
}
