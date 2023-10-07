package handlers

import (
	"songwhip_bot/modules/songwhip"

	"github.com/bwmarrin/discordgo"
)

func NewMessageCreateHandler(b Bot) func(s *discordgo.Session, m *discordgo.MessageCreate) {
	return func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if m.Author.ID == s.State.User.ID {
			return
		}

		songwhip.ProcessMessage(b.GetApp(), s, m)
	}
}
