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

// DogService handles dog operations
type DogService struct {
	dogs []public.Dog
	rand *rand.Rand
}

// NewDogService creates a new dog service
func NewDogService(dataPath string) (*DogService, error) {
	service := &DogService{
		rand: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
	
	// Load dogs from JSON file
	if err := service.loadDogs(dataPath); err != nil {
		return nil, err
	}
	
	return service, nil
}

// loadDogs loads dogs from the JSON file
func (s *DogService) loadDogs(dataPath string) error {
	filePath := filepath.Join(dataPath, "dogs.json")
	
	data, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to read dogs.json: %w", err)
	}
	
	if err := json.Unmarshal(data, &s.dogs); err != nil {
		return fmt.Errorf("failed to parse dogs.json: %w", err)
	}
	
	return nil
}

// GetAll returns all dogs with pagination
func (s *DogService) GetAll(page, limit int) ([]public.Dog, int, error) {
	total := len(s.dogs)
	
	// Calculate offset
	offset := (page - 1) * limit
	
	// Validate pagination
	if offset >= total {
		return []public.Dog{}, total, nil
	}
	
	// Calculate end index
	end := offset + limit
	if end > total {
		end = total
	}
	
	return s.dogs[offset:end], total, nil
}

// GetByID returns a dog by ID
func (s *DogService) GetByID(id int) (*public.Dog, error) {
	for _, dog := range s.dogs {
		if dog.ID == id {
			return &dog, nil
		}
	}
	return nil, fmt.Errorf("dog with ID %d not found", id)
}

// GetRandom returns a random dog
func (s *DogService) GetRandom() (*public.Dog, error) {
	if len(s.dogs) == 0 {
		return nil, fmt.Errorf("no dogs available")
	}
	
	// Return a truly random dog
	index := s.rand.Intn(len(s.dogs))
	return &s.dogs[index], nil
}
