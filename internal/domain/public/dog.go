package public

// DogWeight represents weight measurements
type DogWeight struct {
	Imperial string `json:"imperial"`
	Metric   string `json:"metric"`
}

// DogHeight represents height measurements
type DogHeight struct {
	Imperial string `json:"imperial"`
	Metric   string `json:"metric"`
}

// DogImage represents dog image information
type DogImage struct {
	ID     string `json:"id"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
	URL    string `json:"url"`
}

// Dog represents a dog breed from the dataset
type Dog struct {
	Weight      DogWeight `json:"weight"`
	Height      DogHeight `json:"height"`
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	BredFor     string    `json:"bred_for,omitempty"`
	BreedGroup  string    `json:"breed_group,omitempty"`
	LifeSpan    string    `json:"life_span"`
	Temperament string    `json:"temperament"`
	Origin      string    `json:"origin,omitempty"`
	Image       DogImage  `json:"image"`
}
