package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary Get users
// @Description Retrieve a list of users
// @Tags Users
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{} "Successful response"
// @Router /users [get]
func GetUsers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "List of users",
		"data":    []string{"Alice", "Bob", "Charlie"},
	})
}
