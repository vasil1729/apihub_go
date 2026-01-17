package public

type Book struct {
	ID          int      `json:"id"`
	Title       string   `json:"title"`
	Authors     []string `json:"authors"`
	Publisher   string   `json:"publisher"`
	Description string   `json:"description"`
	PageCount   int      `json:"pageCount"`
	Categories  []string `json:"categories"`
	Thumbnail   string   `json:"thumbnail"`
}
