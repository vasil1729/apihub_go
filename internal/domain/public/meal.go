package public

// Meal represents a meal from the dataset
type Meal struct {
	IDMeal           string   `json:"idMeal"`
	StrMeal          string   `json:"strMeal"`
	StrCategory      string   `json:"strCategory"`
	StrArea          string   `json:"strArea"`
	StrInstructions  string   `json:"strInstructions"`
	StrMealThumb     string   `json:"strMealThumb"`
	StrTags          *string  `json:"strTags"`
	StrYoutube       string   `json:"strYoutube"`
	ID               int      `json:"id"`
}
