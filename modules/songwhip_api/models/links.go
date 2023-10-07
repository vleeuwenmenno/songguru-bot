package models

type Links struct {
	Tidal        []ServiceLink `json:"tidal"`
	Amazon       []ServiceLink `json:"amazon"`
	Deezer       []ServiceLink `json:"deezer"`
	Itunes       []ServiceLink `json:"itunes"`
	Napster      []ServiceLink `json:"napster"`
	Pandora      []ServiceLink `json:"pandora"`
	Spotify      []ServiceLink `json:"spotify"`
	Youtube      []ServiceLink `json:"youtube"`
	AmazonMusic  []ServiceLink `json:"amazonMusic"`
	ItunesStore  []ServiceLink `json:"itunesStore"`
	YoutubeMusic []ServiceLink `json:"youtubeMusic"`
	Yandex       []ServiceLink `json:"yandex"`
	Discogs      []ServiceLink `json:"discogs"`
	Twitter      []ServiceLink `json:"twitter"`
	Facebook     []ServiceLink `json:"facebook"`
	Instagram    []ServiceLink `json:"instagram"`
	Wikipedia    []ServiceLink `json:"wikipedia"`
	Soundcloud   []ServiceLink `json:"soundcloud"`
	MusicBrainz  []ServiceLink `json:"musicBrainz"`
}
