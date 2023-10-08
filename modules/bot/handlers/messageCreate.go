package handlers

import (
	"songguru_bot/modules/bot/processing"

	"github.com/bwmarrin/discordgo"
)

func NewMessageCreateHandler(b Bot) func(s *discordgo.Session, m *discordgo.MessageCreate) {
	return func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if m.Author.ID == s.State.User.ID {
			return
		}

		processing.ProcessMessage(b.GetApp(), s, m)
	}
}
