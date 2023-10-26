package main_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestLoginPage(t *testing.T) {
	r := gin.Default()
	// Setup your Gin routes and middleware here

	// Create a test request to the login page
	req, _ := http.NewRequest("GET", "/login", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Assert the HTTP status code and response content
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Login Page")
}
