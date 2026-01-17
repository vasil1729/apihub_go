package public_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/ultimatum/apihub_go/internal/service/public"
)

func TestRandomUserService_GetAll(t *testing.T) {
	// Create service with test data
	service, err := public.NewRandomUserService("../../data")
	assert.NoError(t, err)
	assert.NotNil(t, service)

	tests := []struct {
		name          string
		page          int
		limit         int
		expectError   bool
		expectEmpty   bool
		minExpected   int
	}{
		{
			name:        "Valid pagination - page 1, limit 10",
			page:        1,
			limit:       10,
			expectError: false,
			minExpected: 1,
		},
		{
			name:        "Valid pagination - page 2, limit 5",
			page:        2,
			limit:       5,
			expectError: false,
			minExpected: 1,
		},
		{
			name:        "Page beyond data",
			page:        1000,
			limit:       10,
			expectError: false,
			expectEmpty: true,
		},
		{
			name:        "First page with limit 1",
			page:        1,
			limit:       1,
			expectError: false,
			minExpected: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			users, total, err := service.GetAll(tt.page, tt.limit)

			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Greater(t, total, 0, "Total should be greater than 0")

				if tt.expectEmpty {
					assert.Empty(t, users)
				} else {
					assert.GreaterOrEqual(t, len(users), tt.minExpected)
					assert.LessOrEqual(t, len(users), tt.limit)
				}
			}
		})
	}
}

func TestRandomUserService_GetByID(t *testing.T) {
	service, err := public.NewRandomUserService("../../data")
	assert.NoError(t, err)

	tests := []struct {
		name        string
		id          int
		expectError bool
	}{
		{
			name:        "Valid ID - 1",
			id:          1,
			expectError: false,
		},
		{
			name:        "Valid ID - 5",
			id:          5,
			expectError: false,
		},
		{
			name:        "Invalid ID - not found",
			id:          99999,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user, err := service.GetByID(tt.id)

			if tt.expectError {
				assert.Error(t, err)
				assert.Nil(t, user)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, user)
				assert.Equal(t, tt.id, user.ID)
				assert.NotEmpty(t, user.Name.First)
				assert.NotEmpty(t, user.Email)
			}
		})
	}
}

func TestRandomUserService_GetRandom(t *testing.T) {
	service, err := public.NewRandomUserService("../../data")
	assert.NoError(t, err)

	user, err := service.GetRandom()
	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.Name.First)
	assert.NotEmpty(t, user.Email)
}
