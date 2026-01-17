package kitchensink

// RedirectsService handles redirect logic
type RedirectsService struct{}

func NewRedirectsService() *RedirectsService {
	return &RedirectsService{}
}

// GetRedirectURL returns the URL to redirect to
func (s *RedirectsService) GetRedirectURL(to string) string {
	if to == "" {
		return "https://google.com"
	}
	return to
}
