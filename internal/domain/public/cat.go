package public

// CatWeight represents weight measurements
type CatWeight struct {
	Imperial string `json:"imperial"`
	Metric   string `json:"metric"`
}

// Cat represents a cat breed from the dataset
type Cat struct {
	Weight      CatWeight `json:"weight"`
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Temperament string    `json:"temperament"`
	Origin      string    `json:"origin"`
	Description string    `json:"description"`
	LifeSpan    string    `json:"life_span"`
	Image       string    `json:"image"`
}
