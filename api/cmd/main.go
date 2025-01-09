package main

import (
	"os"
	"fmt"

	"github.com/gin-gonic/gin"
	"middleware-go/api/controllers" // Import controllers
)

func main() {
	port := os.Getenv("API_PORT")
	fmt.Printf("Port setted as %s \n", port)
	if port == "" {
		port = "8080" // Default port
		fmt.Printf("Using default port \n")
	}

	// Create a new router instance
	r := gin.Default()

	// Register routes for controllers
	r.GET("/users", controllers.GetUsers)

	// Start the server on port 8080
	fmt.Printf("Starting server on port %s...\n", port)
	r.Run(":" + port)
}
