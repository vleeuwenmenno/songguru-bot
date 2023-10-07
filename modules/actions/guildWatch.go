package actions

import (
	"songwhip_bot/models"
	dbModels "songwhip_bot/modules/db/models"
	"songwhip_bot/modules/logging"

	"github.com/bwmarrin/discordgo"
)

func EnsureGuildIsWatched(g *discordgo.Guild, app *models.App) {
	db := app.DB
	config := app.Config
	existingGuild := dbModels.Guild{}
	db.Find(&existingGuild, "ID = ?", g.ID).Limit(1)

	if existingGuild.ID == "" {
		// Add the guild to the database
		pk := db.Create(&dbModels.Guild{
			ID:       g.ID,
			JoinedAt: g.JoinedAt,
		})

		if pk.Error != nil {
			logging.PrintLog("Error storing guild: %s", pk.Error)
		}

		guildSettingsPk := db.Create(&dbModels.GuildSetting{
			ID:                               g.ID,
			GuildRefer:                       g.ID,
			MentionOnlyMode:                  config.DefaultGuildSettings.MentionMode.Enabled,
			AllowOverrideMentionOnlyMode:     config.DefaultGuildSettings.MentionMode.AllowOverride,
			SimpleMode:                       config.DefaultGuildSettings.SimpleMode.Enabled,
			AllowOverrideSimpleMode:          config.DefaultGuildSettings.SimpleMode.AllowOverride,
			KeepOriginalMessage:              config.DefaultGuildSettings.KeepOriginalMessage.Enabled,
			AllowOverrideKeepOriginalMessage: config.DefaultGuildSettings.KeepOriginalMessage.AllowOverride,
		})

		if guildSettingsPk.Error != nil {
			logging.PrintLog("Error storing guild settings: %s", pk.Error)
		}

		logging.PrintLog("Guild seen for the first time: %s (id: %s)", g.Name, g.ID)
	}
}
