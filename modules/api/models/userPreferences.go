package apimodels

type UserPreferences struct {
	SimpleMode          bool `json:"simpleMode"`
	MentionOnlyMode     bool `json:"mentionOnlyMode"`
	KeepOriginalMessage bool `json:"keepOriginalMessage"`
}
