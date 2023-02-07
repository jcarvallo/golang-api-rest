package main

import (
	"log"
	"net/http"
	"os"

	"github.com/jcarvallo/golang-api-rest/db"
	"github.com/jcarvallo/golang-api-rest/routes"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	db.DBConnection()
	db.Migration()

	routes.RoutersIndex()

	http.ListenAndServe(":"+os.Getenv("SERVER_PORT"), routes.Routes)
	log.Println("Server http://localhost:" + os.Getenv("SERVER_PORT"))
}
