package kitchensink

// StatusCodeResponse represents the response when a specific status code is requested
type StatusCodeResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
