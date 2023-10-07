package actions

import (
	"songwhip_bot/models"
	dbModels "songwhip_bot/modules/db/models"
	"songwhip_bot/modules/logging"

	"github.com/bwmarrin/discordgo"
)

func EnsureGuildIsWatched(g *discordgo.Guild, app *models.App) {
	db := app.DB
	existingGuild := &dbModels.Guild{}
	db.First(existingGuild, "ID = ?", g.ID)
	if existingGuild.ID == "" {
		// Add the guild to the database
		pk := db.Create(&dbModels.Guild{
			ID:       g.ID,
			JoinedAt: g.JoinedAt,
		})

		if pk.Error != nil {
			logging.PrintLog("Error storing guild: %s", pk.Error)
		}

		logging.PrintLog("Guild seen for the first time: %s (id: %s)", g.Name, g.ID)
	}
}
