package web

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestNewError(t *testing.T) {

	r := gin.Default()

	r.GET("/test-error", func(c *gin.Context) {
		NewError(c, 400, "Bad Request")
	})

	req, _ := http.NewRequest(http.MethodGet, "/test-error", nil)

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Fatalf("Expected status code %d, got %d", http.StatusBadRequest, w.Code)
	}

	var body map[string]interface{}
	if err := json.Unmarshal(w.Body.Bytes(), &body); err != nil {
		t.Fatalf("Failed to unmarshal response body: %v", err)
	}

	if body["status"].(float64) != 400 {
		t.Errorf("Expected status 400, got %v", body["status"])
	}

	if body["error"].(string) != "Bad Request" {
		t.Errorf("Expected error 'Bad Request', got %v", body["error"])
	}
}

func TestErrorMethod(t *testing.T) {
	err := "Test error"
	errorStruct := &Error{
		Status: 500,
		Err:    err,
	}

	expectedError := "Status: 500, Error: Test error"
	if errorStruct.Error() != expectedError {
		t.Errorf("Expected error: %s, but got: %s", expectedError, errorStruct.Error())
	}
}
