package public

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ultimatum/apihub_go/internal/service/public"
	"github.com/ultimatum/apihub_go/pkg/response"
)

// RandomJokeHandler handles random joke HTTP requests
type RandomJokeHandler struct {
	service *public.RandomJokeService
}

// NewRandomJokeHandler creates a new random joke handler
func NewRandomJokeHandler(service *public.RandomJokeService) *RandomJokeHandler {
	return &RandomJokeHandler{
		service: service,
	}
}

// GetAll godoc
// @Summary Get all random jokes
// @Description Get a paginated list of random jokes
// @Tags Public APIs
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Items per page" default(10)
// @Success 200 {object} response.PaginatedResponse
// @Router /public/randomjokes [get]
func (h *RandomJokeHandler) GetAll(c *gin.Context) {
	// Parse query parameters
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	
	// Validate parameters
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}
	
	// Get jokes
	jokes, total, err := h.service.GetAll(page, limit)
	if err != nil {
		response.InternalServerError(c, "Failed to fetch jokes")
		return
	}
	
	// Calculate pagination metadata
	totalPages := (total + limit - 1) / limit
	
	pagination := response.Pagination{
		Page:       page,
		Limit:      limit,
		TotalPages: totalPages,
		TotalItems: int64(total),
		HasNext:    page < totalPages,
		HasPrev:    page > 1,
	}
	
	response.SuccessWithPagination(c, "Random jokes fetched successfully", jokes, pagination)
}

// GetByID godoc
// @Summary Get random joke by ID
// @Description Get a specific random joke by its ID
// @Tags Public APIs
// @Accept json
// @Produce json
// @Param id path int true "Joke ID"
// @Success 200 {object} response.Response
// @Failure 404 {object} response.Response
// @Router /public/randomjokes/{id} [get]
func (h *RandomJokeHandler) GetByID(c *gin.Context) {
	// Parse ID
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.BadRequest(c, "Invalid joke ID")
		return
	}
	
	// Get joke
	joke, err := h.service.GetByID(id)
	if err != nil {
		response.NotFound(c, "Joke not found")
		return
	}
	
	response.OK(c, "Joke fetched successfully", joke)
}

// GetRandom godoc
// @Summary Get a random joke
// @Description Get a randomly selected joke
// @Tags Public APIs
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Router /public/randomjokes/random [get]
func (h *RandomJokeHandler) GetRandom(c *gin.Context) {
	joke, err := h.service.GetRandom()
	if err != nil {
		response.InternalServerError(c, "Failed to fetch random joke")
		return
	}
	
	response.OK(c, "Random joke fetched successfully", joke)
}
