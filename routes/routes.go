package routes

import (
	controllers "main/controllers"

	"github.com/gorilla/mux"
)

func AppRoute() *mux.Router {

	api := "/api/v1/"

	route := mux.NewRouter()

	// Users Routes
	route.HandleFunc(api+"auth/register", controllers.RegisterUser).Methods("POST")
	route.HandleFunc(api+"auth/login", controllers.LoginUser).Methods("POST")

	// Products Routes
	route.HandleFunc(api+"add-product", controllers.CreateProduct).Methods("POST")
	route.HandleFunc(api+"get-all-products", controllers.GetAllProducts).Methods("GET")

	return route
}
