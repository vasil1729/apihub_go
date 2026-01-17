package kitchensink_test

import (
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/ultimatum/apihub_go/internal/service/kitchensink"
)

func TestRequestInspectionService(t *testing.T) {
	service := kitchensink.NewRequestInspectionService()

	t.Run("GetClientIP", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/", nil)
		req.RemoteAddr = "1.2.3.4:1234"
		
		ip := service.GetClientIP(req)
		assert.Contains(t, ip, "1.2.3.4")
	})

	t.Run("GetUserAgent", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/", nil)
		req.Header.Set("User-Agent", "TestAgent/1.0")
		
		ua := service.GetUserAgent(req)
		assert.Equal(t, "TestAgent/1.0", ua)
	})

	t.Run("GetHeaders", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/", nil)
		req.Header.Set("X-Test", "Value")
		
		headers := service.GetHeaders(req)
		assert.Equal(t, "Value", headers["X-Test"][0])
	})
}
