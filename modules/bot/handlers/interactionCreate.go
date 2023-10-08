package handlers

import (
	"songguru_bot/modules/bot/commands"

	"github.com/bwmarrin/discordgo"
)

func NewInteractionCreateHandler(b Bot) func(session *discordgo.Session, event *discordgo.InteractionCreate) {
	return func(session *discordgo.Session, event *discordgo.InteractionCreate) {
		command := commands.EvaluateCommand(event.ApplicationCommandData().Name)
		err := command(session, event, b.GetApp())

		if err != nil {
			session.InteractionRespond(event.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "`Something went wrong: " + err.Error() + "`",
					Flags:   discordgo.MessageFlagsEphemeral,
				},
			})
		}
	}
}
