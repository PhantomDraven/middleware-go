package main

import (
	"fmt"
	"log"
	"os"

	"middleware-go/api/controllers" // Import controllers
	"middleware-go/api/database"
	"net/http"

	"github.com/joho/godotenv"

	"github.com/gin-gonic/gin"

	_ "middleware-go/api/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Middleware-Go
// @version 0.0.1-alpha
// @description Swagger of middleware-go from https://github.com/PhantomDraven/middleware-go
// @license.name MIT
// @license.url https://github.com/PhantomDraven/middleware-go/blob/main/LICENSE
// @host localhost:3000
// @BasePath /
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error while loading .env: %v", err)
	}

	port := os.Getenv("API_PORT")
	if port == "" {
		port = "3000" // Default port
	}

	fmt.Printf("Port setted as %s \n", port)

	// Initialize Firebase
	database.InitializeFirebase()

	// Create a new router instance
	r := gin.Default()

	// Swagger endpoint
	r.Use(func(c *gin.Context) {
		if c.Request.URL.Path == "/swagger/" {
			c.Redirect(http.StatusMovedPermanently, "/swagger/index.html")
			return
		}
	})
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Register routes for controllers
	r.GET("/status", controllers.GetStatus)
	r.GET("/users", controllers.GetUsers)

	// Namespace routes
	r.POST("/namespaces", controllers.AddNamespace)
	r.DELETE("/namespaces/:id", controllers.RemoveNamespace)
	r.PUT("/namespaces/:id", controllers.EditNamespace)

	// Start the server on port 8080
	fmt.Printf("Starting server on port %s...\n", port)
	r.Run(":" + port)
}
