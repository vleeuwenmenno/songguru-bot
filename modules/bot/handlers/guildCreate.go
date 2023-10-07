package handlers

import (
	"songwhip_bot/modules/actions"
	"songwhip_bot/modules/logging"

	"github.com/bwmarrin/discordgo"
)

func NewGuildCreateHandler(b Bot) func(s *discordgo.Session, g *discordgo.GuildCreate) {
	return func(s *discordgo.Session, g *discordgo.GuildCreate) {
		logging.PrintLog("- %s (id: %s)", g.Name, g.ID)
		actions.UpdateWatchStatus(s)

		actions.EnsureGuildIsWatched(g.Guild, b.GetApp())
		actions.EnsureAdminRoleExists(s, g.Guild, b.GetApp())
		actions.EnsureAdminRoleAssigned(s, g.Guild, b.GetApp())
	}
}
