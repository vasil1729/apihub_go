package public

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"time"
	"github.com/ultimatum/apihub_go/internal/domain/public"
)

type YouTubeService struct {
	videos []public.YouTubeVideo
	rand   *rand.Rand
}

func NewYouTubeService(dataPath string) (*YouTubeService, error) {
	service := &YouTubeService{
		rand: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
	if err := service.loadVideos(dataPath); err != nil {
		return nil, err
	}
	return service, nil
}

func (s *YouTubeService) loadVideos(dataPath string) error {
	data, err := os.ReadFile(filepath.Join(dataPath, "youtube", "videos.json"))
	if err != nil {
		return fmt.Errorf("failed to read videos.json: %w", err)
	}
	if err := json.Unmarshal(data, &s.videos); err != nil {
		return fmt.Errorf("failed to parse videos.json: %w", err)
	}
	return nil
}

func (s *YouTubeService) GetAll(page, limit int) ([]public.YouTubeVideo, int, error) {
	total := len(s.videos)
	offset := (page - 1) * limit
	if offset >= total {
		return []public.YouTubeVideo{}, total, nil
	}
	end := offset + limit
	if end > total {
		end = total
	}
	return s.videos[offset:end], total, nil
}

func (s *YouTubeService) GetByID(id int) (*public.YouTubeVideo, error) {
	for _, video := range s.videos {
		if video.ID == id {
			return &video, nil
		}
	}
	return nil, fmt.Errorf("video with ID %d not found", id)
}

func (s *YouTubeService) GetRandom() (*public.YouTubeVideo, error) {
	if len(s.videos) == 0 {
		return nil, fmt.Errorf("no videos available")
	}
	return &s.videos[s.rand.Intn(len(s.videos))], nil
}
