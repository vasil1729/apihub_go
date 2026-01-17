package kitchensink

import (
	"net/http"

	"github.com/ultimatum/apihub_go/internal/domain/kitchensink"
)

type CookiesService struct{}

func NewCookiesService() *CookiesService {
	return &CookiesService{}
}

func (s *CookiesService) GetCookies(r *http.Request) *kitchensink.CookieResponse {
	cookies := make(map[string]string)
	for _, cookie := range r.Cookies() {
		cookies[cookie.Name] = cookie.Value
	}
	return &kitchensink.CookieResponse{
		Cookies: cookies,
	}
}
