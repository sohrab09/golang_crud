package routes

import (
	controllers "main/controller"

	"github.com/gorilla/mux"
)

func AppRoute() *mux.Router {

	apiEndPoints := "/api/v1/"

	route := mux.NewRouter()

	route.HandleFunc(apiEndPoints+"auth/register", controllers.RegisterUser).Methods("POST")
	return route
}
