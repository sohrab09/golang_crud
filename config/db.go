package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
)

var DB *sql.DB

func ConnectDB() {
	server := "(local)"
	port := 1433
	database := "ECommerce"
	connectionString := fmt.Sprintf("server=%s;port=%d;database=%s;integrated security=true;",
		server, port, database)

	var err error
	DB, err = sql.Open("sqlserver", connectionString)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("Error connecting to database: ", err.Error())
	}

	fmt.Println("Connected to SQL Server successfully.")
}
