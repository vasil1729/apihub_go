package kitchensink_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/ultimatum/apihub_go/internal/service/kitchensink"
)

func TestStatusCodesService_GetResponseForCode(t *testing.T) {
	service := kitchensink.NewStatusCodesService()

	tests := []struct {
		name        string
		code        int
		expectError bool
		expectMsg   string
	}{
		{"Valid 200", 200, false, "OK"},
		{"Valid 404", 404, false, "Not Found"},
		{"Valid 500", 500, false, "Internal Server Error"},
		{"Valid 418", 418, false, "I'm a teapot"},
		{"Invalid Code", 999, true, ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := service.GetResponseForCode(tt.code)
			
			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.code, resp.Code)
				assert.Equal(t, tt.expectMsg, resp.Message)
			}
		})
	}
}
