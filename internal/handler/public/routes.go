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
	
	// Initialize meal service and handler
	mealService, err := publicService.NewMealService(dataPath)
	if err != nil {
		return err
	}
	mealHandler := NewMealHandler(mealService)
	
	// Initialize dog service and handler
	dogService, err := publicService.NewDogService(dataPath)
	if err != nil {
		return err
	}
	dogHandler := NewDogHandler(dogService)
	
	// Initialize cat service and handler
	catService, err := publicService.NewCatService(dataPath)
	if err != nil {
		return err
	}
	catHandler := NewCatHandler(catService)
	
	// Initialize book service and handler
	bookService, err := publicService.NewBookService(dataPath)
	if err != nil {
		return err
	}
	bookHandler := NewBookHandler(bookService)
	
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
	
	// Meals routes
	meals := router.Group("/meals")
	{
		meals.GET("", mealHandler.GetAll)
		meals.GET("/random", mealHandler.GetRandom)
		meals.GET("/:id", mealHandler.GetByID)
	}
	
	// Dogs routes
	dogs := router.Group("/dogs")
	{
		dogs.GET("", dogHandler.GetAll)
		dogs.GET("/random", dogHandler.GetRandom)
		dogs.GET("/:id", dogHandler.GetByID)
	}
	
	// Cats routes
	cats := router.Group("/cats")
	{
		cats.GET("", catHandler.GetAll)
		cats.GET("/random", catHandler.GetRandom)
		cats.GET("/:id", catHandler.GetByID)
	}
	
	// Books routes
	books := router.Group("/books")
	{
		books.GET("", bookHandler.GetAll)
		books.GET("/random", bookHandler.GetRandom)
		books.GET("/:id", bookHandler.GetByID)
	}
	
	return nil
}
