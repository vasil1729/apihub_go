package kitchensink_test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/ultimatum/apihub_go/internal/service/kitchensink"
	"net/http/httptest"
)

func TestCookiesService_GetCookies(t *testing.T) {
	service := kitchensink.NewCookiesService()

	req := httptest.NewRequest("GET", "/", nil)
	req.AddCookie(&http.Cookie{Name: "test-cookie", Value: "test-value"})
	req.AddCookie(&http.Cookie{Name: "another-cookie", Value: "another-value"})

	resp := service.GetCookies(req)
	
	assert.Equal(t, "test-value", resp.Cookies["test-cookie"])
	assert.Equal(t, "another-value", resp.Cookies["another-cookie"])
	assert.Len(t, resp.Cookies, 2)
}
