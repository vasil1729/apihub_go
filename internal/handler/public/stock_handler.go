package public

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ultimatum/apihub_go/internal/service/public"
	"github.com/ultimatum/apihub_go/pkg/response"
)

// StockHandler handles stock HTTP requests
type StockHandler struct {
	service *public.StockService
}

// NewStockHandler creates a new stock handler
func NewStockHandler(service *public.StockService) *StockHandler {
	return &StockHandler{
		service: service,
	}
}

// GetAll godoc
// @Summary Get all stocks
// @Description Get a paginated list of NSE stocks
// @Tags Public APIs
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Items per page" default(10)
// @Success 200 {object} response.PaginatedResponse
// @Router /public/stocks [get]
func (h *StockHandler) GetAll(c *gin.Context) {
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
	
	// Get stocks
	stocks, total, err := h.service.GetAll(page, limit)
	if err != nil {
		response.InternalServerError(c, "Failed to fetch stocks")
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
	
	response.SuccessWithPagination(c, "Stocks fetched successfully", stocks, pagination)
}

// GetBySymbol godoc
// @Summary Get stock by symbol
// @Description Get a specific stock by its symbol (e.g., TCS, RELIANCE)
// @Tags Public APIs
// @Accept json
// @Produce json
// @Param symbol path string true "Stock Symbol"
// @Success 200 {object} response.Response
// @Failure 404 {object} response.Response
// @Router /public/stocks/{symbol} [get]
func (h *StockHandler) GetBySymbol(c *gin.Context) {
	// Get symbol from path
	symbol := c.Param("symbol")
	if symbol == "" {
		response.BadRequest(c, "Stock symbol is required")
		return
	}
	
	// Get stock
	stock, err := h.service.GetBySymbol(symbol)
	if err != nil {
		response.NotFound(c, "Stock not found")
		return
	}
	
	response.OK(c, "Stock fetched successfully", stock)
}
