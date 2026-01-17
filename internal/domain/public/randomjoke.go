package public

// RandomJoke represents a random joke from the dataset
type RandomJoke struct {
	Categories []string `json:"categories"`
	ID         int      `json:"id"`
	Content    string   `json:"content"`
}
