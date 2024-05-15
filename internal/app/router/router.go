package router

import (
	"balance-go/internal/app/handlers"
	"balance-go/internal/app/queries" // Import the queries package

	"github.com/labstack/echo"
	"gorm.io/gorm"
)

// Register routes and start the Echo server
func Register(e *echo.Echo, db *gorm.DB) {

	// Create handler instances with injected dependencies
	userHandler := handlers.NewUserHandler(queries.NewUserQueries(db))
	walletHandler := handlers.NewWalletHandler(queries.NewWalletQueries(db))
	incomeHandler := handlers.NewIncomeHandler(queries.NewIncomeQueries(db))
	incomeCategoriesHandler := handlers.NewIncomeCategoiresHandler(queries.NewIncomeCategoiresQueries(db))

	// Group user routes
	userGroup := e.Group("/users")
	userGroup.GET("", userHandler.GetAllUsers)
	userGroup.GET("/:id", userHandler.GetUser)

	// Group wallet routes
	walletGroup := e.Group("/wallets")
	walletGroup.GET("", walletHandler.GetAllWallets)
	walletGroup.GET("/:id", walletHandler.GetWallet)
	walletGroup.POST("", walletHandler.StoreWallet)

	// Group income routes
	incomeGroup := e.Group("/incomes")
	incomeGroup.GET("", incomeHandler.GetAllIncomes)
	incomeGroup.GET("/:id", incomeHandler.GetIncome)
	incomeGroup.POST("", incomeHandler.StoreIncome)

	// Group income routes
	incomeCategoriesGroup := e.Group("/incomeCategories")
	incomeCategoriesGroup.GET("", incomeCategoriesHandler.GetAllIncomeCategoires)
	incomeCategoriesGroup.GET("/:id", incomeCategoriesHandler.GetIncomeCategoires)
	incomeCategoriesGroup.POST("", incomeCategoriesHandler.StoreIncomeCategoires)

	// ... other user routes

	// ... define routes for other models (if any)

	// Start the Echo server
	e.Logger.Fatal(e.Start(":8880"))
}
