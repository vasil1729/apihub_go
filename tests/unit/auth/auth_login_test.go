package auth_test

import (
	"context"
	"testing"
	
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	authDomain "github.com/ultimatum/apihub_go/internal/domain/auth"
	authServicePkg "github.com/ultimatum/apihub_go/internal/service/auth"
	"github.com/ultimatum/apihub_go/pkg/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
	"golang.org/x/crypto/bcrypt"
)

func TestAuthService_Login(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	// Create a config with JWT Secret
	cfg := &config.Config{
		JWTSecret: "test-secret",
	}

	mt.Run("Success", func(mt *mtest.T) {
		service := authServicePkg.NewAuthService(mt.DB, cfg)

		// Create hashed password
		hashedPwd, _ := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)
		
		// Mock FindOne -> Return User
		userID := primitive.NewObjectID()
		mt.AddMockResponses(mtest.CreateCursorResponse(1, "apihub.users", mtest.FirstBatch, bson.D{
			{Key: "username", Value: "testuser"},
			{Key: "email", Value: "test@example.com"},
			{Key: "password", Value: string(hashedPwd)},
			{Key: "_id", Value: userID}, 
		}))

		req := authDomain.LoginRequest{
			Email:    "test@example.com",
			Password: "password123",
		}

		accessToken, refreshToken, user, err := service.Login(context.Background(), req)
		assert.NoError(mt, err)
		assert.NotEmpty(mt, accessToken)
		assert.NotEmpty(mt, refreshToken)
		require.NotNil(mt, user)
		assert.Equal(mt, "testuser", user.Username)
	})

	mt.Run("Invalid Password", func(mt *mtest.T) {
		service := authServicePkg.NewAuthService(mt.DB, cfg)

		// Create hashed password
		hashedPwd, _ := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)
		
		// Mock FindOne -> Return User
		mt.AddMockResponses(mtest.CreateCursorResponse(1, "apihub.users", mtest.FirstBatch, bson.D{
			{Key: "email", Value: "test@example.com"},
			{Key: "password", Value: string(hashedPwd)},
		}))

		req := authDomain.LoginRequest{
			Email:    "test@example.com",
			Password: "wrongpassword",
		}

		accessToken, refreshToken, user, err := service.Login(context.Background(), req)
		assert.Error(mt, err)
		assert.Empty(mt, accessToken)
		assert.Empty(mt, refreshToken)
		assert.Nil(mt, user)
		assert.Equal(mt, "invalid credentials", err.Error())
	})

	mt.Run("User Not Found", func(mt *mtest.T) {
		service := authServicePkg.NewAuthService(mt.DB, cfg)

		// Mock FindOne -> No Documents
		mt.AddMockResponses(mtest.CreateCursorResponse(0, "apihub.users", mtest.FirstBatch))

		req := authDomain.LoginRequest{
			Email:    "nonexistent@example.com",
			Password: "password123",
		}

		accessToken, refreshToken, user, err := service.Login(context.Background(), req)
		assert.Error(mt, err)
		assert.Empty(mt, accessToken)
		assert.Empty(mt, refreshToken)
		assert.Nil(mt, user)
		assert.Equal(mt, "invalid credentials", err.Error())
	})

	// Refresh Token Tests (placeholder as mocking JWT verification logic purely via mtest is tricky if logic inside service parses token using secret)
	// Actually, RefreshToken method parses the passed token.
	// We can generate a valid token using the same config secret in the test.
}
