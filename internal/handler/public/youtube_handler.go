package public

import (
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/ultimatum/apihub_go/internal/service/public"
	"github.com/ultimatum/apihub_go/pkg/response"
)

type YouTubeHandler struct {
	service *public.YouTubeService
}

func NewYouTubeHandler(service *public.YouTubeService) *YouTubeHandler {
	return &YouTubeHandler{service: service}
}

// @Summary Get all YouTube videos
// @Tags Public APIs
// @Param page query int false "Page" default(1)
// @Param limit query int false "Limit" default(10)
// @Success 200 {object} response.PaginatedResponse
// @Router /public/youtube [get]
func (h *YouTubeHandler) GetAll(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}
	videos, total, err := h.service.GetAll(page, limit)
	if err != nil {
		response.InternalServerError(c, "Failed to fetch videos")
		return
	}
	totalPages := (total + limit - 1) / limit
	response.SuccessWithPagination(c, "Videos fetched successfully", videos, response.Pagination{
		Page: page, Limit: limit, TotalPages: totalPages, TotalItems: int64(total),
		HasNext: page < totalPages, HasPrev: page > 1,
	})
}

// @Summary Get video by ID
// @Tags Public APIs
// @Param id path int true "Video ID"
// @Success 200 {object} response.Response
// @Router /public/youtube/{id} [get]
func (h *YouTubeHandler) GetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.BadRequest(c, "Invalid video ID")
		return
	}
	video, err := h.service.GetByID(id)
	if err != nil {
		response.NotFound(c, "Video not found")
		return
	}
	response.OK(c, "Video fetched successfully", video)
}

// @Summary Get random video
// @Tags Public APIs
// @Success 200 {object} response.Response
// @Router /public/youtube/random [get]
func (h *YouTubeHandler) GetRandom(c *gin.Context) {
	video, err := h.service.GetRandom()
	if err != nil {
		response.InternalServerError(c, "Failed to fetch random video")
		return
	}
	response.OK(c, "Random video fetched successfully", video)
}
