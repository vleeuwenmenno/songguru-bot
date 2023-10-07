package models

type Config struct {
	Discord struct {
		BotToken      string `yaml:"bot_token"`
		AdminRoleName string `yaml:"admin_role_name"`
	} `yaml:"discord"`
	Intents []string `yaml:"intents"`
}
