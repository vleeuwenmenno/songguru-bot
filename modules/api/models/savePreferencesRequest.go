package apimodels

type SavePreferencesRequest struct {
	UserPreferences  UserPreferences  `json:"userPreferences"`
	GuildPreferences GuildPreferences `json:"guildPreferences"`
	GuildID          string           `json:"guildId"`
}
