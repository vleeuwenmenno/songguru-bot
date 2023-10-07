package commands

import (
	"songwhip_bot/models"

	"github.com/bwmarrin/discordgo"
)

func changelog(session *discordgo.Session, event *discordgo.InteractionCreate, app *models.App) error {
	config := app.Config

	session.InteractionRespond(event.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "You can find the changelog for the bot at [the following link](" + config.WebPortal.URL + "/changelogs)",
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	})
	return nil
}
