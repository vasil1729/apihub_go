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

	// Note: Register currently doesn't login automatically (no token returned in this implementation step, usually register might return token too)
	// Update: Domain model for AuthResponse expects Token.
	// But Register method signature only returns *User.
	// We should probably change Register to return token too?
	// Or just return Empty token for now as per previous implementation plan?
	// Previous implementation: Token: "jwt_token_placeholder".
	
	resp := auth.AuthResponse{
		AccessToken:  "", // Client should login after register
		RefreshToken: "",
		User:         *user,
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User registered successfully, please login",
		"data":    resp,
	})
}

// @Summary Login user
// @Tags Authentication
// @Accept json
// @Produce json
// @Param request body auth.LoginRequest true "Login Request"
// @Success 200 {object} response.Response{data=auth.AuthResponse}
// @Failure 400 {object} response.Response
// @Failure 401 {object} response.Response
// @Router /auth/login [post]
func (h *AuthHandler) Login(c *gin.Context) {
	var req auth.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	accessToken, refreshToken, user, err := h.service.Login(c.Request.Context(), req)
	if err != nil {
		if err.Error() == "invalid credentials" {
			response.Unauthorized(c, "Invalid email or password")
			return
		}
		response.InternalServerError(c, "Login failed")
		return
	}

	resp := auth.AuthResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		User:         *user,
	}

	response.OK(c, "Login successful", resp)
}

// @Summary Refresh access token
// @Tags Authentication
// @Accept json
// @Produce json
// @Param request body auth.RefreshTokenRequest true "Refresh Token Request"
// @Success 200 {object} response.Response{data=string}
// @Failure 401 {object} response.Response
// @Router /auth/refresh [post]
func (h *AuthHandler) RefreshToken(c *gin.Context) {
	var req auth.RefreshTokenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	accessToken, err := h.service.RefreshToken(c.Request.Context(), req)
	if err != nil {
		response.Unauthorized(c, err.Error())
		return
	}

	response.OK(c, "Token refreshed successfully", gin.H{"accessToken": accessToken})
}

// @Summary Logout user
// @Tags Authentication
// @Accept json
// @Produce json
// @Success 200 {object} response.Response
// @Router /auth/logout [post]
func (h *AuthHandler) Logout(c *gin.Context) {
	// In a stateless JWT setup, logout is handled client-side by deleting the token.
	// Server-side, we could blacklist the token if we had Redis, but for now just return success.
	response.OK(c, "Logged out successfully", nil)
}
