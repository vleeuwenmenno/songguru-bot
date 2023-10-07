package main

type AppConfig struct {
	Discord struct {
		BotToken string `yaml:"bot_token"`
	} `yaml:"discord"`
}
