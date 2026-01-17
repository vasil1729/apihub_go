package public

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ultimatum/apihub_go/internal/service/public"
	"github.com/ultimatum/apihub_go/pkg/response"
)

type CatHandler struct {
	service *public.CatService
}

func NewCatHandler(service *public.CatService) *CatHandler {
	return &CatHandler{service: service}
}

// GetAll godoc
// @Summary Get all cat breeds
// @Description Get a paginated list of cat breeds
// @Tags Public APIs
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Items per page" default(10)
// @Success 200 {object} response.PaginatedResponse
// @Router /public/cats [get]
func (h *CatHandler) GetAll(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}
	
	cats, total, err := h.service.GetAll(page, limit)
	if err != nil {
		response.InternalServerError(c, "Failed to fetch cat breeds")
		return
	}
	
	totalPages := (total + limit - 1) / limit
	pagination := response.Pagination{
		Page:       page,
		Limit:      limit,
		TotalPages: totalPages,
		TotalItems: int64(total),
		HasNext:    page < totalPages,
		HasPrev:    page > 1,
	}
	
	response.SuccessWithPagination(c, "Cat breeds fetched successfully", cats, pagination)
}

// GetByID godoc
// @Summary Get cat breed by ID
// @Description Get a specific cat breed by its ID
// @Tags Public APIs
// @Param id path int true "Cat ID"
// @Success 200 {object} response.Response
// @Failure 404 {object} response.Response
// @Router /public/cats/{id} [get]
func (h *CatHandler) GetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.BadRequest(c, "Invalid cat ID")
		return
	}
	
	cat, err := h.service.GetByID(id)
	if err != nil {
		response.NotFound(c, "Cat breed not found")
		return
	}
	
	response.OK(c, "Cat breed fetched successfully", cat)
}

// GetRandom godoc
// @Summary Get a random cat breed
// @Description Get a randomly selected cat breed
// @Tags Public APIs
// @Success 200 {object} response.Response
// @Router /public/cats/random [get]
func (h *CatHandler) GetRandom(c *gin.Context) {
	cat, err := h.service.GetRandom()
	if err != nil {
		response.InternalServerError(c, "Failed to fetch random cat breed")
		return
	}
	
	response.OK(c, "Random cat breed fetched successfully", cat)
}
