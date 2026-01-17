package public

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ultimatum/apihub_go/internal/service/public"
	"github.com/ultimatum/apihub_go/pkg/response"
)

// DogHandler handles dog HTTP requests
type DogHandler struct {
	service *public.DogService
}

// NewDogHandler creates a new dog handler
func NewDogHandler(service *public.DogService) *DogHandler {
	return &DogHandler{
		service: service,
	}
}

// GetAll godoc
// @Summary Get all dog breeds
// @Description Get a paginated list of dog breeds
// @Tags Public APIs
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Items per page" default(10)
// @Success 200 {object} response.PaginatedResponse
// @Router /public/dogs [get]
func (h *DogHandler) GetAll(c *gin.Context) {
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
	
	// Get dogs
	dogs, total, err := h.service.GetAll(page, limit)
	if err != nil {
		response.InternalServerError(c, "Failed to fetch dog breeds")
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
	
	response.SuccessWithPagination(c, "Dog breeds fetched successfully", dogs, pagination)
}

// GetByID godoc
// @Summary Get dog breed by ID
// @Description Get a specific dog breed by its ID
// @Tags Public APIs
// @Accept json
// @Produce json
// @Param id path int true "Dog ID"
// @Success 200 {object} response.Response
// @Failure 404 {object} response.Response
// @Router /public/dogs/{id} [get]
func (h *DogHandler) GetByID(c *gin.Context) {
	// Parse ID
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.BadRequest(c, "Invalid dog ID")
		return
	}
	
	// Get dog
	dog, err := h.service.GetByID(id)
	if err != nil {
		response.NotFound(c, "Dog breed not found")
		return
	}
	
	response.OK(c, "Dog breed fetched successfully", dog)
}

// GetRandom godoc
// @Summary Get a random dog breed
// @Description Get a randomly selected dog breed
// @Tags Public APIs
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Router /public/dogs/random [get]
func (h *DogHandler) GetRandom(c *gin.Context) {
	dog, err := h.service.GetRandom()
	if err != nil {
		response.InternalServerError(c, "Failed to fetch random dog breed")
		return
	}
	
	response.OK(c, "Random dog breed fetched successfully", dog)
}
