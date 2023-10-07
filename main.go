package main

import (
	"fmt"
	"songwhip_bot/models"
	"songwhip_bot/modules/bot"
	config "songwhip_bot/modules/config/discord"
	"songwhip_bot/modules/config/services"
	"songwhip_bot/modules/db"
	"songwhip_bot/modules/logging"
)

func NewApp() (*models.App, error) {
	// Load services yaml
	services, err := services.GetStreamingServices()

	if err != nil {
		return nil, fmt.Errorf("error loading services yaml: %w", err)
	}

	// Load config yaml
	config, err := config.GetConfig()

	if err != nil {
		return nil, fmt.Errorf("error loading config yaml: %w", err)
	}

	// Setup database
	db, err := db.SetupDB(config)

	if err != nil {
		return nil, fmt.Errorf("error setting up database: %w", err)
	}

	return &models.App{
		Config:   config,
		Services: services,
		DB:       db,
	}, nil
}

func main() {
	app, err := NewApp()
	if err != nil {
		logging.PrintLog(err.Error(), err)
		panic(err)
	}

	bot, err := bot.NewBot(app)
	if err != nil {
		logging.PrintLog(err.Error(), err)
		panic(err)
	}

	bot.Start(app)
}
