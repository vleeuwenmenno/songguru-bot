package musicapi

type Artist struct {
	ID         string `json:"id"`
	Title      string `json:"title"`
	URL        string `json:"url"`
	PictureURL string `json:"pictureUrl"`
}
