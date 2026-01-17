package kitchensink

import (
	"fmt"
	"net/http"

	"github.com/ultimatum/apihub_go/internal/domain/kitchensink"
)

type StatusCodesService struct{}

func NewStatusCodesService() *StatusCodesService {
	return &StatusCodesService{}
}

func (s *StatusCodesService) GetResponseForCode(code int) (*kitchensink.StatusCodeResponse, error) {
	statusText := http.StatusText(code)
	if statusText == "" {
		return nil, fmt.Errorf("unknown status code: %d", code)
	}

	return &kitchensink.StatusCodeResponse{
		Code:    code,
		Message: statusText,
	}, nil
}
