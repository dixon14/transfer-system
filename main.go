package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/gin-gonic/gin"

	"transfer-system/config"
	"transfer-system/controller"
	"transfer-system/service"
)

func main() {
	// Initialize db
	db := initDB()
	defer db.Close()

	// Initialize services
	accountService := service.NewAccountService(db)
	transactionService := service.NewTransactionService(db)

	// Initialize controllers
	accountController := controller.NewAccountController(accountService)
	transactionController := controller.NewTransactionController(transactionService)

	initServer(accountController, transactionController)
}

func initDB() *sql.DB {
	// Load database configuration
	dbConfig := config.NewDBConfig()

	// Connect to database
	db, err := config.ConnectDB(dbConfig)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	return db
}

func initServer(accountCtrl *controller.AccountController, transactionCtrl *controller.TransactionController) {
	// Set Gin mode
	if os.Getenv("GIN_MODE") == "" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Create router
	router := gin.Default()

	// Health check endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "healthy",
		})
	})

	// Account endpoints
	router.POST("/accounts", accountCtrl.CreateAccount)
	router.GET("/accounts/:account_id", accountCtrl.GetAccount)

	// Transaction endpoints
	router.POST("/transactions", transactionCtrl.Transfer)

	// Get port from environment or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Start server
	log.Printf("Starting server on port %s...", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
