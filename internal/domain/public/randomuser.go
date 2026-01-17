package public

// RandomUser represents a random user from the dataset
type RandomUser struct {
	Gender   string       `json:"gender"`
	Name     Name         `json:"name"`
	Location Location     `json:"location"`
	Email    string       `json:"email"`
	Login    Login        `json:"login"`
	DOB      DateInfo     `json:"dob"`
	Registered DateInfo   `json:"registered"`
	Phone    string       `json:"phone"`
	Cell     string       `json:"cell"`
	ID       int          `json:"id"`
	Picture  Picture      `json:"picture"`
	Nat      string       `json:"nat"`
}

// Name represents a person's name
type Name struct {
	Title string `json:"title"`
	First string `json:"first"`
	Last  string `json:"last"`
}

// Location represents a person's location
type Location struct {
	Street      Street      `json:"street"`
	City        string      `json:"city"`
	State       string      `json:"state"`
	Country     string      `json:"country"`
	Postcode    interface{} `json:"postcode"` // Can be string or int
	Coordinates Coordinates `json:"coordinates"`
	Timezone    Timezone    `json:"timezone"`
}

// Street represents a street address
type Street struct {
	Number int    `json:"number"`
	Name   string `json:"name"`
}

// Coordinates represents geographic coordinates
type Coordinates struct {
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}

// Timezone represents timezone information
type Timezone struct {
	Offset      string `json:"offset"`
	Description string `json:"description"`
}

// Login represents login credentials
type Login struct {
	UUID     string `json:"uuid"`
	Username string `json:"username"`
	Password string `json:"password"`
	Salt     string `json:"salt"`
	MD5      string `json:"md5"`
	SHA1     string `json:"sha1"`
	SHA256   string `json:"sha256"`
}

// DateInfo represents date and age information
type DateInfo struct {
	Date string `json:"date"`
	Age  int    `json:"age"`
}

// Picture represents profile pictures
type Picture struct {
	Large     string `json:"large"`
	Medium    string `json:"medium"`
	Thumbnail string `json:"thumbnail"`
}
