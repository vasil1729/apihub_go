package public_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/ultimatum/apihub_go/internal/service/public"
)

func TestQuoteService_GetAll(t *testing.T) {
	// Create service with test data
	service, err := public.NewQuoteService("../../data")
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
			page:        10000,
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
			quotes, total, err := service.GetAll(tt.page, tt.limit)

			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Greater(t, total, 0, "Total should be greater than 0")

				if tt.expectEmpty {
					assert.Empty(t, quotes)
				} else {
					assert.GreaterOrEqual(t, len(quotes), tt.minExpected)
					assert.LessOrEqual(t, len(quotes), tt.limit)
				}
			}
		})
	}
}

func TestQuoteService_GetByID(t *testing.T) {
	service, err := public.NewQuoteService("../../data")
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
			quote, err := service.GetByID(tt.id)

			if tt.expectError {
				assert.Error(t, err)
				assert.Nil(t, quote)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, quote)
				assert.Equal(t, tt.id, quote.ID)
				assert.NotEmpty(t, quote.Content)
				assert.NotEmpty(t, quote.Author)
			}
		})
	}
}

func TestQuoteService_GetRandom(t *testing.T) {
	service, err := public.NewQuoteService("../../data")
	assert.NoError(t, err)

	// Test that GetRandom returns a quote
	quote, err := service.GetRandom()
	assert.NoError(t, err)
	assert.NotNil(t, quote)
	assert.NotEmpty(t, quote.Content)
	assert.NotEmpty(t, quote.Author)
	assert.Greater(t, quote.ID, 0)
	
	// Test randomness by calling multiple times
	ids := make(map[int]bool)
	for i := 0; i < 10; i++ {
		quote, err := service.GetRandom()
		assert.NoError(t, err)
		ids[quote.ID] = true
	}
	
	// With 10 calls, we should get at least 2 different quotes
	assert.GreaterOrEqual(t, len(ids), 2, "GetRandom should return different quotes")
}
