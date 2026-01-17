package public

// Stock represents a stock from the NSE dataset
type Stock struct {
	Name           string `json:"Name"`
	Symbol         string `json:"Symbol"`
	ListingDate    string `json:"ListingDate"`
	ISIN           string `json:"ISIN"`
	MarketCap      string `json:"MarketCap"`
	CurrentPrice   string `json:"CurrentPrice"`
	HighLow        string `json:"HighLow"`
	StockPE        string `json:"StockPE"`
	BookValue      string `json:"BookValue"`
	DividendYield  string `json:"DividendYield"`
	ROCE           string `json:"ROCE"`
	ROE            string `json:"ROE"`
	FaceValue      string `json:"FaceValue"`
}
