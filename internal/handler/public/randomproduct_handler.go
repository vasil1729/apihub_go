package public

import (
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/ultimatum/apihub_go/internal/service/public"
	"github.com/ultimatum/apihub_go/pkg/response"
)

type RandomProductHandler struct {
	service *public.RandomProductService
}

func NewRandomProductHandler(service *public.RandomProductService) *RandomProductHandler {
	return &RandomProductHandler{service: service}
}

// @Summary Get all products
// @Tags Public APIs
// @Param page query int false "Page" default(1)
// @Param limit query int false "Limit" default(10)
// @Success 200 {object} response.PaginatedResponse
// @Router /public/randomproducts [get]
func (h *RandomProductHandler) GetAll(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}
	products, total, err := h.service.GetAll(page, limit)
	if err != nil {
		response.InternalServerError(c, "Failed to fetch products")
		return
	}
	totalPages := (total + limit - 1) / limit
	response.SuccessWithPagination(c, "Products fetched successfully", products, response.Pagination{
		Page: page, Limit: limit, TotalPages: totalPages, TotalItems: int64(total),
		HasNext: page < totalPages, HasPrev: page > 1,
	})
}

// @Summary Get product by ID
// @Tags Public APIs
// @Param id path int true "Product ID"
// @Success 200 {object} response.Response
// @Router /public/randomproducts/{id} [get]
func (h *RandomProductHandler) GetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.BadRequest(c, "Invalid product ID")
		return
	}
	product, err := h.service.GetByID(id)
	if err != nil {
		response.NotFound(c, "Product not found")
		return
	}
	response.OK(c, "Product fetched successfully", product)
}

// @Summary Get random product
// @Tags Public APIs
// @Success 200 {object} response.Response
// @Router /public/randomproducts/random [get]
func (h *RandomProductHandler) GetRandom(c *gin.Context) {
	product, err := h.service.GetRandom()
	if err != nil {
		response.InternalServerError(c, "Failed to fetch random product")
		return
	}
	response.OK(c, "Random product fetched successfully", product)
}
