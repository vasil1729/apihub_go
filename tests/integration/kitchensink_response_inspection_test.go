package integration_test

import (
	"encoding/json"
	"encoding/xml"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/ultimatum/apihub_go/internal/domain/kitchensink"
)

func TestKitchenSinkResponseInspectionAPI_Integration(t *testing.T) {
	router := setupTestRouter()

	t.Run("GET /api/v1/kitchen-sink/response/json", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api/v1/kitchen-sink/response/json", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"))

		var resp kitchensink.ResponseInspectionResponse
		err := json.Unmarshal(w.Body.Bytes(), &resp)
		assert.NoError(t, err)
		assert.Equal(t, "json", resp.Format)
	})

	t.Run("GET /api/v1/kitchen-sink/response/xml", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api/v1/kitchen-sink/response/xml", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, "application/xml; charset=utf-8", w.Header().Get("Content-Type"))

		var resp kitchensink.ResponseInspectionResponse
		err := xml.Unmarshal(w.Body.Bytes(), &resp)
		assert.NoError(t, err)
		assert.Equal(t, "xml", resp.Format)
	})

	t.Run("GET /api/v1/kitchen-sink/response/html", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api/v1/kitchen-sink/response/html", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, "text/html; charset=utf-8", w.Header().Get("Content-Type"))
		
		body := w.Body.String()
		assert.True(t, strings.Contains(body, "<!DOCTYPE html>"))
	})
}
