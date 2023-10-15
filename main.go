package main

import (
	"fmt"
	"sync"

	"github.com/gin-gonic/gin"

	"songguru_bot/models"
	"songguru_bot/modules/api"
	"songguru_bot/modules/bot"
	config "songguru_bot/modules/config/app"
	"songguru_bot/modules/config/services"
	"songguru_bot/modules/db"
	dbModels "songguru_bot/modules/db/models"
	"songguru_bot/modules/logging"
	"songguru_bot/modules/portal"
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

	// Clear all web tokens
	db.Unscoped().Where("1 = 1").Delete(&dbModels.SettingsWebToken{})

	return &models.App{
		Config:   config,
		Services: services,
		States:   &models.States{Memory: []string{}},
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

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		router := gin.Default()
		portal.StartWebPortal(router, app)
		wg.Done()
	}()
	wg.Add(1)

	go func() {
		router := gin.Default()
		api.StartAPI(router, app)
		wg.Done()
	}()

	bot.Start(app)

	wg.Wait()
}
