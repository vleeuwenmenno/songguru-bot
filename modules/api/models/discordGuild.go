package apimodels

type DiscordGuild struct {
	ID             string   `json:"id"`
	Name           string   `json:"name"`
	Icon           string   `json:"icon"`
	Owner          bool     `json:"owner"`
	Permissions    int      `json:"permissions"`
	PermissionsNew string   `json:"permissions_new"`
	Features       []string `json:"features"`
}
