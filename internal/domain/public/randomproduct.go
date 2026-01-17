package public

type RandomProduct struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	Category    string  `json:"category"`
	Image       string  `json:"thumbnail"`
	Rating      float64 `json:"rating"`
	Brand       string  `json:"brand"`
	Stock       int     `json:"stock"`
}
