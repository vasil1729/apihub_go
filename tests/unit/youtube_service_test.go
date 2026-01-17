package public_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/ultimatum/apihub_go/internal/service/public"
)

func TestYouTubeService_GetAll(t *testing.T) {
	service, err := public.NewYouTubeService("../../data")
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
			videos, total, err := service.GetAll(tt.page, tt.limit)
			assert.NoError(t, err)
			assert.Greater(t, total, 0, "Total should be greater than 0")

			if tt.expectEmpty {
				assert.Empty(t, videos)
			} else {
				assert.GreaterOrEqual(t, len(videos), tt.minExpected)
				assert.LessOrEqual(t, len(videos), tt.limit)
				
				// Verify video structure
				if len(videos) > 0 {
					assert.NotEmpty(t, videos[0].Title)
					assert.Greater(t, videos[0].ID, 0)
				}
			}
		})
	}
}

func TestYouTubeService_GetByID(t *testing.T) {
	service, err := public.NewYouTubeService("../../data")
	assert.NoError(t, err)

	tests := []struct {
		name        string
		id          int
		expectError bool
	}{
		{"Valid ID - 1", 1, false},
		{"Valid ID - 5", 5, false},
		{"Invalid ID - not found", 99999, true},
		{"Invalid ID - 0", 0, true},
		{"Invalid ID - negative", -1, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			video, err := service.GetByID(tt.id)

			if tt.expectError {
				assert.Error(t, err)
				assert.Nil(t, video)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, video)
				assert.Equal(t, tt.id, video.ID)
				assert.NotEmpty(t, video.Title)
			}
		})
	}
}

func TestYouTubeService_GetRandom(t *testing.T) {
	service, err := public.NewYouTubeService("../../data")
	assert.NoError(t, err)

	// Test that GetRandom returns a video
	video, err := service.GetRandom()
	assert.NoError(t, err)
	assert.NotNil(t, video)
	assert.NotEmpty(t, video.Title)
	assert.Greater(t, video.ID, 0)
	
	// Test randomness by calling multiple times
	ids := make(map[int]bool)
	for i := 0; i < 20; i++ {
		video, err := service.GetRandom()
		assert.NoError(t, err)
		ids[video.ID] = true
	}
	
	// With 20 calls, we should get at least 3 different videos
	assert.GreaterOrEqual(t, len(ids), 3, "GetRandom should return different videos")
}
