package kitchensink

// HTTPMethodResponse represents the response for HTTP method tests
type HTTPMethodResponse struct {
	Method        string              `json:"method"`
	URL           string              `json:"url"`
	Headers       map[string][]string `json:"headers"`
	Query         map[string][]string `json:"query"`
	Body          interface{}         `json:"body,omitempty"`
	Origin        string              `json:"origin"`
	ContentLength int64               `json:"contentLength"`
	UserAgent     string              `json:"userAgent"`
}
