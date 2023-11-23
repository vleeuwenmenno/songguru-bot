package api

import (
	"fmt"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"songguru_bot/models"
	"songguru_bot/modules/logging"
)

func StartAPI(router *gin.Engine, app *models.App) {
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{app.Config.API.URL, app.Config.WebPortal.URL},
		AllowMethods:     []string{"GET", "POST", "PUT"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           1 * time.Hour,
	}))

	api := router.Group("/api")
	{
		setupRoutes(api, app)
	}

	logging.PrintLog("starting API on http://0.0.0.0:%d", app.Config.API.Port)
	router.Run(fmt.Sprintf(":%d", app.Config.API.Port))
}
