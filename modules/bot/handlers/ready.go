package handlers

import (
	"songwhip_bot/modules/actions"
	"songwhip_bot/modules/logging"

	"github.com/bwmarrin/discordgo"
)

func NewReadyHandler(b Bot) func(s *discordgo.Session, r *discordgo.Ready) {
	return func(s *discordgo.Session, r *discordgo.Ready) {
		logging.PrintLog("Logged in as: %v#%v", s.State.User.Username, s.State.User.Discriminator)
		logging.PrintLog("Bot is serving in %d servers!", len(s.State.Guilds))

		actions.GuildsCleanup(b.GetApp(), s.State.Guilds)
	}
}
