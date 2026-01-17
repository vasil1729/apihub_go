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

// QuoteService handles quote operations
type QuoteService struct {
	quotes []public.Quote
	rand   *rand.Rand
}

// NewQuoteService creates a new quote service
func NewQuoteService(dataPath string) (*QuoteService, error) {
	service := &QuoteService{
		rand: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
	
	// Load quotes from JSON file
	if err := service.loadQuotes(dataPath); err != nil {
		return nil, err
	}
	
	return service, nil
}

// loadQuotes loads quotes from the JSON file
func (s *QuoteService) loadQuotes(dataPath string) error {
	filePath := filepath.Join(dataPath, "quotes.json")
	
	data, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to read quotes.json: %w", err)
	}
	
	if err := json.Unmarshal(data, &s.quotes); err != nil {
		return fmt.Errorf("failed to parse quotes.json: %w", err)
	}
	
	return nil
}

// GetAll returns all quotes with pagination
func (s *QuoteService) GetAll(page, limit int) ([]public.Quote, int, error) {
	total := len(s.quotes)
	
	// Calculate offset
	offset := (page - 1) * limit
	
	// Validate pagination
	if offset >= total {
		return []public.Quote{}, total, nil
	}
	
	// Calculate end index
	end := offset + limit
	if end > total {
		end = total
	}
	
	return s.quotes[offset:end], total, nil
}

// GetByID returns a quote by ID
func (s *QuoteService) GetByID(id int) (*public.Quote, error) {
	for _, quote := range s.quotes {
		if quote.ID == id {
			return &quote, nil
		}
	}
	return nil, fmt.Errorf("quote with ID %d not found", id)
}

// GetRandom returns a random quote
func (s *QuoteService) GetRandom() (*public.Quote, error) {
	if len(s.quotes) == 0 {
		return nil, fmt.Errorf("no quotes available")
	}
	
	// Return a truly random quote
	index := s.rand.Intn(len(s.quotes))
	return &s.quotes[index], nil
}
