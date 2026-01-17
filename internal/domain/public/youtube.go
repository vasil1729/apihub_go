package public

type YouTubeVideo struct {
	ID           int    `json:"id"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	ChannelTitle string `json:"channelTitle"`
	PublishedAt  string `json:"publishedAt"`
}
