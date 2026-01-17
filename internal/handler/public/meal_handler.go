package public

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ultimatum/apihub_go/internal/service/public"
	"github.com/ultimatum/apihub_go/pkg/response"
)

// MealHandler handles meal HTTP requests
type MealHandler struct {
	service *public.MealService
}

// NewMealHandler creates a new meal handler
func NewMealHandler(service *public.MealService) *MealHandler {
	return &MealHandler{
		service: service,
	}
}

// GetAll godoc
// @Summary Get all meals
// @Description Get a paginated list of meals
// @Tags Public APIs
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Items per page" default(10)
// @Success 200 {object} response.PaginatedResponse
// @Router /public/meals [get]
func (h *MealHandler) GetAll(c *gin.Context) {
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
	
	// Get meals
	meals, total, err := h.service.GetAll(page, limit)
	if err != nil {
		response.InternalServerError(c, "Failed to fetch meals")
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
	
	response.SuccessWithPagination(c, "Meals fetched successfully", meals, pagination)
}

// GetByID godoc
// @Summary Get meal by ID
// @Description Get a specific meal by its ID
// @Tags Public APIs
// @Accept json
// @Produce json
// @Param id path int true "Meal ID"
// @Success 200 {object} response.Response
// @Failure 404 {object} response.Response
// @Router /public/meals/{id} [get]
func (h *MealHandler) GetByID(c *gin.Context) {
	// Parse ID
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.BadRequest(c, "Invalid meal ID")
		return
	}
	
	// Get meal
	meal, err := h.service.GetByID(id)
	if err != nil {
		response.NotFound(c, "Meal not found")
		return
	}
	
	response.OK(c, "Meal fetched successfully", meal)
}

// GetRandom godoc
// @Summary Get a random meal
// @Description Get a randomly selected meal
// @Tags Public APIs
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Router /public/meals/random [get]
func (h *MealHandler) GetRandom(c *gin.Context) {
	meal, err := h.service.GetRandom()
	if err != nil {
		response.InternalServerError(c, "Failed to fetch random meal")
		return
	}
	
	response.OK(c, "Random meal fetched successfully", meal)
}
