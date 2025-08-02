package main

import (
	"fmt"
	"log"
	"main/config"
	"main/routes"
	"net/http"
)

func main() {
	config.ConnectDB()

	r := routes.AppRoute()

	fmt.Println("Server started on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
