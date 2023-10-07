package models

type HasLinks struct {
	Tidal        bool `json:"tidal"`
	Amazon       bool `json:"amazon"`
	Deezer       bool `json:"deezer"`
	Itunes       bool `json:"itunes"`
	Napster      bool `json:"napster"`
	Pandora      bool `json:"pandora"`
	Spotify      bool `json:"spotify"`
	Youtube      bool `json:"youtube"`
	AmazonMusic  bool `json:"amazonMusic"`
	ItunesStore  bool `json:"itunesStore"`
	YoutubeMusic bool `json:"youtubeMusic"`
}
