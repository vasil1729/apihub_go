package kitchensink

import (
	"net/http"

	"github.com/gin-gonic/gin"
	kitchenSinkService "github.com/ultimatum/apihub_go/internal/service/kitchensink"
)

type ResponseInspectionHandler struct {
	service *kitchenSinkService.ResponseInspectionService
}

func NewResponseInspectionHandler(service *kitchenSinkService.ResponseInspectionService) *ResponseInspectionHandler {
	return &ResponseInspectionHandler{service: service}
}

// @Summary Get JSON response
// @Tags Kitchen Sink
// @Produce json
// @Success 200 {object} kitchensink.ResponseInspectionResponse
// @Router /kitchen-sink/response/json [get]
func (h *ResponseInspectionHandler) GetJSON(c *gin.Context) {
	resp := h.service.GetJSONResponse()
	c.JSON(http.StatusOK, resp)
}

// @Summary Get XML response
// @Tags Kitchen Sink
// @Produce xml
// @Success 200 {object} kitchensink.ResponseInspectionResponse
// @Router /kitchen-sink/response/xml [get]
func (h *ResponseInspectionHandler) GetXML(c *gin.Context) {
	resp := h.service.GetXMLResponse()
	c.XML(http.StatusOK, resp)
}

// @Summary Get HTML response
// @Tags Kitchen Sink
// @Produce html
// @Success 200 {string} string "HTML Response"
// @Router /kitchen-sink/response/html [get]
func (h *ResponseInspectionHandler) GetHTML(c *gin.Context) {
	resp := h.service.GetHTMLResponse()
	c.Header("Content-Type", "text/html; charset=utf-8")
	c.String(http.StatusOK, resp)
}
