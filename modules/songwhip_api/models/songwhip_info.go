package models

type SongwhipInfo struct {
	Type           string      `json:"type"`
	ID             int         `json:"id"`
	Path           string      `json:"path"`
	Name           string      `json:"name"`
	URL            string      `json:"url"`
	SourceURL      string      `json:"sourceUrl"`
	SourceCountry  string      `json:"sourceCountry"`
	Image          string      `json:"image"`
	Links          HasLinks    `json:"links"`
	LinksCountries []string    `json:"linksCountries"`
	Artists        []Artist    `json:"artists"`
	Overrides      interface{} `json:"overrides"`
}
