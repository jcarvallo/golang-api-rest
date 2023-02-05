package main

import (
	"net/http"

	"github.com/jcarvallo/golang-api-rest/db"
	"github.com/jcarvallo/golang-api-rest/routes"
)

func main() {
	db.DBConnection()
	db.Migration()

	routes.RoutersIndex()

	http.ListenAndServe(":3000", routes.Routes)
}
