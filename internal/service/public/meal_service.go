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

// MealService handles meal operations
type MealService struct {
	meals []public.Meal
	rand  *rand.Rand
}

// NewMealService creates a new meal service
func NewMealService(dataPath string) (*MealService, error) {
	service := &MealService{
		rand: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
	
	// Load meals from JSON file
	if err := service.loadMeals(dataPath); err != nil {
		return nil, err
	}
	
	return service, nil
}

// loadMeals loads meals from the JSON file
func (s *MealService) loadMeals(dataPath string) error {
	filePath := filepath.Join(dataPath, "meals.json")
	
	data, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to read meals.json: %w", err)
	}
	
	if err := json.Unmarshal(data, &s.meals); err != nil {
		return fmt.Errorf("failed to parse meals.json: %w", err)
	}
	
	return nil
}

// GetAll returns all meals with pagination
func (s *MealService) GetAll(page, limit int) ([]public.Meal, int, error) {
	total := len(s.meals)
	
	// Calculate offset
	offset := (page - 1) * limit
	
	// Validate pagination
	if offset >= total {
		return []public.Meal{}, total, nil
	}
	
	// Calculate end index
	end := offset + limit
	if end > total {
		end = total
	}
	
	return s.meals[offset:end], total, nil
}

// GetByID returns a meal by ID
func (s *MealService) GetByID(id int) (*public.Meal, error) {
	for _, meal := range s.meals {
		if meal.ID == id {
			return &meal, nil
		}
	}
	return nil, fmt.Errorf("meal with ID %d not found", id)
}

// GetRandom returns a random meal
func (s *MealService) GetRandom() (*public.Meal, error) {
	if len(s.meals) == 0 {
		return nil, fmt.Errorf("no meals available")
	}
	
	// Return a truly random meal
	index := s.rand.Intn(len(s.meals))
	return &s.meals[index], nil
}
