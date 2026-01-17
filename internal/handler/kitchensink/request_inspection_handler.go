package kitchensink

import (
	"github.com/gin-gonic/gin"
	kitchenSinkDomain "github.com/ultimatum/apihub_go/internal/domain/kitchensink"
	kitchenSinkService "github.com/ultimatum/apihub_go/internal/service/kitchensink"
	"github.com/ultimatum/apihub_go/pkg/response"
)

type RequestInspectionHandler struct {
	service *kitchenSinkService.RequestInspectionService
}

func NewRequestInspectionHandler(service *kitchenSinkService.RequestInspectionService) *RequestInspectionHandler {
	return &RequestInspectionHandler{service: service}
}

// @Summary Inspect IP address
// @Tags Kitchen Sink
// @Success 200 {object} response.Response
// @Router /kitchen-sink/ip [get]
func (h *RequestInspectionHandler) GetIP(c *gin.Context) {
	ip := h.service.GetClientIP(c.Request)
	// If behind a proxy, Gin has ClientIP() method which is better
	if c.ClientIP() != "" {
		ip = c.ClientIP()
	}
	
	resp := kitchenSinkDomain.IPResponse{IP: ip}
	response.OK(c, "Client IP fetched successfully", resp)
}

// @Summary Inspect User Agent
// @Tags Kitchen Sink
// @Success 200 {object} response.Response
// @Router /kitchen-sink/user-agent [get]
func (h *RequestInspectionHandler) GetUserAgent(c *gin.Context) {
	ua := h.service.GetUserAgent(c.Request)
	resp := kitchenSinkDomain.UserAgentResponse{UserAgent: ua}
	response.OK(c, "User Agent fetched successfully", resp)
}

// @Summary Inspect Headers
// @Tags Kitchen Sink
// @Success 200 {object} response.Response
// @Router /kitchen-sink/headers [get]
func (h *RequestInspectionHandler) GetHeaders(c *gin.Context) {
	headers := h.service.GetHeaders(c.Request)
	resp := kitchenSinkDomain.HeadersResponse{Headers: headers}
	response.OK(c, "Headers fetched successfully", resp)
}
