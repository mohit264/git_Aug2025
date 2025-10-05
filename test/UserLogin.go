// Implement Backend User Login
package test
import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"your_project/handlers"
	"your_project/models"
	"your_project/utils"
	"github.com/stretchr/testify/assert"
)

func TestUserLogin(t *testing.T) {
	// Initialize the router and database
	router := handlers.SetupRouter()
	utils.InitDB()
	defer utils.CloseDB()

	// Create a test user
	testUser := models.User{
		Username: "testuser",
		Password: "testpassword",
	}
	utils.CreateUser(&testUser)

	// Create a login request
	loginRequest := models.LoginRequest{
		Username: "testuser",
		Password: "testpassword",
	}
	jsonValue, _ := json.Marshal(loginRequest)

	// Send a POST request to the login endpoint
	req, err := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonValue))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Record the response
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	// Check the response code
	assert.Equal(t, http.StatusOK, recorder.Code)

	// Check the response body
	var loginResponse models.LoginResponse
	err = json.Unmarshal(recorder.Body.Bytes(), &loginResponse)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotEmpty(t, loginResponse.Token)
}