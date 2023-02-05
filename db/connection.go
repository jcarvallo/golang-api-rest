package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var DSN = "host=localhost user=postgres password=123456 dbname=db-api-golang port=5432 sslmode=disable"

func DBConnection() {
	var error error
	DB, error = gorm.Open(postgres.Open(DSN), &gorm.Config{})
	if error != nil {
		log.Fatal(error)
	} else {
		log.Println("DB connected")
	}
}
