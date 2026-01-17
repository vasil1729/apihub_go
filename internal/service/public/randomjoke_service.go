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

// RandomJokeService handles random joke operations
type RandomJokeService struct {
	jokes []public.RandomJoke
	rand  *rand.Rand
}

// NewRandomJokeService creates a new random joke service
func NewRandomJokeService(dataPath string) (*RandomJokeService, error) {
	service := &RandomJokeService{
		rand: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
	
	// Load jokes from JSON file
	if err := service.loadJokes(dataPath); err != nil {
		return nil, err
	}
	
	return service, nil
}

// loadJokes loads jokes from the JSON file
func (s *RandomJokeService) loadJokes(dataPath string) error {
	filePath := filepath.Join(dataPath, "randomjoke.json")
	
	data, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to read randomjoke.json: %w", err)
	}
	
	if err := json.Unmarshal(data, &s.jokes); err != nil {
		return fmt.Errorf("failed to parse randomjoke.json: %w", err)
	}
	
	return nil
}

// GetAll returns all jokes with pagination
func (s *RandomJokeService) GetAll(page, limit int) ([]public.RandomJoke, int, error) {
	total := len(s.jokes)
	
	// Calculate offset
	offset := (page - 1) * limit
	
	// Validate pagination
	if offset >= total {
		return []public.RandomJoke{}, total, nil
	}
	
	// Calculate end index
	end := offset + limit
	if end > total {
		end = total
	}
	
	return s.jokes[offset:end], total, nil
}

// GetByID returns a joke by ID
func (s *RandomJokeService) GetByID(id int) (*public.RandomJoke, error) {
	for _, joke := range s.jokes {
		if joke.ID == id {
			return &joke, nil
		}
	}
	return nil, fmt.Errorf("joke with ID %d not found", id)
}

// GetRandom returns a random joke
func (s *RandomJokeService) GetRandom() (*public.RandomJoke, error) {
	if len(s.jokes) == 0 {
		return nil, fmt.Errorf("no jokes available")
	}
	
	// Return a truly random joke
	index := s.rand.Intn(len(s.jokes))
	return &s.jokes[index], nil
}
