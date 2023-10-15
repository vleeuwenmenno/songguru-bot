package handlers

import (
	"github.com/bwmarrin/discordgo"

	"songguru_bot/modules/actions"
	"songguru_bot/modules/actions/notify"
	"songguru_bot/modules/logging"
)

func NewGuildCreateHandler(b Bot) func(s *discordgo.Session, g *discordgo.GuildCreate) {
	return func(s *discordgo.Session, g *discordgo.GuildCreate) {
		logging.PrintLog("- %s (id: %s)", g.Name, g.ID)
		actions.UpdateWatchStatus(s)

		var role *discordgo.Role

		// Ensure basics
		actions.EnsureGuildIsWatched(s, g.Guild, b.GetApp())
		actions.EnsureGuildPreferences(s, g.Guild, b.GetApp())
		actions.EnsureBotRoleExists(&role, s, g.Guild, b.GetApp())

		if role == nil {
			notify.EnsureNotifyGuildOwner(s, b.GetApp(), g.Guild, notify.BotRoleCreateFailed)
		}

		if role != nil {
			actions.EnsureGuildOwnerHasBotRole(s, g.Guild, b.GetApp())
			actions.EnsureBotHasBotRole(s, g.Guild, b.GetApp())
		}

		// Send out legacy guilds a message (New guilds automatically get the legacy notification marked as already sent out)
		notify.EnsureNotifyGuildOwner(s, b.GetApp(), g.Guild, notify.Legacy)
	}
}
