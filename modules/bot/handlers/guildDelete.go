package handlers

import (
	"songwhip_bot/modules/actions"
	"songwhip_bot/modules/logging"

	"github.com/bwmarrin/discordgo"
)

func NewGuildDeleteHandler(b Bot) func(s *discordgo.Session, g *discordgo.GuildDelete) {
	return func(s *discordgo.Session, g *discordgo.GuildDelete) {
		logging.PrintLog("- %s (id: %s)", g.Name, g.ID)
		actions.UpdateWatchStatus(s)
	}
}
