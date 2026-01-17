package integration_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKitchenSinkImagesAPI_Integration(t *testing.T) {
	router := setupTestRouter()

	t.Run("GET JPEG", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api/v1/kitchen-sink/images/jpeg", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, "image/jpeg", w.Header().Get("Content-Type"))
		assert.NotEmpty(t, w.Body.Bytes())
	})

	t.Run("GET PNG", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api/v1/kitchen-sink/images/png", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, "image/png", w.Header().Get("Content-Type"))
		assert.NotEmpty(t, w.Body.Bytes())
	})

	t.Run("GET SVG", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api/v1/kitchen-sink/images/svg", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Contains(t, w.Header().Get("Content-Type"), "image/svg+xml") // Gin might append charset?
		assert.NotEmpty(t, w.Body.String())
	})

	t.Run("GET WebP", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api/v1/kitchen-sink/images/webp", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, "image/webp", w.Header().Get("Content-Type"))
		assert.NotEmpty(t, w.Body.Bytes())
	})
}
