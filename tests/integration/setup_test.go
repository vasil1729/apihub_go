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
	// Note: We might need a dummy DB connection for public routes if they used it?
	// Public routes currently don't use DB in their handler constructor in main.go (wait, main.go passes services).
	// But `SetupPublicRoutes` in `internal/handler/public/routes.go` creates handlers internally?
	// Let's check `public/routes.go`.
	// Actually `setupTestRouter` currently calls `publicHandler.SetupPublicRoutes`.
	// If public handlers need services that need DB, `SetupPublicRoutes` might be creating them?
	// `SetupPublicRoutes` in `routes.go` instantiates minimal services?
	// Let's look at `setup_test.go` again.
	// It calls `publicHandler.SetupPublicRoutes(publicRoutes, "../../data")`.
	
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
