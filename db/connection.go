package db

import (
	"log"
	"os"
	"strings"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

var DSN string

func DBConnection() {
	var error error
	var dsnString []string

	dsnString = append(dsnString, "host="+os.Getenv("HOST_DB"))
	dsnString = append(dsnString, "user="+os.Getenv("USER_DB"))
	dsnString = append(dsnString, "password="+os.Getenv("PASSWORD_DB"))
	dsnString = append(dsnString, "dbname="+os.Getenv("NAME_DB"))
	dsnString = append(dsnString, "port="+os.Getenv("PORT_DB"))
	dsnString = append(dsnString, "sslmode="+os.Getenv("SSLMODE"))
	DSN = strings.Join(dsnString, " ")

	DB, error = gorm.Open(postgres.Open(DSN), &gorm.Config{})
	if error != nil {
		log.Fatal(error)
	} else {
		log.Println("DB connected")
	}
}
