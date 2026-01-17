package kitchensink

import (
	"github.com/gin-gonic/gin"
	kitchenSinkService "github.com/ultimatum/apihub_go/internal/service/kitchensink"
)

// SetupKitchenSinkRoutes sets up all kitchen sink API routes
func SetupKitchenSinkRoutes(router *gin.RouterGroup) error {
	// Initialize HTTP Methods service and handler
	httpMethodsService := kitchenSinkService.NewHTTPMethodsService()
	httpMethodsHandler := NewHTTPMethodsHandler(httpMethodsService)
	
	// Initialize Status Codes service and handler
	statusCodesService := kitchenSinkService.NewStatusCodesService()
	statusCodesHandler := NewStatusCodesHandler(statusCodesService)

	// Initialize Request Inspection service and handler
	reqInspectionService := kitchenSinkService.NewRequestInspectionService()
	reqInspectionHandler := NewRequestInspectionHandler(reqInspectionService)
	
	// Initialize Response Inspection service and handler
	respInspectionService := kitchenSinkService.NewResponseInspectionService()
	respInspectionHandler := NewResponseInspectionHandler(respInspectionService)
	
	// Initialize Cookies service and handler
	cookiesService := kitchenSinkService.NewCookiesService()
	cookiesHandler := NewCookiesHandler(cookiesService)
	
	// Initialize Redirects service and handler
	redirectsService := kitchenSinkService.NewRedirectsService()
	redirectsHandler := NewRedirectsHandler(redirectsService)

	httpMethods := router.Group("/http-methods")
	{
		httpMethods.GET("/get", httpMethodsHandler.HandleGet)
		httpMethods.POST("/post", httpMethodsHandler.HandlePost)
		httpMethods.PUT("/put", httpMethodsHandler.HandlePut)
		httpMethods.PATCH("/patch", httpMethodsHandler.HandlePatch)
		httpMethods.DELETE("/delete", httpMethodsHandler.HandleDelete)
	}

	statusCodes := router.Group("/status")
	{
		statusCodes.Any("/:code", statusCodesHandler.HandleStatus)
	}
	
	// Request Inspection endpoints
	router.GET("/ip", reqInspectionHandler.GetIP)
	router.GET("/user-agent", reqInspectionHandler.GetUserAgent)
	router.GET("/headers", reqInspectionHandler.GetHeaders)
	
	// Response Inspection endpoints
	responseGroup := router.Group("/response")
	{
		responseGroup.GET("/json", respInspectionHandler.GetJSON)
		responseGroup.GET("/xml", respInspectionHandler.GetXML)
		responseGroup.GET("/html", respInspectionHandler.GetHTML)
	}
	
	// Cookies endpoints
	cookiesGroup := router.Group("/cookies")
	{
		cookiesGroup.GET("/get", cookiesHandler.GetCookies)
		cookiesGroup.GET("/set", cookiesHandler.SetCookie)
		cookiesGroup.GET("/delete", cookiesHandler.DeleteCookie)
	}

	// Redirects endpoints
	redirectsGroup := router.Group("/redirects")
	{
		redirectsGroup.GET("/301", redirectsHandler.Redirect301)
		redirectsGroup.GET("/302", redirectsHandler.Redirect302)
		redirectsGroup.GET("/307", redirectsHandler.Redirect307)
		redirectsGroup.GET("/308", redirectsHandler.Redirect308)
	}
	
	return nil
}
