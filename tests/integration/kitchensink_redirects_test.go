package integration_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKitchenSinkRedirectsAPI_Integration(t *testing.T) {
	router := setupTestRouter()

	tests := []struct {
		name       string
		path       string
		expectCode int
		toURL      string
	}{
		{"301 Moved Permanently", "/api/v1/kitchen-sink/redirects/301", 301, "https://google.com"},
		{"302 Found", "/api/v1/kitchen-sink/redirects/302", 302, "https://google.com"},
		{"307 Temporary Redirect", "/api/v1/kitchen-sink/redirects/307", 307, "https://google.com"},
		{"308 Permanent Redirect", "/api/v1/kitchen-sink/redirects/308", 308, "https://google.com"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", tt.path, nil)
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.expectCode, w.Code)
			assert.Equal(t, tt.toURL, w.Header().Get("Location"))
		})
	}

	t.Run("Custom URL Redirect", func(t *testing.T) {
		w := httptest.NewRecorder()
		customURL := "https://example.com/foo"
		req, _ := http.NewRequest("GET", "/api/v1/kitchen-sink/redirects/302?url="+customURL, nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, 302, w.Code)
		assert.Equal(t, customURL, w.Header().Get("Location"))
	})
}
