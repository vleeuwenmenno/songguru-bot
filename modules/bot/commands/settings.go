package commands

import (
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/google/uuid"

	"songguru_bot/models"
	"songguru_bot/modules/actions"
	dbModels "songguru_bot/modules/db/models"
	"songguru_bot/modules/logging"
)

func settings(session *discordgo.Session, event *discordgo.InteractionCreate, app *models.App) error {
	db := app.DB
	config := app.Config
	token := uuid.Must(uuid.NewRandom()).String()

	// If requester has bot role
	guild, err := session.Guild(event.GuildID)

	if err != nil {
		logging.PrintLog("Error getting bot guild (ID: %s)", event.GuildID)
		return err
	}

	BotRole, err := actions.GetBotRole(session, guild, app.Config.Discord.ModeratorRoleName)
	if err != nil {
		logging.PrintLog("Error getting bot role for guild (ID: %s)", event.GuildID)
		return err
	}

	webToken := dbModels.SettingsWebToken{
		ID:            token,
		GuildID:       event.Interaction.GuildID,
		DiscordUserID: event.Interaction.Member.User.ID,
		AdminToken:    false,
		ExpiresAt:     time.Now().Add(15 * time.Minute),
	}

	for _, role := range event.Member.Roles {
		if role == BotRole.ID {
			webToken = dbModels.SettingsWebToken{
				ID:            token,
				GuildID:       event.GuildID,
				DiscordUserID: event.Member.User.ID,
				AdminToken:    true,
				ExpiresAt:     time.Now().Add(15 * time.Minute),
			}
			break
		}
	}

	db.Create(&webToken)

	session.InteractionRespond(event.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "To change your settings click on [the following link](" + config.WebPortal.URL + "?t=" + webToken.ID + ")\n**Don't share this link with anyone! The link will expire in 15 minutes.**",
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	})
	return nil
}
