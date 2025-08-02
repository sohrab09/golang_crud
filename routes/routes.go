package routes

import (
	"github.com/gorilla/mux"
)

func RegisterRoute() *mux.Router {
	route := mux.NewRouter()
	route.HandleFunc("/register", controllers.RegisterUser).Methods("POST")
	return route
}
