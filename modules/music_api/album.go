package musicapi

type Album struct {
	ID         string   `json:"id"`
	Title      string   `json:"title"`
	Artists    []Artist `json:"artist"`
	URL        string   `json:"url"`
	PictureURL string   `json:"pictureUrl"`
}
