package actions

import (
	"time"

	"github.com/bwmarrin/discordgo"

	"songguru_bot/models"
	"songguru_bot/modules/actions/notify"
	dbModels "songguru_bot/modules/db/models"
	"songguru_bot/modules/logging"
)

func EnsureGuildIsWatched(s *discordgo.Session, g *discordgo.Guild, app *models.App) {
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
			logging.PrintLog("- error storing guild: %s", pk.Error)
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
			logging.PrintLog("- error storing guild settings: %s", pk.Error)
		}

		logging.PrintLog("- %s (id: %s) seen for the first time!", g.Name, g.ID)
		notify.EnsureNotifyGuildOwner(s, app, g, notify.Welcome)

		// Save legacy notify to registry to prevent new guilds from getting the legacy message.
		notifyRegistry := dbModels.NoticeRegistry{
			NoticeID: string(notify.Legacy),
			GuildID:  g.ID,
			OwnerID:  g.OwnerID,
			SentAt:   time.Now(),
		}
		err := app.DB.Create(&notifyRegistry).Error
		if err != nil {
			logging.PrintLog("- error storing legacy notification registry for guild %s : %s", g.ID, pk.Error)
			return
		}
	}
}
