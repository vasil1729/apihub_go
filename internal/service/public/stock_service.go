package public

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	
	"github.com/ultimatum/apihub_go/internal/domain/public"
)

// StockService handles stock operations
type StockService struct {
	stocks []public.Stock
}

// NewStockService creates a new stock service
func NewStockService(dataPath string) (*StockService, error) {
	service := &StockService{}
	
	// Load stocks from JSON file
	if err := service.loadStocks(dataPath); err != nil {
		return nil, err
	}
	
	return service, nil
}

// loadStocks loads stocks from the JSON file
func (s *StockService) loadStocks(dataPath string) error {
	filePath := filepath.Join(dataPath, "nse-stocks.json")
	
	data, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to read nse-stocks.json: %w", err)
	}
	
	if err := json.Unmarshal(data, &s.stocks); err != nil {
		return fmt.Errorf("failed to parse nse-stocks.json: %w", err)
	}
	
	return nil
}

// GetAll returns all stocks with pagination
func (s *StockService) GetAll(page, limit int) ([]public.Stock, int, error) {
	total := len(s.stocks)
	
	// Calculate offset
	offset := (page - 1) * limit
	
	// Validate pagination
	if offset >= total {
		return []public.Stock{}, total, nil
	}
	
	// Calculate end index
	end := offset + limit
	if end > total {
		end = total
	}
	
	return s.stocks[offset:end], total, nil
}

// GetBySymbol returns a stock by its symbol (case-insensitive)
func (s *StockService) GetBySymbol(symbol string) (*public.Stock, error) {
	symbolUpper := strings.ToUpper(symbol)
	for _, stock := range s.stocks {
		if strings.ToUpper(stock.Symbol) == symbolUpper {
			return &stock, nil
		}
	}
	return nil, fmt.Errorf("stock with symbol %s not found", symbol)
}
