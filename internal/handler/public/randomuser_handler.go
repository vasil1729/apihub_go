package public

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ultimatum/apihub_go/internal/service/public"
	"github.com/ultimatum/apihub_go/pkg/response"
)

// RandomUserHandler handles random user HTTP requests
type RandomUserHandler struct {
	service *public.RandomUserService
}

// NewRandomUserHandler creates a new random user handler
func NewRandomUserHandler(service *public.RandomUserService) *RandomUserHandler {
	return &RandomUserHandler{
		service: service,
	}
}

// GetAll godoc
// @Summary Get all random users
// @Description Get a paginated list of random users
// @Tags Public APIs
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Items per page" default(10)
// @Success 200 {object} response.PaginatedResponse
// @Router /public/randomusers [get]
func (h *RandomUserHandler) GetAll(c *gin.Context) {
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
	
	// Get users
	users, total, err := h.service.GetAll(page, limit)
	if err != nil {
		response.InternalServerError(c, "Failed to fetch users")
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
	
	response.SuccessWithPagination(c, "Random users fetched successfully", users, pagination)
}

// GetByID godoc
// @Summary Get random user by ID
// @Description Get a specific random user by their ID
// @Tags Public APIs
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} response.Response
// @Failure 404 {object} response.Response
// @Router /public/randomusers/{id} [get]
func (h *RandomUserHandler) GetByID(c *gin.Context) {
	// Parse ID
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.BadRequest(c, "Invalid user ID")
		return
	}
	
	// Get user
	user, err := h.service.GetByID(id)
	if err != nil {
		response.NotFound(c, "User not found")
		return
	}
	
	response.OK(c, "User fetched successfully", user)
}

// GetRandom godoc
// @Summary Get a random user
// @Description Get a randomly selected user
// @Tags Public APIs
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Router /public/randomusers/random [get]
func (h *RandomUserHandler) GetRandom(c *gin.Context) {
	user, err := h.service.GetRandom()
	if err != nil {
		response.InternalServerError(c, "Failed to fetch random user")
		return
	}
	
	response.OK(c, "Random user fetched successfully", user)
}
