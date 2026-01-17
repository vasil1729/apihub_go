package public

import (
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/ultimatum/apihub_go/internal/service/public"
	"github.com/ultimatum/apihub_go/pkg/response"
)

type BookHandler struct {
	service *public.BookService
}

func NewBookHandler(service *public.BookService) *BookHandler {
	return &BookHandler{service: service}
}

// @Summary Get all books
// @Tags Public APIs
// @Param page query int false "Page" default(1)
// @Param limit query int false "Limit" default(10)
// @Success 200 {object} response.PaginatedResponse
// @Router /public/books [get]
func (h *BookHandler) GetAll(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}
	books, total, err := h.service.GetAll(page, limit)
	if err != nil {
		response.InternalServerError(c, "Failed to fetch books")
		return
	}
	totalPages := (total + limit - 1) / limit
	response.SuccessWithPagination(c, "Books fetched successfully", books, response.Pagination{
		Page: page, Limit: limit, TotalPages: totalPages, TotalItems: int64(total),
		HasNext: page < totalPages, HasPrev: page > 1,
	})
}

// @Summary Get book by ID
// @Tags Public APIs
// @Param id path int true "Book ID"
// @Success 200 {object} response.Response
// @Router /public/books/{id} [get]
func (h *BookHandler) GetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.BadRequest(c, "Invalid book ID")
		return
	}
	book, err := h.service.GetByID(id)
	if err != nil {
		response.NotFound(c, "Book not found")
		return
	}
	response.OK(c, "Book fetched successfully", book)
}

// @Summary Get random book
// @Tags Public APIs
// @Success 200 {object} response.Response
// @Router /public/books/random [get]
func (h *BookHandler) GetRandom(c *gin.Context) {
	book, err := h.service.GetRandom()
	if err != nil {
		response.InternalServerError(c, "Failed to fetch random book")
		return
	}
	response.OK(c, "Random book fetched successfully", book)
}
