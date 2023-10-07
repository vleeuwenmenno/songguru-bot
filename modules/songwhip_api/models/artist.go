package models

import "time"

type Artist struct {
	Type           string     `json:"type"`
	ID             int        `json:"id"`
	Path           string     `json:"path"`
	Name           string     `json:"name"`
	SourceURL      string     `json:"sourceUrl"`
	SourceCountry  string     `json:"sourceCountry"`
	URL            string     `json:"url"`
	Image          string     `json:"image"`
	CreatedAt      time.Time  `json:"createdAt"`
	UpdatedAt      time.Time  `json:"updatedAt"`
	RefreshedAt    time.Time  `json:"refreshedAt"`
	LinksCountries []string   `json:"linksCountries"`
	Links          Links      `json:"links"`
	Description    string     `json:"description"`
	ServiceIds     ServiceIds `json:"serviceIds"`
}
