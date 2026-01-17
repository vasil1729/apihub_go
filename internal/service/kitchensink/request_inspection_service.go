package kitchensink

import (
	"net/http"

	"github.com/ultimatum/apihub_go/internal/domain/kitchensink"
)

type RequestInspectionService struct{}

func NewRequestInspectionService() *RequestInspectionService {
	return &RequestInspectionService{}
}

func (s *RequestInspectionService) InspectRequest(r *http.Request) *kitchensink.RequestInspectionResponse {
	return &kitchensink.RequestInspectionResponse{
		IP:        s.GetClientIP(r),
		UserAgent: r.UserAgent(),
		Headers:   r.Header,
		Method:    r.Method,
		URL:       r.URL.String(),
	}
}

func (s *RequestInspectionService) GetClientIP(r *http.Request) string {
	// Simple IP extraction, could be improved to check X-Forwarded-For etc.
	// For now, RemoteAddr or ClientIP from framework/headers
	ip := r.Header.Get("X-Forwarded-For")
	if ip == "" {
		ip = r.RemoteAddr
	}
	return ip
}

func (s *RequestInspectionService) GetUserAgent(r *http.Request) string {
	return r.UserAgent()
}

func (s *RequestInspectionService) GetHeaders(r *http.Request) map[string][]string {
	return r.Header
}
