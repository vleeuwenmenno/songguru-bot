package handlers

import (
	"github.com/bwmarrin/discordgo"

	"songguru_bot/modules/actions"
	dbModels "songguru_bot/modules/db/models"
	"songguru_bot/modules/logging"
)

func NewGuildRoleDeleteHandler(b Bot) func(session *discordgo.Session, event *discordgo.GuildRoleDelete) {
	return func(session *discordgo.Session, event *discordgo.GuildRoleDelete) {
		// Find the guild in the database
		existingGuild := dbModels.Guild{}
		err := b.GetApp().DB.Find(&existingGuild, "ID = ?", event.GuildID).Limit(1).Error
		if err != nil {
			logging.PrintLog("warning, received a role deleted event for a unknown guild %s : %s", event.GuildID, err)
			return
		}

		if event.RoleID == *existingGuild.AdminRoleID {
			actions.ClearGuildBotRoleID(event.GuildID, b.GetApp())
			logging.PrintLog("moderator role deleted for guild %s : %s", event.GuildID)
			return
		}
	}
}
