package integration_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/ultimatum/apihub_go/internal/domain/auth"
	authHandlerPkg "github.com/ultimatum/apihub_go/internal/handler/auth"
)

func setupRefreshRouter(service auth.AuthService) *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	v1 := router.Group("/api/v1")
	authGroup := v1.Group("/auth")
	
	handler := authHandlerPkg.NewAuthHandler(service)
	authHandlerPkg.SetupAuthRoutes(authGroup, handler)
	return router
}

func TestAuthAPI_RefreshToken_Integration(t *testing.T) {
	mockService := new(MockAuthService)
	router := setupRefreshRouter(mockService)
	
	t.Run("Refresh Token Success", func(t *testing.T) {
		refreshReq := auth.RefreshTokenRequest{
			RefreshToken: "valid_refresh_token",
		}

		mockService.On("RefreshToken", mock.Anything, refreshReq).Return("new_access_token", nil).Once()

		w := httptest.NewRecorder()
		body, _ := json.Marshal(refreshReq)
		req, _ := http.NewRequest("POST", "/api/v1/auth/refresh", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		
		var resp struct {
			Data map[string]string `json:"data"`
		}
		err := json.Unmarshal(w.Body.Bytes(), &resp)
		assert.NoError(t, err)
		assert.Equal(t, "new_access_token", resp.Data["accessToken"])

		mockService.AssertExpectations(t)
	})
}
