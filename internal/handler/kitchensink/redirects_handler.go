package kitchensink

import (
	"github.com/gin-gonic/gin"
	kitchenSinkService "github.com/ultimatum/apihub_go/internal/service/kitchensink"
	"net/http"
)

type RedirectsHandler struct {
	service *kitchenSinkService.RedirectsService
}

func NewRedirectsHandler(service *kitchenSinkService.RedirectsService) *RedirectsHandler {
	return &RedirectsHandler{service: service}
}

// @Summary Trigger a 301 Moved Permanently redirect
// @Tags Kitchen Sink
// @Param url query string false "URL to redirect to"
// @Success 301 {string} string "Redirecting..."
// @Router /kitchen-sink/redirects/301 [get]
func (h *RedirectsHandler) Redirect301(c *gin.Context) {
	to := c.Query("url")
	url := h.service.GetRedirectURL(to)
	c.Redirect(http.StatusMovedPermanently, url)
}

// @Summary Trigger a 302 Found redirect
// @Tags Kitchen Sink
// @Param url query string false "URL to redirect to"
// @Success 302 {string} string "Redirecting..."
// @Router /kitchen-sink/redirects/302 [get]
func (h *RedirectsHandler) Redirect302(c *gin.Context) {
	to := c.Query("url")
	url := h.service.GetRedirectURL(to)
	c.Redirect(http.StatusFound, url)
}

// @Summary Trigger a 307 Temporary Redirect
// @Tags Kitchen Sink
// @Param url query string false "URL to redirect to"
// @Success 307 {string} string "Redirecting..."
// @Router /kitchen-sink/redirects/307 [get]
func (h *RedirectsHandler) Redirect307(c *gin.Context) {
	to := c.Query("url")
	url := h.service.GetRedirectURL(to)
	c.Redirect(http.StatusTemporaryRedirect, url)
}

// @Summary Trigger a 308 Permanent Redirect
// @Tags Kitchen Sink
// @Param url query string false "URL to redirect to"
// @Success 308 {string} string "Redirecting..."
// @Router /kitchen-sink/redirects/308 [get]
func (h *RedirectsHandler) Redirect308(c *gin.Context) {
	to := c.Query("url")
	url := h.service.GetRedirectURL(to)
	c.Redirect(http.StatusPermanentRedirect, url)
}
