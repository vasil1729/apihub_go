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

func TestKitchenSinkCookiesAPI_Integration(t *testing.T) {
	router := setupTestRouter()

	t.Run("Set Cookie", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api/v1/kitchen-sink/cookies/set?key=mycookie&value=myvalue", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		
		// Check Set-Cookie header
		cookies := w.Result().Cookies()
		assert.NotEmpty(t, cookies)
		found := false
		for _, c := range cookies {
			if c.Name == "mycookie" && c.Value == "myvalue" {
				found = true
				break
			}
		}
		assert.True(t, found, "Cookie 'mycookie' should be set")
	})

	t.Run("Get Cookies", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api/v1/kitchen-sink/cookies/get", nil)
		req.AddCookie(&http.Cookie{Name: "existing", Value: "cookie"})
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var resp struct {
			response.Response
			Data kitchensink.CookieResponse `json:"data"`
		}
		err := json.Unmarshal(w.Body.Bytes(), &resp)
		assert.NoError(t, err)
		assert.Equal(t, "cookie", resp.Data.Cookies["existing"])
	})

	t.Run("Delete Cookie", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api/v1/kitchen-sink/cookies/delete?key=todelete", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		// Check Set-Cookie header for deletion (max-age < 0 or expire in past)
		cookies := w.Result().Cookies()
		found := false
		for _, c := range cookies {
			if c.Name == "todelete" && c.MaxAge < 0 {
				found = true
				break
			}
		}
		assert.True(t, found, "Cookie 'todelete' should be expired/deleted")
	})
	
	t.Run("Set Cookie - Missing Params", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/api/v1/kitchen-sink/cookies/set?key=foo", nil) // missing value
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}
