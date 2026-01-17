package public

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	
	"github.com/ultimatum/apihub_go/internal/domain/public"
)

// RandomUserService handles random user operations
type RandomUserService struct {
	users []public.RandomUser
}

// NewRandomUserService creates a new random user service
func NewRandomUserService(dataPath string) (*RandomUserService, error) {
	service := &RandomUserService{}
	
	// Load users from JSON file
	if err := service.loadUsers(dataPath); err != nil {
		return nil, err
	}
	
	return service, nil
}

// loadUsers loads users from the JSON file
func (s *RandomUserService) loadUsers(dataPath string) error {
	filePath := filepath.Join(dataPath, "randomuser.json")
	
	data, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to read randomuser.json: %w", err)
	}
	
	if err := json.Unmarshal(data, &s.users); err != nil {
		return fmt.Errorf("failed to parse randomuser.json: %w", err)
	}
	
	return nil
}

// GetAll returns all users with pagination
func (s *RandomUserService) GetAll(page, limit int) ([]public.RandomUser, int, error) {
	total := len(s.users)
	
	// Calculate offset
	offset := (page - 1) * limit
	
	// Validate pagination
	if offset >= total {
		return []public.RandomUser{}, total, nil
	}
	
	// Calculate end index
	end := offset + limit
	if end > total {
		end = total
	}
	
	return s.users[offset:end], total, nil
}

// GetByID returns a user by ID
func (s *RandomUserService) GetByID(id int) (*public.RandomUser, error) {
	for _, user := range s.users {
		if user.ID == id {
			return &user, nil
		}
	}
	return nil, fmt.Errorf("user with ID %d not found", id)
}

// GetRandom returns a random user
func (s *RandomUserService) GetRandom() (*public.RandomUser, error) {
	if len(s.users) == 0 {
		return nil, fmt.Errorf("no users available")
	}
	
	// For simplicity, return the first user
	// In production, you'd use a proper random selection
	return &s.users[0], nil
}
