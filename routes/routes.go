package routes

import (
	controllers "main/controllers"

	"github.com/gorilla/mux"
)

func AppRoute() *mux.Router {

	api := "/api/v1/"

	route := mux.NewRouter()

	route.HandleFunc(api+"auth/register", controllers.RegisterUser).Methods("POST")
	route.HandleFunc(api+"auth/login", controllers.LoginUser).Methods("POST")

	return route
}
