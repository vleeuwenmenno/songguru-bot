package commands

import (
	"songwhip_bot/models"
	"songwhip_bot/modules/actions"
	dbModels "songwhip_bot/modules/db/models"
	"songwhip_bot/modules/logging"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/google/uuid"
)

func settings(session *discordgo.Session, event *discordgo.InteractionCreate, app *models.App) error {
	db := app.DB
	config := app.Config
	token := uuid.Must(uuid.NewRandom()).String()

	// If requester has role admin
	roles, err := session.GuildRoles(event.GuildID)

	if err != nil {
		logging.PrintLog("Error getting admin roles for guild (ID: %s)", event.GuildID)
		return err
	}

	adminRoleID, err := actions.GuildHasAdminRole(roles, config)
	if err != nil {
		logging.PrintLog("Error getting admin roles for guild (ID: %s)", event.GuildID)
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
		if role == adminRoleID {
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
			Content: "To change your settings click on [the following link](" + config.WebPortal.URL + "?t=" + webToken.ID + ")\n**This link will expire in 15 minutes.**",
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	})
	return nil
}
