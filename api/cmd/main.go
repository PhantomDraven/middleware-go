package main

import (
	"os"
	"fmt"

	"net/http"
	"github.com/gin-gonic/gin"
	"middleware-go/api/controllers" // Import controllers

	_ "middleware-go/api/docs"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

// @title Middleware-Go
// @version 0.0.1-alpha
// @description Swagger of middleware-go from https://github.com/PhantomDraven/middleware-go 
// @license.name MIT
// @license.url https://github.com/PhantomDraven/middleware-go/blob/main/LICENSE
// @host localhost:3000
// @BasePath /
func main() {
	port := os.Getenv("API_PORT")
	fmt.Printf("Port setted as %s \n", port)
	if port == "" {
		port = "3000" // Default port
		fmt.Printf("Using default port \n")
	}

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

	// Start the server on port 8080
	fmt.Printf("Starting server on port %s...\n", port)
	r.Run(":" + port)
}
