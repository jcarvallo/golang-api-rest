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
	dsnString := []string{"host=" + os.Getenv("HOST_DB"), "user=" + os.Getenv("USER_DB"), "password=" + os.Getenv("PASSWORD_DB"), "dbname=" + os.Getenv("NAME_DB"), "port=" + os.Getenv("PORT_DB"), "sslmode=" + os.Getenv("SSLMODE")}
	DSN = strings.Join(dsnString, " ")
	DB, error = gorm.Open(postgres.Open(DSN), &gorm.Config{})
	if error != nil {
		log.Fatal(error)
	} else {
		log.Println("DB connected")
	}
}
