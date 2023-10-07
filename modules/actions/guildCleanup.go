package actions

import (
	"songwhip_bot/models"
	dbModels "songwhip_bot/modules/db/models"
	"songwhip_bot/modules/logging"

	"github.com/bwmarrin/discordgo"
)

func GuildsCleanup(app *models.App, guilds []*discordgo.Guild) {
	db := app.DB

	logging.PrintLog("Purging guilds from database that bot no longer has access to ...")

	// Clean up guilds that are no longer in s.State.Guilds (e.g. bot got kicked while offline)
	for _, guild := range guilds {
		existingGuild := &dbModels.Guild{}
		db.First(existingGuild, "ID = ?", guild.ID)
		if existingGuild.ID != "" {
			continue
		}

		db.Delete(existingGuild)
		logging.PrintLog("- purged %s", guild.ID)
	}
}
