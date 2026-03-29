package tests

import (
	"fmt"
	"gin-app/handlers"
	"gin-app/router"
	"gin-app/storage"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func performRequest(r http.Handler, method, path string, body string, token string) *httptest.ResponseRecorder {
	var bodyReader io.Reader
	if body != "" {
		bodyReader = strings.NewReader(body)
	}

	req := httptest.NewRequest(method, path, bodyReader)
	if body != "" {
		req.Header.Set("Content-Type", "application/json")
	}
	if token != "" {
		req.Header.Set("Authorization", token)
	}
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func TestCreateUser(t *testing.T) {
	tests := []struct {
		name           string
		body           string
		token          string
		expectedStatus int
		expectedMsg    string
	}{
		{
			name:           "Correct creating",
			body:           `{"name":"Piotr","email":"test@email.com"}`,
			token:          "test",
			expectedStatus: http.StatusCreated,
			expectedMsg:    `{"id":1,"name":"Piotr","email":"test@email.com"}`,
		},
		{
			name:           "Unauthorized",
			body:           `{"name":"Piotr"}`,
			token:          "",
			expectedStatus: http.StatusUnauthorized,
			expectedMsg:    "missing auth token",
		},
		{
			name:           "Bad mail format",
			body:           `{"name":"Piotr","email":"not-an-email"}`,
			token:          "test",
			expectedStatus: http.StatusBadRequest,
			expectedMsg:    "invalid request body",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			store := storage.NewMemoryStorage()
			handler := handlers.NewUserHandler(store)
			r := router.SetupRouter(handler)

			w := performRequest(r, "POST", "/api/private/users", tt.body, tt.token)

			assert.Equal(t, tt.expectedStatus, w.Code)

			if tt.expectedStatus == http.StatusCreated {
				assert.Contains(t, w.Body.String(), tt.expectedMsg)
				users := store.GetUsers()
				assert.Len(t, users, 1)
			}
		})
	}
}

func TestUpdateUser(t *testing.T) {
	tests := []struct {
		name           string
		body           string
		token          string
		expectedStatus int
		expectedMsg    string
		userID         string
	}{
		{
			name:           "Correct updating",
			body:           `{"name":"Piotr","email":"test@email.com"}`,
			token:          "test",
			expectedStatus: http.StatusOK,
			expectedMsg:    `{"id":1,"name":"Piotr","email":"test@email.com"}`,
			userID:         "1",
		},
		{
			name:           "Unauthorized",
			body:           `{"name":"Piotr"}`,
			token:          "",
			expectedStatus: http.StatusUnauthorized,
			expectedMsg:    "missing auth token",
			userID:         "1",
		},
		{
			name:           "Bad mail format",
			body:           `{"name":"Piotr","email":"not-an-email"}`,
			token:          "test",
			expectedStatus: http.StatusBadRequest,
			expectedMsg:    "invalid request body",
			userID:         "1",
		},
		{
			name:           "User not found",
			body:           `{"name":"Piotr","email":"test@email.com"}`,
			token:          "test",
			expectedStatus: http.StatusNotFound,
			expectedMsg:    "user not found",
			userID:         "2",
		},
		{
			name:           "Invalid id",
			body:           `{"name":"Piotr","email":"test@email.com"}`,
			token:          "test",
			expectedStatus: http.StatusBadRequest,
			expectedMsg:    "invalid id",
			userID:         "asd",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			store := storage.NewMemoryStorage()
			store.AddUser("Piotr", "piotr@email.com")
			handler := handlers.NewUserHandler(store)
			r := router.SetupRouter(handler)
			path := fmt.Sprintf("/api/private/users/%s", tt.userID)
			w := performRequest(r, "PUT", path, tt.body, tt.token)

			assert.Equal(t, tt.expectedStatus, w.Code)
		})
	}
}

func TestDeleteUser(t *testing.T) {
	tests := []struct {
		name           string
		body           string
		token          string
		expectedStatus int
		expectedMsg    string
		userID         string
	}{
		{
			name:           "Correct deleting",
			body:           "",
			token:          "test",
			expectedStatus: http.StatusOK,
			expectedMsg:    `{"id":1,"name":"Piotr","email":"test@email.com"}`,
			userID:         "1",
		},
		{
			name:           "Unauthorized",
			body:           "",
			token:          "",
			expectedStatus: http.StatusUnauthorized,
			expectedMsg:    "missing auth token",
			userID:         "1",
		},
		{
			name:           "User not found",
			body:           "",
			token:          "test",
			expectedStatus: http.StatusNotFound,
			expectedMsg:    "user not found",
			userID:         "2",
		},
		{
			name:           "Invalid id",
			body:           "",
			token:          "test",
			expectedStatus: http.StatusBadRequest,
			expectedMsg:    "invalid id",
			userID:         "asd",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			store := storage.NewMemoryStorage()
			store.AddUser("Piotr", "piotr@email.com")
			handler := handlers.NewUserHandler(store)
			r := router.SetupRouter(handler)
			path := fmt.Sprintf("/api/private/users/%s", tt.userID)
			w := performRequest(r, "DELETE", path, tt.body, tt.token)

			assert.Equal(t, tt.expectedStatus, w.Code)
		})
	}
}

func TestGetUser(t *testing.T) {
	tests := []struct {
		name           string
		body           string
		token          string
		expectedStatus int
		expectedMsg    string
		userID         string
	}{
		{
			name:           "Correct getting",
			body:           "",
			token:          "test",
			expectedStatus: http.StatusOK,
			expectedMsg:    `{"id":1,"name":"Piotr","email":"piotr@email.com"}`,
			userID:         "1",
		},
		{
			name:           "Unauthorized",
			body:           "",
			token:          "",
			expectedStatus: http.StatusUnauthorized,
			expectedMsg:    "missing auth token",
			userID:         "1",
		},
		{
			name:           "User not found",
			body:           "",
			token:          "test",
			expectedStatus: http.StatusNotFound,
			expectedMsg:    "user not found",
			userID:         "2",
		},
		{
			name:           "Invalid id",
			body:           "",
			token:          "test",
			expectedStatus: http.StatusBadRequest,
			expectedMsg:    "invalid id",
			userID:         "asd",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			store := storage.NewMemoryStorage()
			store.AddUser("Piotr", "piotr@email.com")
			handler := handlers.NewUserHandler(store)
			r := router.SetupRouter(handler)
			path := fmt.Sprintf("/api/private/users/%s", tt.userID)
			w := performRequest(r, "GET", path, tt.body, tt.token)

			assert.Equal(t, tt.expectedStatus, w.Code)
		})
	}
}
