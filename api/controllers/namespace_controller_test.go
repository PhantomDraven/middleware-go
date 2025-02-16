package controllers

import (
	"bytes"
	"encoding/json"
	"middleware-go/api/database"
	"middleware-go/api/models"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func setupRouter() *gin.Engine {
	// Load environment variables
	err := godotenv.Load("../.env")
	if err != nil {
		panic("Error loading .env file")
	}

	// Initialize Firebase
	database.InitializeFirebase()

	router := gin.Default()
	router.POST("/namespaces", AddNamespace)
	router.DELETE("/namespaces/:id", RemoveNamespace)
	router.PUT("/namespaces/:id", EditNamespace)
	return router
}

func setupBaseNamespace(router *gin.Engine) string {
	// create the namespace (to be deleted)
	namespace := models.Namespace{
		Name: "Test Namespace",
	}
	jsonValue, _ := json.Marshal(namespace)
	req, _ := http.NewRequest("POST", "/namespaces", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	var responseNamespace models.Namespace
	json.Unmarshal(w.Body.Bytes(), &responseNamespace)

	return responseNamespace.ID
}

func TestAddNamespace(t *testing.T) {
	router := setupRouter()

	// Test successful addition
	namespace := models.Namespace{
		Name: "Test Namespace",
	}

	jsonValue, _ := json.Marshal(namespace)
	req, _ := http.NewRequest("POST", "/namespaces", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var responseNamespace models.Namespace
	err := json.Unmarshal(w.Body.Bytes(), &responseNamespace)
	assert.NoError(t, err)
	assert.Equal(t, namespace.Name, responseNamespace.Name)
	assert.NotEmpty(t, responseNamespace.ID)
	assert.WithinDuration(t, time.Now(), responseNamespace.CreatedAt, time.Second)
	assert.WithinDuration(t, time.Now(), responseNamespace.UpdatedAt, time.Second)

	// Test error on missing name attribute
	invalidNamespace := models.Namespace{}

	jsonValue, _ = json.Marshal(invalidNamespace)
	req, _ = http.NewRequest("POST", "/namespaces", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)

	var errorResponse map[string]string
	err = json.Unmarshal(w.Body.Bytes(), &errorResponse)
	assert.NoError(t, err)
	assert.Contains(t, errorResponse["error"], "Name attribute is required")

	// remove tests
	req, _ = http.NewRequest("DELETE", "/namespaces/"+responseNamespace.ID, nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
}

func TestRemoveNamespace(t *testing.T) {
	router := setupRouter()

	idNamespace := setupBaseNamespace(router)

	// Test successful removal
	req, _ := http.NewRequest("DELETE", "/namespaces/"+idNamespace, nil)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "Namespace deleted successfully", response["message"])

	// Test error on invalid ID
	req, _ = http.NewRequest("DELETE", "/namespaces/invalid-id", nil)

	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)

	err = json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Contains(t, response["error"], "Namespace does not exist")
}

func TestEditNamespace(t *testing.T) {
	router := setupRouter()

	idNamespace := setupBaseNamespace(router)

	// Test successful edit
	namespace := models.Namespace{
		Name: "Updated Namespace",
	}

	jsonValue, _ := json.Marshal(namespace)
	req, _ := http.NewRequest("PUT", "/namespaces/"+idNamespace, bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var responseNamespace models.Namespace
	err := json.Unmarshal(w.Body.Bytes(), &responseNamespace)
	assert.NoError(t, err)
	assert.Equal(t, namespace.Name, responseNamespace.Name)
	assert.WithinDuration(t, time.Now(), responseNamespace.UpdatedAt, time.Second)

	// Test error on invalid ID
	var response map[string]string
	req, _ = http.NewRequest("PUT", "/namespaces/invalid-id", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)

	err = json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Contains(t, response["error"], "Namespace does not exist")

	// remove tests
	req, _ = http.NewRequest("DELETE", "/namespaces/"+idNamespace, nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)
}
