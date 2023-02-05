package routes

import "github.com/jcarvallo/golang-api-rest/controllers"

func RoutesTasks() {
	Routes.HandleFunc("/tasks", controllers.GetTasks).Methods("GET")
	Routes.HandleFunc("/tasks/{id}", controllers.GetTask).Methods("GET")
	Routes.HandleFunc("/tasks", controllers.PostTask).Methods("POST")
	Routes.HandleFunc("/tasks/{id}", controllers.DeleteTask).Methods("DELETE")

}
