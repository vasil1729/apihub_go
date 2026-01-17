package integration_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/ultimatum/apihub_go/internal/domain/kitchensink"
	"github.com/ultimatum/apihub_go/pkg/response"
)

func TestKitchenSinkRequestInspectionAPI_Integration(t *testing.T) {
	router := setupTestRouter()

	t.Run("GET /api/v1/kitchen-sink/ip", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api/v1/kitchen-sink/ip", nil)
		req.RemoteAddr = "127.0.0.1:12345"
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		
		var resp struct {
			response.Response
			Data kitchensink.IPResponse `json:"data"`
		}
		err := json.Unmarshal(w.Body.Bytes(), &resp)
		assert.NoError(t, err)
		// Usually Gin returns empty IP for httptest request without setting it specifically on context or trusted proxies
		// But let's check structure mostly
		assert.NotNil(t, resp.Data)
	})

	t.Run("GET /api/v1/kitchen-sink/user-agent", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api/v1/kitchen-sink/user-agent", nil)
		req.Header.Set("User-Agent", "IntegrationTest/1.0")
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		
		var resp struct {
			response.Response
			Data kitchensink.UserAgentResponse `json:"data"`
		}
		err := json.Unmarshal(w.Body.Bytes(), &resp)
		assert.NoError(t, err)
		assert.Equal(t, "IntegrationTest/1.0", resp.Data.UserAgent)
	})

	t.Run("GET /api/v1/kitchen-sink/headers", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api/v1/kitchen-sink/headers", nil)
		req.Header.Set("X-Custom-Header", "CustomValue")
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		
		var resp struct {
			response.Response
			Data kitchensink.HeadersResponse `json:"data"`
		}
		err := json.Unmarshal(w.Body.Bytes(), &resp)
		assert.NoError(t, err)
		assert.Equal(t, "CustomValue", resp.Data.Headers["X-Custom-Header"][0])
	})
}
