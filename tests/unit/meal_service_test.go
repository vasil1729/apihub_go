package public_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/ultimatum/apihub_go/internal/service/public"
)

func TestMealService_GetAll(t *testing.T) {
	// Create service with test data
	service, err := public.NewMealService("../../data")
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
			page:        100000,
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
			meals, total, err := service.GetAll(tt.page, tt.limit)

			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Greater(t, total, 0, "Total should be greater than 0")

				if tt.expectEmpty {
					assert.Empty(t, meals)
				} else {
					assert.GreaterOrEqual(t, len(meals), tt.minExpected)
					assert.LessOrEqual(t, len(meals), tt.limit)
				}
			}
		})
	}
}

func TestMealService_GetByID(t *testing.T) {
	service, err := public.NewMealService("../../data")
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
			name:        "Valid ID - 50",
			id:          50,
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
			meal, err := service.GetByID(tt.id)

			if tt.expectError {
				assert.Error(t, err)
				assert.Nil(t, meal)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, meal)
				assert.Equal(t, tt.id, meal.ID)
				assert.NotEmpty(t, meal.StrMeal)
			}
		})
	}
}

func TestMealService_GetRandom(t *testing.T) {
	service, err := public.NewMealService("../../data")
	assert.NoError(t, err)

	// Test that GetRandom returns a meal
	meal, err := service.GetRandom()
	assert.NoError(t, err)
	assert.NotNil(t, meal)
	assert.NotEmpty(t, meal.StrMeal)
	assert.Greater(t, meal.ID, 0)
	
	// Test randomness by calling multiple times
	ids := make(map[int]bool)
	for i := 0; i < 10; i++ {
		meal, err := service.GetRandom()
		assert.NoError(t, err)
		ids[meal.ID] = true
	}
	
	// With 10 calls, we should get at least 2 different meals
	assert.GreaterOrEqual(t, len(ids), 2, "GetRandom should return different meals")
}
