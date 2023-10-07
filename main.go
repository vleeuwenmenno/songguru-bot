package main

import (
	"fmt"
	"songwhip_bot/models"
	"songwhip_bot/modules/bot"
	config "songwhip_bot/modules/config/app"
	"songwhip_bot/modules/config/services"
	"songwhip_bot/modules/db"
	dbModels "songwhip_bot/modules/db/models"
	"songwhip_bot/modules/logging"
	"songwhip_bot/modules/web"
	"sync"

	"github.com/gin-gonic/gin"
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
		if app.Config.WebPortal.DebugMode {
			gin.SetMode(gin.DebugMode)
		} else {
			gin.SetMode(gin.ReleaseMode)
		}

		web.StartWebPortal(router, app)
		wg.Done()
	}()

	bot.Start(app)

	wg.Wait()
}
