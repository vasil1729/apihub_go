package auth

import (
	"github.com/gin-gonic/gin"
)

func SetupAuthRoutes(router *gin.RouterGroup, handler *AuthHandler) {
	router.POST("/register", handler.Register)
	router.POST("/login", handler.Login)
	router.POST("/logout", handler.Logout)
}
