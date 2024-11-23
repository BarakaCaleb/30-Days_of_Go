package main

import (
	"log"
	"online-song-library/config"
	"online-song-library/handlers"
	"os"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Online Song Library API
// @version 1.0
// @description This is an API for managing an online song library
// @termsOfService http://example.com/terms/

// @contact.name API Support
// @contact.url http://example.com/contact
// @contact.email support@example.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /
func main() {
	// Load environment variables
	err := config.LoadEnv()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Initialize database
	db, err := config.ConnectDB()
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	// Initialize router
	router := gin.Default()

	// Swagger route
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Register routes
	handlers.RegisterRoutes(router, db)

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Server running on port %s", port)
	router.Run(":" + port)
}
