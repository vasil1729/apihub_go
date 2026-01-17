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
	
	// Initialize joke service and handler
	randomJokeService, err := publicService.NewRandomJokeService(dataPath)
	if err != nil {
		return err
	}
	randomJokeHandler := NewRandomJokeHandler(randomJokeService)
	
	// Initialize quote service and handler
	quoteService, err := publicService.NewQuoteService(dataPath)
	if err != nil {
		return err
	}
	quoteHandler := NewQuoteHandler(quoteService)
	
	// Initialize stock service and handler
	stockService, err := publicService.NewStockService(dataPath)
	if err != nil {
		return err
	}
	stockHandler := NewStockHandler(stockService)
	
	// Random Users routes
	randomUsers := router.Group("/randomusers")
	{
		randomUsers.GET("", randomUserHandler.GetAll)
		randomUsers.GET("/random", randomUserHandler.GetRandom)
		randomUsers.GET("/:id", randomUserHandler.GetByID)
	}
	
	// Random Jokes routes
	randomJokes := router.Group("/randomjokes")
	{
		randomJokes.GET("", randomJokeHandler.GetAll)
		randomJokes.GET("/random", randomJokeHandler.GetRandom)
		randomJokes.GET("/:id", randomJokeHandler.GetByID)
	}
	
	// Quotes routes
	quotes := router.Group("/quotes")
	{
		quotes.GET("", quoteHandler.GetAll)
		quotes.GET("/random", quoteHandler.GetRandom)
		quotes.GET("/:id", quoteHandler.GetByID)
	}
	
	// Stocks routes
	stocks := router.Group("/stocks")
	{
		stocks.GET("", stockHandler.GetAll)
		stocks.GET("/:symbol", stockHandler.GetBySymbol)
	}
	
	return nil
}
