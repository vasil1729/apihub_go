package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/ultimatum/apihub_go/pkg/config"
)

// CORS returns a CORS middleware
func CORS(cfg *config.Config) gin.HandlerFunc {
	config := cors.DefaultConfig()
	
	if len(cfg.CORSOrigin) == 1 && cfg.CORSOrigin[0] == "*" {
		config.AllowAllOrigins = true
	} else {
		config.AllowOrigins = cfg.CORSOrigin
	}
	
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Accept", "Authorization"}
	config.AllowCredentials = true
	
	return cors.New(config)
}
