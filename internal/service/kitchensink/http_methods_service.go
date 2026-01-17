package kitchensink

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/ultimatum/apihub_go/internal/domain/kitchensink"
)

type HTTPMethodsService struct{}

func NewHTTPMethodsService() *HTTPMethodsService {
	return &HTTPMethodsService{}
}

func (s *HTTPMethodsService) ProcessRequest(r *http.Request) (*kitchensink.HTTPMethodResponse, error) {
	// Parse body if present
	var body interface{}
	if r.Body != nil {
		bodyBytes, err := io.ReadAll(r.Body)
		if err == nil && len(bodyBytes) > 0 {
			// Try to parse as JSON first
			if err := json.Unmarshal(bodyBytes, &body); err != nil {
				// If not JSON, use string
				body = string(bodyBytes)
			}
		}
	}

	response := &kitchensink.HTTPMethodResponse{
		Method:        r.Method,
		URL:           r.URL.String(),
		Headers:       r.Header,
		Query:         r.URL.Query(),
		Body:          body,
		Origin:        r.RemoteAddr,
		ContentLength: r.ContentLength,
		UserAgent:     r.UserAgent(),
	}

	return response, nil
}
