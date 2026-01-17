package public_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/ultimatum/apihub_go/internal/service/public"
)

func TestRandomProductService_GetAll(t *testing.T) {
	service, err := public.NewRandomProductService("../../data")
	assert.NoError(t, err)
	assert.NotNil(t, service)

	tests := []struct {
		name        string
		page        int
		limit       int
		expectEmpty bool
		minExpected int
	}{
		{"Valid pagination - page 1, limit 10", 1, 10, false, 1},
		{"Valid pagination - page 2, limit 5", 2, 5, false, 1},
		{"Page beyond data", 100000, 10, true, 0},
		{"First page with limit 1", 1, 1, false, 1},
		{"Large limit", 1, 50, false, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			products, total, err := service.GetAll(tt.page, tt.limit)
			assert.NoError(t, err)
			assert.Greater(t, total, 0, "Total should be greater than 0")

			if tt.expectEmpty {
				assert.Empty(t, products)
			} else {
				assert.GreaterOrEqual(t, len(products), tt.minExpected)
				assert.LessOrEqual(t, len(products), tt.limit)
				
				// Verify product structure
				if len(products) > 0 {
					assert.NotEmpty(t, products[0].Title)
					assert.Greater(t, products[0].Price, 0.0)
				}
			}
		})
	}
}

func TestRandomProductService_GetByID(t *testing.T) {
	service, err := public.NewRandomProductService("../../data")
	assert.NoError(t, err)

	tests := []struct {
		name        string
		id          int
		expectError bool
	}{
		{"Valid ID - 1", 1, false},
		{"Valid ID - 10", 10, false},
		{"Invalid ID - not found", 99999, true},
		{"Invalid ID - 0", 0, true},
		{"Invalid ID - negative", -1, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			product, err := service.GetByID(tt.id)

			if tt.expectError {
				assert.Error(t, err)
				assert.Nil(t, product)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, product)
				assert.Equal(t, tt.id, product.ID)
				assert.NotEmpty(t, product.Title)
				assert.Greater(t, product.Price, 0.0)
				assert.NotEmpty(t, product.Category)
			}
		})
	}
}

func TestRandomProductService_GetRandom(t *testing.T) {
	service, err := public.NewRandomProductService("../../data")
	assert.NoError(t, err)

	// Test that GetRandom returns a product
	product, err := service.GetRandom()
	assert.NoError(t, err)
	assert.NotNil(t, product)
	assert.NotEmpty(t, product.Title)
	assert.Greater(t, product.ID, 0)
	assert.Greater(t, product.Price, 0.0)
	
	// Test randomness by calling multiple times
	ids := make(map[int]bool)
	for i := 0; i < 20; i++ {
		product, err := service.GetRandom()
		assert.NoError(t, err)
		ids[product.ID] = true
	}
	
	// With 20 calls, we should get at least 3 different products
	assert.GreaterOrEqual(t, len(ids), 3, "GetRandom should return different products")
}
