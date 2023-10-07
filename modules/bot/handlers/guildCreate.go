package handlers

import (
	"songwhip_bot/modules/actions"
	"songwhip_bot/modules/db/models"
	"songwhip_bot/modules/logging"

	"github.com/bwmarrin/discordgo"
)

func NewGuildCreateHandler(b Bot) func(s *discordgo.Session, g *discordgo.GuildCreate) {
	return func(s *discordgo.Session, g *discordgo.GuildCreate) {
		db := b.GetApp().DB

		// Log the guilds seen
		logging.PrintLog("- %s (id: %s)", g.Name, g.ID)
		actions.UpdateWatchStatus(s)

		// Check if the guild is already in the database, if so we're done here.
		existingGuild := &models.Guild{}
		db.First(existingGuild, "ID = ?", g.ID)
		if existingGuild.ID == "" {
			role, err := actions.EnsureAdminRole(s, g.Guild)

			if err != nil {
				logging.PrintLog("Error adding/getting admin role for/from guild: %s", err)
				return
			}

			// Add the guild to the database
			pk := db.Create(&models.Guild{
				ID:          g.ID,
				AdminRoleID: *role,
				JoinedAt:    g.JoinedAt,
			})

			if pk.Error != nil {
				logging.PrintLog("Error storing guild: %s", pk.Error)
			}

			logging.PrintLog("Guild seen for the first time: %s (id: %s) (Admin Role ID is %s)", g.Name, g.ID, role)
		}

		role, err := actions.EnsureAdminRole(s, g.Guild)

		if err != nil {
			logging.PrintLog("Error adding/getting admin role for/from guild: %s", err)
			return
		}
		db.Model(&models.Guild{}).Where("ID = ?", g.ID).Update("AdminRoleID", *role)
	}
}
