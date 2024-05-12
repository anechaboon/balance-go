package main

import (
	"balance-go/internal/app/mysql"
	"balance-go/internal/app/router"
	"fmt"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4" // Use v4 for Echo import
)

func main() {

	// Load values from .env file
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	// ... establish database connection
	db, err := mysql.ConnectToMySQL()
	if err != nil {
		fmt.Printf("\nlog:ConnectToMySQL-Error:")

		panic(err) // Handle error appropriately
	}

	// Create an Echo instance with the correct version
	e := echo.New() // Assuming using v4 of Echo

	// Register routes and start server with dependency injection
	router.Register(e, db)

	// Start the Echo server
	e.Logger.Fatal(e.Start(":8080")) // Replace with your desired port
}
