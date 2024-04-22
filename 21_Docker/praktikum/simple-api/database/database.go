package database

import (
	"database/sql"
	"fmt"
	"simple-api-docker/utils"

	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		utils.GetConfig("DB_USERNAME"),
		utils.GetConfig("DB_PASSWORD"),
		utils.GetConfig("DB_HOST"),
		utils.GetConfig("DB_PORT"),
		utils.GetConfig("DB_NAME"),
	)

	var err error

	DB, err = sql.Open("mysql", dsn)

	if err != nil {
		log.Fatalf("error when connecting to the database: %v", err)
	}

	log.Println("connected to the database")

	Migrate()
}

func Migrate() {
	_, err := DB.Exec(`CREATE TABLE IF NOT EXISTS todos(
		id INT PRIMARY KEY AUTO_INCREMENT,
		task VARCHAR(255) NOT NULL,
		is_completed bool NOT NULL
	);`)

	if err != nil {
		log.Fatalf("error when perform database migration: %v", err)
	}

	log.Println("database migration success")
}
