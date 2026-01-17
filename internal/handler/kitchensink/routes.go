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
	
	return nil
}
