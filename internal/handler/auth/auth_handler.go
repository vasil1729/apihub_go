package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ultimatum/apihub_go/internal/domain/auth"
	"github.com/ultimatum/apihub_go/pkg/response"
)

type AuthHandler struct {
	service auth.AuthService
}

func NewAuthHandler(service auth.AuthService) *AuthHandler {
	return &AuthHandler{service: service}
}

// @Summary Register a new user
// @Tags Authentication
// @Accept json
// @Produce json
// @Param request body auth.RegisterRequest true "Registration Request"
// @Success 201 {object} response.Response{data=auth.AuthResponse}
// @Failure 400 {object} response.Response
// @Failure 409 {object} response.Response
// @Router /auth/register [post]
func (h *AuthHandler) Register(c *gin.Context) {
	var req auth.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	user, err := h.service.Register(c.Request.Context(), req)
	if err != nil {
		if err.Error() == "username or email already exists" {
			response.Conflict(c, err.Error())
			return
		}
		response.InternalServerError(c, "Failed to register user")
		return
	}

	// TODO: Generate JWT Token
	resp := auth.AuthResponse{
		Token: "jwt_token_placeholder",
		User:  *user,
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User registered successfully",
		"data":    resp,
	})
}
