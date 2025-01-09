package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary Get status
// @Description Get the server status
// @Tags Status
// @Accept json
// @Produce json
// @Success 200 {object} map[string]string
// @Router /status [get]
func GetStatus(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}
