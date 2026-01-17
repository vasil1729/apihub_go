package integration_test

import (
	"github.com/gin-gonic/gin"
	publicHandler "github.com/ultimatum/apihub_go/internal/handler/public"
	kitchenSinkHandler "github.com/ultimatum/apihub_go/internal/handler/kitchensink"
)

func setupTestRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	
	v1 := router.Group("/api/v1")
	
	// Setup public routes
	publicRoutes := v1.Group("/public")
	err := publicHandler.SetupPublicRoutes(publicRoutes, "../../data")
	if err != nil {
		panic(err)
	}
	
	// Setup kitchen sink routes
	kitchenSinkRoutes := v1.Group("/kitchen-sink")
	err = kitchenSinkHandler.SetupKitchenSinkRoutes(kitchenSinkRoutes)
	if err != nil {
		panic(err)
	}
	
	return router
}
