package public_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/ultimatum/apihub_go/internal/service/public"
)

func TestStockService_GetAll(t *testing.T) {
	// Create service with test data
	service, err := public.NewStockService("../../data")
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
			stocks, total, err := service.GetAll(tt.page, tt.limit)

			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Greater(t, total, 0, "Total should be greater than 0")

				if tt.expectEmpty {
					assert.Empty(t, stocks)
				} else {
					assert.GreaterOrEqual(t, len(stocks), tt.minExpected)
					assert.LessOrEqual(t, len(stocks), tt.limit)
				}
			}
		})
	}
}

func TestStockService_GetBySymbol(t *testing.T) {
	service, err := public.NewStockService("../../data")
	assert.NoError(t, err)

	tests := []struct {
		name        string
		symbol      string
		expectError bool
	}{
		{
			name:        "Valid symbol - TCS",
			symbol:      "TCS",
			expectError: false,
		},
		{
			name:        "Valid symbol - lowercase",
			symbol:      "reliance",
			expectError: false,
		},
		{
			name:        "Valid symbol - INFY",
			symbol:      "INFY",
			expectError: false,
		},
		{
			name:        "Invalid symbol - not found",
			symbol:      "NOTEXIST",
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stock, err := service.GetBySymbol(tt.symbol)

			if tt.expectError {
				assert.Error(t, err)
				assert.Nil(t, stock)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, stock)
				assert.NotEmpty(t, stock.Symbol)
				assert.NotEmpty(t, stock.Name)
			}
		})
	}
}
