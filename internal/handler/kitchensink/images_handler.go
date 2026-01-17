package kitchensink

import (
	"github.com/gin-gonic/gin"
	kitchenSinkService "github.com/ultimatum/apihub_go/internal/service/kitchensink"
	"github.com/ultimatum/apihub_go/pkg/response"
	"net/http"
)

type ImagesHandler struct {
	service *kitchenSinkService.ImagesService
}

func NewImagesHandler(service *kitchenSinkService.ImagesService) *ImagesHandler {
	return &ImagesHandler{service: service}
}

// @Summary Get a generated JPEG image
// @Tags Kitchen Sink
// @Produce image/jpeg
// @Success 200 {string} binary "JPEG Image"
// @Router /kitchen-sink/images/jpeg [get]
func (h *ImagesHandler) GetJPEG(c *gin.Context) {
	imgData, err := h.service.GenerateJPEG()
	if err != nil {
		response.InternalServerError(c, "Failed to generate JPEG")
		return
	}
	c.Data(http.StatusOK, "image/jpeg", imgData)
}

// @Summary Get a generated PNG image
// @Tags Kitchen Sink
// @Produce image/png
// @Success 200 {string} binary "PNG Image"
// @Router /kitchen-sink/images/png [get]
func (h *ImagesHandler) GetPNG(c *gin.Context) {
	imgData, err := h.service.GeneratePNG()
	if err != nil {
		response.InternalServerError(c, "Failed to generate PNG")
		return
	}
	c.Data(http.StatusOK, "image/png", imgData)
}

// @Summary Get a generated SVG image
// @Tags Kitchen Sink
// @Produce image/svg+xml
// @Success 200 {string} string "SVG Image"
// @Router /kitchen-sink/images/svg [get]
func (h *ImagesHandler) GetSVG(c *gin.Context) {
	svg := h.service.GenerateSVG()
	c.Header("Content-Type", "image/svg+xml")
	c.String(http.StatusOK, svg)
}

// @Summary Get a static WebP image
// @Tags Kitchen Sink
// @Produce image/webp
// @Success 200 {string} binary "WebP Image"
// @Router /kitchen-sink/images/webp [get]
func (h *ImagesHandler) GetWebP(c *gin.Context) {
	webpData := h.service.GetWebP()
	c.Data(http.StatusOK, "image/webp", webpData)
}
