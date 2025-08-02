package main

import (
	"fmt"
	"main/config"
)

func main() {
	config.ConnectDB()
	fmt.Println("Server started on port 8080")
}
