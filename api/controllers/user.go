package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetUsers example
func GetUsers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "List of users",
		"data":    []string{"Alice", "Bob", "Charlie"},
	})
}
