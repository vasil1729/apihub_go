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

type rawBook struct {
	ID         int `json:"id"`
	VolumeInfo struct {
		Title       string   `json:"title"`
		Authors     []string `json:"authors"`
		Publisher   string   `json:"publisher"`
		Description string   `json:"description"`
		PageCount   int      `json:"pageCount"`
		Categories  []string `json:"categories"`
		ImageLinks  struct {
			Thumbnail string `json:"thumbnail"`
		} `json:"imageLinks"`
	} `json:"volumeInfo"`
}

type BookService struct {
	books []public.Book
	rand  *rand.Rand
}

func NewBookService(dataPath string) (*BookService, error) {
	service := &BookService{
		rand: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
	if err := service.loadBooks(dataPath); err != nil {
		return nil, err
	}
	return service, nil
}

func (s *BookService) loadBooks(dataPath string) error {
	data, err := os.ReadFile(filepath.Join(dataPath, "books.json"))
	if err != nil {
		return fmt.Errorf("failed to read books.json: %w", err)
	}
	
	var rawBooks []rawBook
	if err := json.Unmarshal(data, &rawBooks); err != nil {
		return fmt.Errorf("failed to parse books.json: %w", err)
	}
	
	for _, rb := range rawBooks {
		s.books = append(s.books, public.Book{
			ID:          rb.ID,
			Title:       rb.VolumeInfo.Title,
			Authors:     rb.VolumeInfo.Authors,
			Publisher:   rb.VolumeInfo.Publisher,
			Description: rb.VolumeInfo.Description,
			PageCount:   rb.VolumeInfo.PageCount,
			Categories:  rb.VolumeInfo.Categories,
			Thumbnail:   rb.VolumeInfo.ImageLinks.Thumbnail,
		})
	}
	return nil
}

func (s *BookService) GetAll(page, limit int) ([]public.Book, int, error) {
	total := len(s.books)
	offset := (page - 1) * limit
	if offset >= total {
		return []public.Book{}, total, nil
	}
	end := offset + limit
	if end > total {
		end = total
	}
	return s.books[offset:end], total, nil
}

func (s *BookService) GetByID(id int) (*public.Book, error) {
	for _, book := range s.books {
		if book.ID == id {
			return &book, nil
		}
	}
	return nil, fmt.Errorf("book with ID %d not found", id)
}

func (s *BookService) GetRandom() (*public.Book, error) {
	if len(s.books) == 0 {
		return nil, fmt.Errorf("no books available")
	}
	return &s.books[s.rand.Intn(len(s.books))], nil
}
