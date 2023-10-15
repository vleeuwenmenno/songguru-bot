package apimodels

type GuildPreferences struct {
	SimpleMode                       bool `json:"simpleMode"`
	MentionOnlyMode                  bool `json:"mentionOnlyMode"`
	KeepOriginalMessage              bool `json:"keepOriginalMessage"`
	AllowOverrideSimpleMode          bool `json:"allowOverridesimpleMode"`
	AllowOverrideMentionOnlyMode     bool `json:"allowOverridementionOnlyMode"`
	AllowOverrideKeepOriginalMessage bool `json:"allowOverridekeepOriginalMessage"`
}
