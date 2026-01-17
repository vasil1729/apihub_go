package public

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ultimatum/apihub_go/internal/service/public"
	"github.com/ultimatum/apihub_go/pkg/response"
)

// QuoteHandler handles quote HTTP requests
type QuoteHandler struct {
	service *public.QuoteService
}

// NewQuoteHandler creates a new quote handler
func NewQuoteHandler(service *public.QuoteService) *QuoteHandler {
	return &QuoteHandler{
		service: service,
	}
}

// GetAll godoc
// @Summary Get all quotes
// @Description Get a paginated list of quotes
// @Tags Public APIs
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Items per page" default(10)
// @Success 200 {object} response.PaginatedResponse
// @Router /public/quotes [get]
func (h *QuoteHandler) GetAll(c *gin.Context) {
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
	
	// Get quotes
	quotes, total, err := h.service.GetAll(page, limit)
	if err != nil {
		response.InternalServerError(c, "Failed to fetch quotes")
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
	
	response.SuccessWithPagination(c, "Quotes fetched successfully", quotes, pagination)
}

// GetByID godoc
// @Summary Get quote by ID
// @Description Get a specific quote by its ID
// @Tags Public APIs
// @Accept json
// @Produce json
// @Param id path int true "Quote ID"
// @Success 200 {object} response.Response
// @Failure 404 {object} response.Response
// @Router /public/quotes/{id} [get]
func (h *QuoteHandler) GetByID(c *gin.Context) {
	// Parse ID
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.BadRequest(c, "Invalid quote ID")
		return
	}
	
	// Get quote
	quote, err := h.service.GetByID(id)
	if err != nil {
		response.NotFound(c, "Quote not found")
		return
	}
	
	response.OK(c, "Quote fetched successfully", quote)
}

// GetRandom godoc
// @Summary Get a random quote
// @Description Get a randomly selected quote
// @Tags Public APIs
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Router /public/quotes/random [get]
func (h *QuoteHandler) GetRandom(c *gin.Context) {
	quote, err := h.service.GetRandom()
	if err != nil {
		response.InternalServerError(c, "Failed to fetch random quote")
		return
	}
	
	response.OK(c, "Random quote fetched successfully", quote)
}
