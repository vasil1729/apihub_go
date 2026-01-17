package public

import (
	"github.com/gin-gonic/gin"
	publicService "github.com/ultimatum/apihub_go/internal/service/public"
)

// SetupPublicRoutes sets up all public API routes
func SetupPublicRoutes(router *gin.RouterGroup, dataPath string) error {
	// Initialize services
	randomUserService, err := publicService.NewRandomUserService(dataPath)
	if err != nil {
		return err
	}
	
	// Initialize handlers
	randomUserHandler := NewRandomUserHandler(randomUserService)
	
	// Random Users routes
	randomUsers := router.Group("/randomusers")
	{
		randomUsers.GET("", randomUserHandler.GetAll)
		randomUsers.GET("/random", randomUserHandler.GetRandom)
		randomUsers.GET("/:id", randomUserHandler.GetByID)
	}
	
	return nil
}
