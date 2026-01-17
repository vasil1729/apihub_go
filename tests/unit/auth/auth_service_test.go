package auth_test

import (
	"context"
	"testing"
	
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	authDomain "github.com/ultimatum/apihub_go/internal/domain/auth"
	authServicePkg "github.com/ultimatum/apihub_go/internal/service/auth"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestAuthService_Register(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("Success", func(mt *mtest.T) {
		service := authServicePkg.NewAuthService(mt.DB)
		
		// Mock FindOne -> No Documents (user doesn't exist) -> ErrNoDocuments
		mt.AddMockResponses(mtest.CreateCursorResponse(0, "apihub.users", mtest.FirstBatch))
		
		// Mock InsertOne -> Success
		mt.AddMockResponses(mtest.CreateSuccessResponse())

		req := authDomain.RegisterRequest{
			Username: "testuser",
			Email:    "test@example.com",
			Password: "password123",
		}

		user, err := service.Register(context.Background(), req)
		assert.NoError(mt, err)
		require.NotNil(mt, user)
		assert.Equal(mt, "testuser", user.Username)
		assert.NotEmpty(mt, user.Password)
	})

	mt.Run("User Already Exists", func(mt *mtest.T) {
		service := authServicePkg.NewAuthService(mt.DB)

		// Mock FindOne -> Success (user exists)
		mt.AddMockResponses(mtest.CreateCursorResponse(1, "apihub.users", mtest.FirstBatch, bson.D{
			{Key: "username", Value: "testuser"},
			{Key: "email", Value: "test@example.com"},
		}))

		req := authDomain.RegisterRequest{
			Username: "testuser",
			Email:    "test@example.com",
			Password: "password123",
		}

		user, err := service.Register(context.Background(), req)
		assert.Error(mt, err)
		assert.Nil(mt, user)
		assert.Contains(mt, err.Error(), "username or email already exists")
	})
}
