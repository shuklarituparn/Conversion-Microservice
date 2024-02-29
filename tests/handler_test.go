package tests

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/shuklarituparn/Conversion-Microservice/internal/handlers"
	_ "github.com/smartystreets/goconvey/convey"
	"github.com/stretchr/testify/assert"
)

// TestAccountDeleteSuccess tests the successful deletion of an account.
func TestAccountDeleteSuccess(t *testing.T) {
	// Setup Gin router
	r := gin.Default()
	r.GET("/profile/delete", handlers.AccountDelete)

	// Prepare a request
	req, err := http.NewRequest("GET", "/account/delete?userID=1&code=secure_code", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a response recorder to record the response
	w := httptest.NewRecorder()

	// Serve the request
	r.ServeHTTP(w, req)

	// Assertions
	assert.Equal(t, http.StatusOK, w.Code)
	// Add more assertions as needed
}

// TestAccountDeleteUnauthorized tests handling unauthorized deletion attempts.
func TestAccountDeleteUnauthorized(t *testing.T) {
	// Setup Gin router
	r := gin.Default()
	r.GET("/profile/delete", handlers.AccountDelete)

	// Prepare a request
	req, err := http.NewRequest("GET", "/account/delete?userID=2&code=secure_code", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a response recorder to record the response
	w := httptest.NewRecorder()

	// Serve the request
	r.ServeHTTP(w, req)

	// Assertions
	assert.Equal(t, http.StatusForbidden, w.Code)
	// Add more assertions as needed
}

// Annotate tests with Allure annotations
func TestMain(m *testing.M) {
	// Run tests with Allure annotation setup
	os.Exit(m.Run())
}
