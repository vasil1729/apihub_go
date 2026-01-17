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

type RandomProductService struct {
	products []public.RandomProduct
	rand     *rand.Rand
}

func NewRandomProductService(dataPath string) (*RandomProductService, error) {
	service := &RandomProductService{
		rand: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
	if err := service.loadProducts(dataPath); err != nil {
		return nil, err
	}
	return service, nil
}

func (s *RandomProductService) loadProducts(dataPath string) error {
	data, err := os.ReadFile(filepath.Join(dataPath, "randomproduct.json"))
	if err != nil {
		return fmt.Errorf("failed to read randomproduct.json: %w", err)
	}
	if err := json.Unmarshal(data, &s.products); err != nil {
		return fmt.Errorf("failed to parse randomproduct.json: %w", err)
	}
	return nil
}

func (s *RandomProductService) GetAll(page, limit int) ([]public.RandomProduct, int, error) {
	total := len(s.products)
	offset := (page - 1) * limit
	if offset >= total {
		return []public.RandomProduct{}, total, nil
	}
	end := offset + limit
	if end > total {
		end = total
	}
	return s.products[offset:end], total, nil
}

func (s *RandomProductService) GetByID(id int) (*public.RandomProduct, error) {
	for _, product := range s.products {
		if product.ID == id {
			return &product, nil
		}
	}
	return nil, fmt.Errorf("product with ID %d not found", id)
}

func (s *RandomProductService) GetRandom() (*public.RandomProduct, error) {
	if len(s.products) == 0 {
		return nil, fmt.Errorf("no products available")
	}
	return &s.products[s.rand.Intn(len(s.products))], nil
}
