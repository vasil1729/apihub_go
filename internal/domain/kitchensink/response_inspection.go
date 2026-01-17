package kitchensink

// ResponseInspectionResponse is a generic struct for JSON/XML responses
type ResponseInspectionResponse struct {
	Message string `json:"message" xml:"message"`
	Format  string `json:"format" xml:"format"`
}
