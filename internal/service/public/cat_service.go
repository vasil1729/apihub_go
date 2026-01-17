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

type CatService struct {
	cats []public.Cat
	rand *rand.Rand
}

func NewCatService(dataPath string) (*CatService, error) {
	service := &CatService{
		rand: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
	
	if err := service.loadCats(dataPath); err != nil {
		return nil, err
	}
	
	return service, nil
}

func (s *CatService) loadCats(dataPath string) error {
	filePath := filepath.Join(dataPath, "cats.json")
	
	data, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to read cats.json: %w", err)
	}
	
	if err := json.Unmarshal(data, &s.cats); err != nil {
		return fmt.Errorf("failed to parse cats.json: %w", err)
	}
	
	return nil
}

func (s *CatService) GetAll(page, limit int) ([]public.Cat, int, error) {
	total := len(s.cats)
	offset := (page - 1) * limit
	
	if offset >= total {
		return []public.Cat{}, total, nil
	}
	
	end := offset + limit
	if end > total {
		end = total
	}
	
	return s.cats[offset:end], total, nil
}

func (s *CatService) GetByID(id int) (*public.Cat, error) {
	for _, cat := range s.cats {
		if cat.ID == id {
			return &cat, nil
		}
	}
	return nil, fmt.Errorf("cat with ID %d not found", id)
}

func (s *CatService) GetRandom() (*public.Cat, error) {
	if len(s.cats) == 0 {
		return nil, fmt.Errorf("no cats available")
	}
	
	index := s.rand.Intn(len(s.cats))
	return &s.cats[index], nil
}
