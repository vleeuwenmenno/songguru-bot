package models

type Config struct {
	Discord struct {
		BotToken          string `yaml:"bot_token"`
		ModeratorRoleName string `yaml:"moderator_role_name"`
	} `yaml:"discord"`
	DefaultGuildSettings struct {
		KeepOriginalMessage struct {
			Enabled       bool `yaml:"enabled"`
			AllowOverride bool `yaml:"allow_members_override"`
		} `yaml:"keep_original_messages"`
		MentionMode struct {
			Enabled       bool `yaml:"enabled"`
			AllowOverride bool `yaml:"allow_members_override"`
		} `yaml:"mention_mode"`
		SimpleMode struct {
			Enabled       bool `yaml:"enabled"`
			AllowOverride bool `yaml:"allow_members_override"`
		} `yaml:"simple_mode"`
	} `yaml:"default_guild_settings"`
	Intents      []string `yaml:"intents"`
	DatabasePath string   `yaml:"database_path"`
	WebPortal    struct {
		Port      int    `yaml:"port"`
		URL       string `yaml:"url"`
		DebugMode bool   `yaml:"debug_mode"`
	} `yaml:"web_portal"`
	NotifyMessages []struct {
		ID      string `yaml:"id"`
		Message string `yaml:"message"`
	} `yaml:"notify_messages"`
}
