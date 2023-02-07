package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jcarvallo/golang-api-rest/db"
	"github.com/jcarvallo/golang-api-rest/models"
	"github.com/jcarvallo/golang-api-rest/utils"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	db.DB.Find(&users)
	json.NewEncoder(w).Encode(&users)
}
func GetUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	params := mux.Vars(r)
	db.DB.First(&user, params["id"])
	if user.ID == 0 {
		utils.RespondNotFound("User not found", w)
	}
	db.DB.Model(&user).Association("Tasks").Find(&user.Tasks)
	json.NewEncoder(w).Encode(&user)

}
func PostUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	createUser := db.DB.Create(&user)
	error := createUser.Error
	if error != nil {
		utils.RespondWithError(error, w)
	}
	json.NewEncoder(w).Encode(&user)
}
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	params := mux.Vars(r)
	db.DB.First(&user, params["id"])
	if user.ID == 0 {
		utils.RespondNotFound("User not found", w)
	}
	db.DB.Unscoped().Delete(&user)
	utils.RespondNotContent("User delete", w)
}
