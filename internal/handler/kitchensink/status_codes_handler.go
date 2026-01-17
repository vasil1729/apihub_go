package kitchensink

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ultimatum/apihub_go/internal/service/kitchensink"
	"github.com/ultimatum/apihub_go/pkg/response"
)

type StatusCodesHandler struct {
	service *kitchensink.StatusCodesService
}

func NewStatusCodesHandler(service *kitchensink.StatusCodesService) *StatusCodesHandler {
	return &StatusCodesHandler{service: service}
}

// @Summary Return response with specified status code
// @Tags Kitchen Sink
// @Param code path int true "HTTP Status Code"
// @Success 200 {object} kitchensink.StatusCodeResponse
// @Failure 400 {object} response.Response
// @Router /kitchen-sink/status/{code} [get]
// @Router /kitchen-sink/status/{code} [post]
// @Router /kitchen-sink/status/{code} [put]
// @Router /kitchen-sink/status/{code} [patch]
// @Router /kitchen-sink/status/{code} [delete]
func (h *StatusCodesHandler) HandleStatus(c *gin.Context) {
	codeStr := c.Param("code")
	code, err := strconv.Atoi(codeStr)
	if err != nil {
		response.BadRequest(c, "Invalid status code format")
		return
	}

	resp, err := h.service.GetResponseForCode(code)
	if err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	// We return the requested status code directly
	// We also return a JSON body explaining the code
	c.JSON(code, resp)
}
