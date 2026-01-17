package kitchensink_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/ultimatum/apihub_go/internal/service/kitchensink"
)

func TestResponseInspectionService(t *testing.T) {
	service := kitchensink.NewResponseInspectionService()

	t.Run("GetJSONResponse", func(t *testing.T) {
		resp := service.GetJSONResponse()
		assert.Equal(t, "This is a JSON response", resp.Message)
		assert.Equal(t, "json", resp.Format)
	})

	t.Run("GetXMLResponse", func(t *testing.T) {
		resp := service.GetXMLResponse()
		assert.Equal(t, "This is an XML response", resp.Message)
		assert.Equal(t, "xml", resp.Format)
	})

	t.Run("GetHTMLResponse", func(t *testing.T) {
		resp := service.GetHTMLResponse()
		assert.True(t, strings.Contains(resp, "<!DOCTYPE html>"))
		assert.True(t, strings.Contains(resp, "<h1>This is an HTML response</h1>"))
	})
}
