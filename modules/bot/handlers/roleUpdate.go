package handlers

import (
	"github.com/bwmarrin/discordgo"

	"songguru_bot/models"
	dbModels "songguru_bot/modules/db/models"
	"songguru_bot/modules/logging"
)

func NewGuildRoleUpdateHandler(b Bot) func(session *discordgo.Session, event *discordgo.GuildRoleUpdate) {
	return func(session *discordgo.Session, event *discordgo.GuildRoleUpdate) {
		CheckGuildRole(event.Role, event.GuildID, b.GetApp())
	}
}

func CheckGuildRole(role *discordgo.Role, guildID string, app *models.App) {
	if role != nil && role.Name == app.Config.Discord.ModeratorRoleName {
		// Find the guild in the database
		existingGuild := dbModels.Guild{}
		err := app.DB.Find(&existingGuild, "ID = ?", guildID).Limit(1).Error
		if err != nil {
			logging.PrintLog("warning, received a role created event for a unknown guild %s : %s", guildID, err)
			return
		}

		// Update the AdminRoleId column on the guild table
		existingGuild.AdminRoleID = &role.ID
		err = app.DB.Save(&existingGuild).Error
		if err != nil {
			logging.PrintLog("failed to store new role id for guild %s : %s", guildID, err)
			return
		}
		logging.PrintLog("received and stored new moderator role id %s for guild %s : %s", role.ID, guildID)
	}
}
