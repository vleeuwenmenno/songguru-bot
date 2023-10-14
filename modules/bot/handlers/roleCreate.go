package handlers

import (
	"github.com/bwmarrin/discordgo"
)

func NewGuildRoleCreateHandler(b Bot) func(session *discordgo.Session, event *discordgo.GuildRoleCreate) {
	return func(session *discordgo.Session, event *discordgo.GuildRoleCreate) {
		CheckGuildRole(event.Role, event.GuildID, b.GetApp())
	}
}
