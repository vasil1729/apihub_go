package kitchensink

import (
	"github.com/gin-gonic/gin"
	kitchenSinkDomain "github.com/ultimatum/apihub_go/internal/domain/kitchensink"
	kitchenSinkService "github.com/ultimatum/apihub_go/internal/service/kitchensink"
	"github.com/ultimatum/apihub_go/pkg/response"
)

type CookiesHandler struct {
	service *kitchenSinkService.CookiesService
}

func NewCookiesHandler(service *kitchenSinkService.CookiesService) *CookiesHandler {
	return &CookiesHandler{service: service}
}

// @Summary Get all cookies
// @Tags Kitchen Sink
// @Success 200 {object} response.Response
// @Router /kitchen-sink/cookies/get [get]
func (h *CookiesHandler) GetCookies(c *gin.Context) {
	cookies := h.service.GetCookies(c.Request)
	response.OK(c, "Cookies fetched successfully", cookies)
}

// @Summary Set a cookie
// @Tags Kitchen Sink
// @Param key query string true "Cookie Key"
// @Param value query string true "Cookie Value"
// @Success 200 {object} response.Response
// @Router /kitchen-sink/cookies/set [get]
func (h *CookiesHandler) SetCookie(c *gin.Context) {
	key := c.Query("key")
	value := c.Query("value")
	
	if key == "" || value == "" {
		response.BadRequest(c, "key and value query parameters are required")
		return
	}

	// Set cookie with HttpOnly=false so it can be accessed by JS, basic path /, max age 3600
	c.SetCookie(key, value, 3600, "/", "", false, false)
	
	cookies := h.service.GetCookies(c.Request)
	// Add the new/updated cookie manually since request cookies won't contain it yet in this specific request object typically
	// But let's just return success message with the set values
	responseData := map[string]string{
		key: value,
	}
	// Also returning other cookies
	for k, v := range cookies.Cookies {
		responseData[k] = v
	}
	// Override/Ensure the new one is there for display
	responseData[key] = value

	response.OK(c, "Cookie set successfully", kitchenSinkDomain.CookieResponse{Cookies: responseData})
}

// @Summary Delete a cookie
// @Tags Kitchen Sink
// @Param key query string true "Cookie Key"
// @Success 200 {object} response.Response
// @Router /kitchen-sink/cookies/delete [get]
func (h *CookiesHandler) DeleteCookie(c *gin.Context) {
	key := c.Query("key")
	if key == "" {
		response.BadRequest(c, "key query parameter is required")
		return
	}

	// Delete cookie by setting max age to -1
	c.SetCookie(key, "", -1, "/", "", false, false)
	
	cookies := h.service.GetCookies(c.Request)
	// Remove from response data for display
	delete(cookies.Cookies, key)
	
	response.OK(c, "Cookie deleted successfully", cookies)
}
