package kitchensink_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/ultimatum/apihub_go/internal/service/kitchensink"
)

func TestImagesService_GenerateImages(t *testing.T) {
	service := kitchensink.NewImagesService()

	t.Run("GenerateJPEG", func(t *testing.T) {
		data, err := service.GenerateJPEG()
		assert.NoError(t, err)
		assert.NotEmpty(t, data)
		// Check Magic Bytes for JPEG (FF D8)
		assert.Equal(t, uint8(0xFF), data[0])
		assert.Equal(t, uint8(0xD8), data[1])
	})

	t.Run("GeneratePNG", func(t *testing.T) {
		data, err := service.GeneratePNG()
		assert.NoError(t, err)
		assert.NotEmpty(t, data)
		// Check Magic Bytes for PNG (89 50 4E 47 0D 0A 1A 0A)
		assert.Equal(t, uint8(0x89), data[0])
		assert.Equal(t, uint8(0x50), data[1])
	})

	t.Run("GenerateSVG", func(t *testing.T) {
		svg := service.GenerateSVG()
		assert.NotEmpty(t, svg)
		assert.True(t, strings.Contains(svg, "<svg"))
		assert.True(t, strings.Contains(svg, "xmlns=\"http://www.w3.org/2000/svg\""))
	})

	t.Run("GetWebP", func(t *testing.T) {
		data := service.GetWebP()
		assert.NotEmpty(t, data)
		// Check WebP / RIFF header
		assert.Equal(t, uint8(0x52), data[0]) // R
		assert.Equal(t, uint8(0x49), data[1]) // I
	})
}
