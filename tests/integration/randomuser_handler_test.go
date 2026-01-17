package integration_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	publicHandler "github.com/ultimatum/apihub_go/internal/handler/public"
	"github.com/ultimatum/apihub_go/pkg/response"
)

func setupTestRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	
	// Setup public routes
	v1 := router.Group("/api/v1")
	publicRoutes := v1.Group("/public")
	err := publicHandler.SetupPublicRoutes(publicRoutes, "../../data")
	if err != nil {
		panic(err)
	}
	
	return router
}

func TestRandomUsersAPI_Integration(t *testing.T) {
	router := setupTestRouter()

	t.Run("GET /api/v1/public/randomusers - success", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api/v1/public/randomusers?page=1&limit=10", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var resp response.PaginatedResponse
		err := json.Unmarshal(w.Body.Bytes(), &resp)
		assert.NoError(t, err)
		assert.Equal(t, 200, resp.StatusCode)
		assert.Equal(t, "Random users fetched successfully", resp.Message)
		assert.NotNil(t, resp.Data)
		assert.NotNil(t, resp.Pagination)
		assert.Equal(t, 1, resp.Pagination.Page)
		assert.Equal(t, 10, resp.Pagination.Limit)
	})

	t.Run("GET /api/v1/public/randomusers/:id - success", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api/v1/public/randomusers/1", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var resp response.Response
		err := json.Unmarshal(w.Body.Bytes(), &resp)
		assert.NoError(t, err)
		assert.Equal(t, 200, resp.StatusCode)
		assert.Equal(t, "User fetched successfully", resp.Message)
		assert.NotNil(t, resp.Data)
	})

	t.Run("GET /api/v1/public/randomusers/:id - not found", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api/v1/public/randomusers/99999", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusNotFound, w.Code)

		var resp response.Response
		err := json.Unmarshal(w.Body.Bytes(), &resp)
		assert.NoError(t, err)
		assert.Equal(t, 404, resp.StatusCode)
		assert.Equal(t, "User not found", resp.Message)
	})

	t.Run("GET /api/v1/public/randomusers/:id - invalid ID", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api/v1/public/randomusers/abc", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)

		var resp response.Response
		err := json.Unmarshal(w.Body.Bytes(), &resp)
		assert.NoError(t, err)
		assert.Equal(t, 400, resp.StatusCode)
		assert.Equal(t, "Invalid user ID", resp.Message)
	})

	t.Run("GET /api/v1/public/randomusers/random - success", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api/v1/public/randomusers/random", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var resp response.Response
		err := json.Unmarshal(w.Body.Bytes(), &resp)
		assert.NoError(t, err)
		assert.Equal(t, 200, resp.StatusCode)
		assert.Equal(t, "Random user fetched successfully", resp.Message)
		assert.NotNil(t, resp.Data)
	})
}
