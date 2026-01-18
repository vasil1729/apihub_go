package integration_test

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/ultimatum/apihub_go/internal/domain/auth"
	authHandlerPkg "github.com/ultimatum/apihub_go/internal/handler/auth"
)

// MockAuthService
type MockAuthService struct {
	mock.Mock
}

func (m *MockAuthService) Register(ctx context.Context, req auth.RegisterRequest) (*auth.User, error) {
	args := m.Called(ctx, req)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*auth.User), args.Error(1)
}

func (m *MockAuthService) Login(ctx context.Context, req auth.LoginRequest) (string, string, *auth.User, error) {
	args := m.Called(ctx, req)
	if args.Get(0) == "" {
		return "", "", nil, args.Error(3)
	}
	return args.String(0), args.String(1), args.Get(2).(*auth.User), args.Error(3)
}

func (m *MockAuthService) RefreshToken(ctx context.Context, req auth.RefreshTokenRequest) (string, error) {
	args := m.Called(ctx, req)
	return args.String(0), args.Error(1)
}

func setupAuthRouter(service auth.AuthService) *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	v1 := router.Group("/api/v1")
	authGroup := v1.Group("/auth")
	
	handler := authHandlerPkg.NewAuthHandler(service)
	authHandlerPkg.SetupAuthRoutes(authGroup, handler)
	return router
}

func TestAuthAPI_Integration(t *testing.T) {
	mockService := new(MockAuthService)
	router := setupAuthRouter(mockService)
	
	testUserReq := auth.RegisterRequest{
		Username: "testuser",
		Email:    "test@example.com",
		Password: "password123",
	}

	testUserRes := &auth.User{
		Username: "testuser",
		Email:    "test@example.com",
	}

	t.Run("Register Success", func(t *testing.T) {
		mockService.On("Register", mock.Anything, testUserReq).Return(testUserRes, nil).Once()

		w := httptest.NewRecorder()
		body, _ := json.Marshal(testUserReq)
		req, _ := http.NewRequest("POST", "/api/v1/auth/register", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusCreated, w.Code)
		mockService.AssertExpectations(t)
	})

	t.Run("Register Duplicate Failure", func(t *testing.T) {
		mockService.On("Register", mock.Anything, testUserReq).Return(nil, errors.New("username or email already exists")).Once()
		
		w := httptest.NewRecorder()
		body, _ := json.Marshal(testUserReq)
		req, _ := http.NewRequest("POST", "/api/v1/auth/register", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusConflict, w.Code)
		mockService.AssertExpectations(t)
	})

	t.Run("Login Success", func(t *testing.T) {
		loginReq := auth.LoginRequest{
			Email:    "test@example.com",
			Password: "password123",
		}
		
		mockService.On("Login", mock.Anything, loginReq).Return("access_token", "refresh_token", testUserRes, nil).Once()

		w := httptest.NewRecorder()
		body, _ := json.Marshal(loginReq)
		req, _ := http.NewRequest("POST", "/api/v1/auth/login", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		
		var resp struct {
			Data auth.AuthResponse `json:"data"`
		}
		err := json.Unmarshal(w.Body.Bytes(), &resp)
		assert.NoError(t, err)
		assert.Equal(t, "access_token", resp.Data.AccessToken)
		assert.Equal(t, "refresh_token", resp.Data.RefreshToken)
		
		mockService.AssertExpectations(t)
	})

	t.Run("Login Invalid Credentials", func(t *testing.T) {
		loginReq := auth.LoginRequest{
			Email:    "test@example.com",
			Password: "wrong",
		}
		
		mockService.On("Login", mock.Anything, loginReq).Return("", "", nil, errors.New("invalid credentials")).Once()

		w := httptest.NewRecorder()
		body, _ := json.Marshal(loginReq)
		req, _ := http.NewRequest("POST", "/api/v1/auth/login", bytes.NewBuffer(body))
		req.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
		mockService.AssertExpectations(t)
	})

	t.Run("Logout Success", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/api/v1/auth/logout", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Contains(t, w.Body.String(), "Logged out successfully")
	})
}
