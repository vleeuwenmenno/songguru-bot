package models

type Config struct {
	Discord struct {
		BotToken string `yaml:"bot_token"`
	} `yaml:"discord"`
	Intents []string `yaml:"intents"`
}
