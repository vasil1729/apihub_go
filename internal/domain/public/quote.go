package public

// Quote represents a quote from the dataset
type Quote struct {
	Author       string   `json:"author"`
	Content      string   `json:"content"`
	Tags         []string `json:"tags"`
	AuthorSlug   string   `json:"authorSlug"`
	Length       int      `json:"length"`
	DateAdded    string   `json:"dateAdded"`
	DateModified string   `json:"dateModified"`
	ID           int      `json:"id"`
}
