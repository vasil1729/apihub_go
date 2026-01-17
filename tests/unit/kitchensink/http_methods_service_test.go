package kitchensink_test

import (
	"bytes"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/ultimatum/apihub_go/internal/service/kitchensink"
)

func TestHTTPMethodsService_ProcessRequest(t *testing.T) {
	service := kitchensink.NewHTTPMethodsService()

	t.Run("GET Request", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/test?q=hello", nil)
		req.Header.Set("X-Custom-Header", "value")

		resp, err := service.ProcessRequest(req)
		assert.NoError(t, err)
		assert.Equal(t, "GET", resp.Method)
		assert.Equal(t, "/test?q=hello", resp.URL)
		assert.Equal(t, "value", resp.Headers["X-Custom-Header"][0])
		assert.Equal(t, "hello", resp.Query["q"][0])
	})

	t.Run("POST Request with Body", func(t *testing.T) {
		body := []byte(`{"name":"test"}`)
		req := httptest.NewRequest("POST", "/test", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")

		resp, err := service.ProcessRequest(req)
		assert.NoError(t, err)
		assert.Equal(t, "POST", resp.Method)
		
		// Body parsing verification
		bodyMap, ok := resp.Body.(map[string]interface{})
		assert.True(t, ok)
		assert.Equal(t, "test", bodyMap["name"])
	})
}
