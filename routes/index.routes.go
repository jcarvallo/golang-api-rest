package routes

import (
	"github.com/gorilla/mux"
)

var Routes = mux.NewRouter()

func RoutersIndex() {
	RoutesUser()
	RoutesTasks()
}
