package public_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/ultimatum/apihub_go/internal/service/public"
)

func TestDogService_GetAll(t *testing.T) {
	service, err := public.NewDogService("../../data")
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
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dogs, total, err := service.GetAll(tt.page, tt.limit)
			assert.NoError(t, err)
			assert.Greater(t, total, 0)
			if tt.expectEmpty {
				assert.Empty(t, dogs)
			} else {
				assert.GreaterOrEqual(t, len(dogs), tt.minExpected)
				assert.LessOrEqual(t, len(dogs), tt.limit)
			}
		})
	}
}

func TestDogService_GetByID(t *testing.T) {
	service, err := public.NewDogService("../../data")
	assert.NoError(t, err)

	tests := []struct {
		name        string
		id          int
		expectError bool
	}{
		{"Valid ID - 1", 1, false},
		{"Valid ID - 50", 50, false},
		{"Invalid ID - not found", 99999, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dog, err := service.GetByID(tt.id)
			if tt.expectError {
				assert.Error(t, err)
				assert.Nil(t, dog)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, dog)
				assert.Equal(t, tt.id, dog.ID)
				assert.NotEmpty(t, dog.Name)
			}
		})
	}
}

func TestDogService_GetRandom(t *testing.T) {
	service, err := public.NewDogService("../../data")
	assert.NoError(t, err)

	dog, err := service.GetRandom()
	assert.NoError(t, err)
	assert.NotNil(t, dog)
	assert.NotEmpty(t, dog.Name)
	assert.Greater(t, dog.ID, 0)
	
	ids := make(map[int]bool)
	for i := 0; i < 10; i++ {
		dog, err := service.GetRandom()
		assert.NoError(t, err)
		ids[dog.ID] = true
	}
	assert.GreaterOrEqual(t, len(ids), 2, "GetRandom should return different dogs")
}
