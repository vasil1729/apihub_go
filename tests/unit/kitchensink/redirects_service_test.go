package kitchensink_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/ultimatum/apihub_go/internal/service/kitchensink"
)

func TestRedirectsService_GetRedirectURL(t *testing.T) {
	service := kitchensink.NewRedirectsService()

	t.Run("Default URL", func(t *testing.T) {
		url := service.GetRedirectURL("")
		assert.Equal(t, "https://google.com", url)
	})

	t.Run("Custom URL", func(t *testing.T) {
		custom := "https://example.com"
		url := service.GetRedirectURL(custom)
		assert.Equal(t, custom, url)
	})
}
