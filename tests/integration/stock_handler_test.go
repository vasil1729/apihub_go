package integration_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/ultimatum/apihub_go/pkg/response"
)

func TestStocksAPI_Integration(t *testing.T) {
	router := setupTestRouter()

	t.Run("GET /api/v1/public/stocks - success", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api/v1/public/stocks?page=1&limit=10", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var resp response.PaginatedResponse
		err := json.Unmarshal(w.Body.Bytes(), &resp)
		assert.NoError(t, err)
		assert.Equal(t, 200, resp.StatusCode)
		assert.Equal(t, "Stocks fetched successfully", resp.Message)
		assert.NotNil(t, resp.Data)
		assert.NotNil(t, resp.Pagination)
	})

	t.Run("GET /api/v1/public/stocks/:symbol - success (TCS)", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api/v1/public/stocks/TCS", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var resp response.Response
		err := json.Unmarshal(w.Body.Bytes(), &resp)
		assert.NoError(t, err)
		assert.Equal(t, 200, resp.StatusCode)
		assert.Equal(t, "Stock fetched successfully", resp.Message)
	})

	t.Run("GET /api/v1/public/stocks/:symbol - case insensitive", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api/v1/public/stocks/reliance", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("GET /api/v1/public/stocks/:symbol - not found", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api/v1/public/stocks/NOTEXIST", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusNotFound, w.Code)
	})
}
