package kitchensink

import (
	"github.com/gin-gonic/gin"
	"github.com/ultimatum/apihub_go/internal/service/kitchensink"
	"github.com/ultimatum/apihub_go/pkg/response"
)

type HTTPMethodsHandler struct {
	service *kitchensink.HTTPMethodsService
}

func NewHTTPMethodsHandler(service *kitchensink.HTTPMethodsService) *HTTPMethodsHandler {
	return &HTTPMethodsHandler{service: service}
}

// @Summary Test GET method
// @Tags Kitchen Sink
// @Success 200 {object} response.Response
// @Router /kitchen-sink/http-methods/get [get]
func (h *HTTPMethodsHandler) HandleGet(c *gin.Context) {
	resp, _ := h.service.ProcessRequest(c.Request)
	response.OK(c, "GET request received successfully", resp)
}

// @Summary Test POST method
// @Tags Kitchen Sink
// @Success 200 {object} response.Response
// @Router /kitchen-sink/http-methods/post [post]
func (h *HTTPMethodsHandler) HandlePost(c *gin.Context) {
	// We need to read body from Gin context if not read by binding
	// Ideally service should take req info, but passing request is fine for this demo
	resp, _ := h.service.ProcessRequest(c.Request)
	response.OK(c, "POST request received successfully", resp)
}

// @Summary Test PUT method
// @Tags Kitchen Sink
// @Success 200 {object} response.Response
// @Router /kitchen-sink/http-methods/put [put]
func (h *HTTPMethodsHandler) HandlePut(c *gin.Context) {
	resp, _ := h.service.ProcessRequest(c.Request)
	response.OK(c, "PUT request received successfully", resp)
}

// @Summary Test PATCH method
// @Tags Kitchen Sink
// @Success 200 {object} response.Response
// @Router /kitchen-sink/http-methods/patch [patch]
func (h *HTTPMethodsHandler) HandlePatch(c *gin.Context) {
	resp, _ := h.service.ProcessRequest(c.Request)
	response.OK(c, "PATCH request received successfully", resp)
}

// @Summary Test DELETE method
// @Tags Kitchen Sink
// @Success 200 {object} response.Response
// @Router /kitchen-sink/http-methods/delete [delete]
func (h *HTTPMethodsHandler) HandleDelete(c *gin.Context) {
	resp, _ := h.service.ProcessRequest(c.Request)
	response.OK(c, "DELETE request received successfully", resp)
}
