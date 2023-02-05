package db

import "github.com/jcarvallo/golang-api-rest/models"

func Migration() {
	DB.AutoMigrate(models.Task{})
	DB.AutoMigrate(models.User{})
}
