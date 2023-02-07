package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jcarvallo/golang-api-rest/db"
	"github.com/jcarvallo/golang-api-rest/models"
	"github.com/jcarvallo/golang-api-rest/utils"
)

func GetTasks(w http.ResponseWriter, r *http.Request) {
	var tasks []models.Task
	db.DB.Find(&tasks)
	json.NewEncoder(w).Encode(&tasks)
}
func GetTask(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	params := mux.Vars(r)
	db.DB.First(&task, params["id"])
	if task.ID == 0 {
		utils.RespondNotFound("Task not found", w)

	}
	db.DB.Model(&task).Association("Task").Find(&task.UserID)
	json.NewEncoder(w).Encode(&task)
}
func PostTask(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	json.NewDecoder(r.Body).Decode(&task)

	createTask := db.DB.Create(&task)
	error := createTask.Error
	if error != nil {
		utils.RespondWithError(error, w)

	}
	json.NewEncoder(w).Encode(&task)
}
func DeleteTask(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	params := mux.Vars(r)
	db.DB.First(&task, params["id"])
	if task.ID == 0 {
		utils.RespondNotFound("Task not found", w)

	}
	db.DB.Unscoped().Delete(&task)
	utils.RespondNotContent("task delete", w)

}
