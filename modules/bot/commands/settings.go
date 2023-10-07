package commands

import (
	"songwhip_bot/models"
	dbModels "songwhip_bot/modules/db/models"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/google/uuid"
)

func settings(session *discordgo.Session, event *discordgo.InteractionCreate, app *models.App) error {
	db := app.DB
	config := app.Config
	token := uuid.Must(uuid.NewRandom()).String()
	webToken := dbModels.SettingsWebToken{
		ID:            token,
		GuildID:       event.Interaction.GuildID,
		DiscordUserID: event.Interaction.Member.User.ID,
		AdminToken:    false,
		ExpiresAt:     time.Now().Add(15 * time.Minute),
	}

	db.Create(&webToken)

	session.InteractionRespond(event.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "To change your settings click on [the following link](" + config.WebPortal.URL + "?t=" + webToken.ID + ")\n**This link will expire in 15 minutes.**",
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	})
	return nil
}
