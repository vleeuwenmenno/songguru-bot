package handlers

import "github.com/bwmarrin/discordgo"

func AddHandlers(dg *discordgo.Session) {
	dg.AddHandler(messageCreate)
	dg.AddHandler(ready)
}
