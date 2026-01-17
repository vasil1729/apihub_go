package integration_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/ultimatum/apihub_go/internal/domain/kitchensink"
	"github.com/ultimatum/apihub_go/pkg/response"
)

func TestKitchenSinkHTTPMethodsAPI_Integration(t *testing.T) {
	router := setupTestRouter()

	t.Run("GET /api/v1/kitchen-sink/http-methods/get", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api/v1/kitchen-sink/http-methods/get?foo=bar", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		
		var resp struct {
			response.Response
			Data kitchensink.HTTPMethodResponse `json:"data"`
		}
		err := json.Unmarshal(w.Body.Bytes(), &resp)
		assert.NoError(t, err)
		assert.Equal(t, "GET", resp.Data.Method)
		assert.Equal(t, "bar", resp.Data.Query["foo"][0])
	})

	t.Run("POST /api/v1/kitchen-sink/http-methods/post", func(t *testing.T) {
		w := httptest.NewRecorder()
		body := []byte(`{"key":"value"}`)
		req, _ := http.NewRequest("POST", "/api/v1/kitchen-sink/http-methods/post", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		
		var resp struct {
			response.Response
			Data kitchensink.HTTPMethodResponse `json:"data"`
		}
		err := json.Unmarshal(w.Body.Bytes(), &resp)
		assert.NoError(t, err)
		assert.Equal(t, "POST", resp.Data.Method)
		
		bodyMap, ok := resp.Data.Body.(map[string]interface{})
		assert.True(t, ok)
		assert.Equal(t, "value", bodyMap["key"])
	})

	t.Run("PUT /api/v1/kitchen-sink/http-methods/put", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("PUT", "/api/v1/kitchen-sink/http-methods/put", nil)
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("DELETE /api/v1/kitchen-sink/http-methods/delete", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("DELETE", "/api/v1/kitchen-sink/http-methods/delete", nil)
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
	})
}
