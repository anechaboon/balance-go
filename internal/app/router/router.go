package router

import (
	"balance-go/internal/app/handlers"
	"balance-go/internal/app/queries" // Import the queries package

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// Register routes and start the Echo server
func Register(e *echo.Echo, db *gorm.DB) {

	// Create handler instances with injected dependencies
	userHandler := handlers.NewUserHandler(queries.NewUserQueries(db))

	// Group user routes
	userGroup := e.Group("/users")
	userGroup.GET("", userHandler.GetAllUsers)
	// ... other user routes

	// ... define routes for other models (if any)

	// Start the Echo server
	e.Logger.Fatal(e.Start(":8080"))
}
