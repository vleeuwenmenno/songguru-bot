package actions

import (
	"github.com/bwmarrin/discordgo"

	"songguru_bot/models"
	dbModels "songguru_bot/modules/db/models"
	"songguru_bot/modules/logging"
)

func EnsureGuildPreferences(s *discordgo.Session, g *discordgo.Guild, app *models.App) {
	db := app.DB
	existingGuildPreferences := dbModels.GuildSetting{}
	db.Find(&existingGuildPreferences, "ID = ?", g.ID).Limit(1)

	existingGuild := dbModels.Guild{}
	db.Find(&existingGuild, "ID = ?", g.ID).Limit(1)

	if existingGuildPreferences.ID == "" && existingGuild.ID != "" {
		// Add the guild to the database
		pk := db.Create(&dbModels.GuildSetting{
			ID:                      g.ID,
			SimpleMode:              app.Config.DefaultGuildSettings.SimpleMode.Enabled,
			AllowOverrideSimpleMode: app.Config.DefaultGuildSettings.SimpleMode.AllowOverride,

			MentionOnlyMode:              app.Config.DefaultGuildSettings.MentionMode.Enabled,
			AllowOverrideMentionOnlyMode: app.Config.DefaultGuildSettings.MentionMode.AllowOverride,

			KeepOriginalMessage:              app.Config.DefaultGuildSettings.KeepOriginalMessage.Enabled,
			AllowOverrideKeepOriginalMessage: app.Config.DefaultGuildSettings.KeepOriginalMessage.AllowOverride,

			GuildRefer: g.ID,
		})

		if pk.Error != nil {
			logging.PrintLog("- error storing default guild preferences: %s", pk.Error)
		}
		logging.PrintLog("- stored initial guild preferences for %s (id: %s)", g.Name, g.ID)
	}
}
