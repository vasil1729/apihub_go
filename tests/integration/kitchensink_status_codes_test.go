package integration_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/ultimatum/apihub_go/internal/domain/kitchensink"
)

func TestKitchenSinkStatusCodesAPI_Integration(t *testing.T) {
	router := setupTestRouter()

	tests := []struct {
		name     string
		code     int
		expected int
	}{
		{"Status 200", 200, 200},
		{"Status 201", 201, 201},
		{"Status 400", 400, 400},
		{"Status 404", 404, 404},
		{"Status 418", 418, 418},
		{"Status 500", 500, 500},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			url := fmt.Sprintf("/api/v1/kitchen-sink/status/%d", tt.code)
			req, _ := http.NewRequest("GET", url, nil)
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.expected, w.Code)
			
			var resp kitchensink.StatusCodeResponse
			err := json.Unmarshal(w.Body.Bytes(), &resp)
			assert.NoError(t, err)
			assert.Equal(t, tt.code, resp.Code)
			assert.NotEmpty(t, resp.Message)
		})
	}
	
	t.Run("Invalid Status Code", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api/v1/kitchen-sink/status/999", nil)
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
	
	t.Run("Invalid Format", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api/v1/kitchen-sink/status/abc", nil)
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}
