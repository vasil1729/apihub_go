package kitchensink

// RequestInspectionResponse represents the response for request inspection
type RequestInspectionResponse struct {
	IP        string              `json:"ip"`
	UserAgent string              `json:"userAgent"`
	Headers   map[string][]string `json:"headers"`
	Method    string              `json:"method"`
	URL       string              `json:"url"`
}

// IPResponse represents just the IP address
type IPResponse struct {
	IP string `json:"ip"`
}

// UserAgentResponse represents just the User-Agent
type UserAgentResponse struct {
	UserAgent string `json:"userAgent"`
}

// HeadersResponse represents just the headers
type HeadersResponse struct {
	Headers map[string][]string `json:"headers"`
}
