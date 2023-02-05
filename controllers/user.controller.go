package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jcarvallo/golang-api-rest/db"
	"github.com/jcarvallo/golang-api-rest/models"
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
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User not found"))
		return
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
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(error.Error()))
		return
	}
	json.NewEncoder(w).Encode(&user)
}
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	params := mux.Vars(r)
	db.DB.First(&user, params["id"])
	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User not found"))
		return
	}
	db.DB.Unscoped().Delete(&user)
	w.WriteHeader(http.StatusNoContent)
	w.Write([]byte("User delete"))
}
