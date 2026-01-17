package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	publicHandler "github.com/ultimatum/apihub_go/internal/handler/public"
	kitchenSinkHandler "github.com/ultimatum/apihub_go/internal/handler/kitchensink"
	authHandlerPkg "github.com/ultimatum/apihub_go/internal/handler/auth"
	authServicePkg "github.com/ultimatum/apihub_go/internal/service/auth"
	"github.com/ultimatum/apihub_go/internal/middleware"
	"github.com/ultimatum/apihub_go/pkg/config"
	"github.com/ultimatum/apihub_go/pkg/database"
	"github.com/ultimatum/apihub_go/pkg/logger"
	"github.com/ultimatum/apihub_go/pkg/response"
	
	_ "github.com/ultimatum/apihub_go/docs" // Import generated docs
)

// @title FreeAPI - API Hub
// @version 1.0
// @description A comprehensive API hub with multiple API categories for learning and development
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url https://github.com/ultimatum/apihub_go
// @contact.email support@apihub.com

// @license.name ISC
// @license.url https://opensource.org/licenses/ISC

// @host localhost:8080
// @BasePath /api/v1

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Initialize logger
	logger.Init(cfg.LogLevel, cfg.LogFormat)
	logger.Info("Starting FreeAPI server...")

	// Connect to MongoDB
	db, err := database.Connect(cfg.MongoDBURI, cfg.DBName)
	if err != nil {
		logger.Fatal("Failed to connect to database", err)
	}
	defer db.Disconnect()
	logger.Info("Connected to MongoDB successfully")

	// Set Gin mode
	if cfg.IsProduction() {
		gin.SetMode(gin.ReleaseMode)
	}

	// Create Gin router
	router := gin.New()

	// Global middleware
	router.Use(gin.Recovery())
	router.Use(middleware.Logger())
	router.Use(middleware.CORS(cfg))
	router.Use(middleware.RateLimit(cfg))

	// Health check endpoint
	router.GET("/api/v1/healthcheck", func(c *gin.Context) {
		// Check database health
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := db.HealthCheck(ctx); err != nil {
			response.InternalServerError(c, "Database connection failed")
			return
		}

		response.OK(c, "Server is healthy", gin.H{
			"status":   "ok",
			"database": "connected",
		})
	})

	// API v1 routes
	v1 := router.Group("/api/v1")
	
	// Public APIs
	publicRoutes := v1.Group("/public")
	if err := publicHandler.SetupPublicRoutes(publicRoutes, cfg.DataPath); err != nil {
		log.Fatalf("Failed to setup public routes: %v", err)
	}

	// Kitchen Sink APIs
	kitchenSinkRoutes := v1.Group("/kitchen-sink")
	if err := kitchenSinkHandler.SetupKitchenSinkRoutes(kitchenSinkRoutes); err != nil {
		log.Fatalf("Failed to setup kitchen sink routes: %v", err)
	}

	// Auth APIs
	authService := authServicePkg.NewAuthService(db)
	authHandler := authHandlerPkg.NewAuthHandler(authService)
	authHandlerPkg.SetupAuthRoutes(v1.Group("/auth"), authHandler)

	// Swagger documentation
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	
	// Root redirect to Swagger
	router.GET("/", func(c *gin.Context) {
		c.Redirect(302, "/swagger/index.html")
	})

	// 404 handler
	router.NoRoute(func(c *gin.Context) {
		response.NotFound(c, "Route not found")
	})

	// Error handling middleware (should be last)
	router.Use(middleware.ErrorHandler())

	// Create HTTP server
	srv := &http.Server{
		Addr:    ":" + cfg.Port,
		Handler: router,
	}

	// Start server in a goroutine
	go func() {
		logger.Info(fmt.Sprintf("Server starting on port %s", cfg.Port))
		logger.Info(fmt.Sprintf("Environment: %s", cfg.Env))
		logger.Info(fmt.Sprintf("API Documentation: %s/swagger/index.html", cfg.HostURL))

		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatal("Failed to start server", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logger.Info("Shutting down server...")

	// Graceful shutdown with 5 second timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logger.Fatal("Server forced to shutdown", err)
	}

	logger.Info("Server exited successfully")
}
