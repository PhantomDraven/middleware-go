package controllers

import (
	"context"
	"net/http"
	"time"

	"middleware-go/api/database"
	"middleware-go/api/models"

	gin "github.com/gin-gonic/gin"
	logrus "github.com/sirupsen/logrus"
)

// @Summary Add a new namespace
// @Description Add a new namespace to the database
// @Tags Namespace
// @Accept json
// @Produce json
// @Param namespace body models.Namespace true "Namespace"
// @Success 200 {object} models.Namespace
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /namespaces [post]
func AddNamespace(c *gin.Context) {
	var namespace models.Namespace
	if err := c.ShouldBindJSON(&namespace); err != nil {
		logrus.WithError(err).Error("Failed to bind JSON")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
		return
	}

	if namespace.Name == "" {
		logrus.Error("Name attribute is required")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Name attribute is required"})
		return
	}

	id, err := generateID()
	if err != nil {
		logrus.WithError(err).Error("Failed to generate unique ID")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate unique ID"})
		return
	}
	namespace.ID = id
	namespace.CreatedAt = time.Now()
	namespace.UpdatedAt = time.Now()

	err = database.FirestoreDB.NewRef("namespaces/"+namespace.ID).Set(context.Background(), namespace)
	if err != nil {
		logrus.WithError(err).Error("Failed to add namespace to database")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add namespace to database"})
		return
	}

	logrus.WithField("namespace_id", namespace.ID).Info("Namespace added successfully")
	c.JSON(http.StatusOK, namespace)
}

func generateID() (string, error) {
	ref, err := database.FirestoreDB.NewRef("namespaces").Push(context.Background(), nil)
	if err != nil {
		return "", err
	}
	return ref.Key, nil
}

func checkIfNamespaceExists(id string) bool {
	var namespace models.Namespace
	database.FirestoreDB.NewRef("namespaces/"+id).Get(context.Background(), &namespace)
	return namespace.ID != ""
}

// @Summary Remove a namespace
// @Description Remove a namespace from the database
// @Tags Namespace
// @Accept json
// @Produce json
// @Param id path string true "Namespace ID"
// @Success 200 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /namespaces/{id} [delete]
func RemoveNamespace(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		logrus.Error("Namespace ID is required")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Namespace ID is required"})
		return
	}

	exists := checkIfNamespaceExists(id)
	if !exists {
		logrus.Error("Namespace does not exist")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Namespace does not exist"})
		return
	}

	err := database.FirestoreDB.NewRef("namespaces/" + id).Delete(context.Background())
	if err != nil {
		logrus.WithError(err).Error("Failed to remove namespace from database")
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	logrus.WithField("namespace_id", id).Info("Namespace removed successfully")
	c.JSON(http.StatusOK, gin.H{"message": "Namespace deleted successfully"})
}

// @Summary Edit a namespace
// @Description Edit an existing namespace in the database
// @Tags Namespace
// @Accept json
// @Produce json
// @Param id path string true "Namespace ID"
// @Param namespace body models.Namespace true "Namespace"
// @Success 200 {object} models.Namespace
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /namespaces/{id} [put]
func EditNamespace(c *gin.Context) {
	id := c.Param("id")
	exists := checkIfNamespaceExists(id)
	if !exists {
		logrus.Error("Namespace does not exist")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Namespace does not exist"})
		return
	}

	var namespace models.Namespace
	if err := c.ShouldBindJSON(&namespace); err != nil {
		logrus.WithError(err).Error("Failed to bind JSON")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
		return
	}

	namespace.UpdatedAt = time.Now()

	err := database.FirestoreDB.NewRef("namespaces/"+id).Update(context.Background(), map[string]interface{}{
		"name":       namespace.Name,
		"updated_at": namespace.UpdatedAt,
	})
	if err != nil {
		logrus.WithError(err).Error("Failed to update namespace in database")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update namespace in database"})
		return
	}

	logrus.WithField("namespace_id", id).Info("Namespace updated successfully")
	c.JSON(http.StatusOK, namespace)
}
