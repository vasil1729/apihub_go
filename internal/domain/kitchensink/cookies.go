package kitchensink

// CookieResponse represents a cookie in the response
type CookieResponse struct {
	Cookies map[string]string `json:"cookies"`
}
