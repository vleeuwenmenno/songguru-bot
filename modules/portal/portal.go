package portal

import (
	"fmt"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"

	"songguru_bot/models"
	"songguru_bot/modules/logging"
)

func StartWebPortal(router *gin.Engine, app *models.App) {
	router.Use(static.Serve("/", static.LocalFile("./client/dist", true)))
	logging.PrintLog("starting web portal on http://0.0.0.0:%d", app.Config.WebPortal.Port)
	router.Run(fmt.Sprintf(":%d", app.Config.WebPortal.Port))
}
