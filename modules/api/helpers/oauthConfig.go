package helpers

import (
	"github.com/ravener/discord-oauth2"
	"golang.org/x/oauth2"

	"songguru_bot/models"
)

func GetOAuth2Config(app *models.App) *oauth2.Config {
	return &oauth2.Config{
		Endpoint:     discord.Endpoint,
		Scopes:       []string{discord.ScopeIdentify, discord.ScopeGuilds},
		RedirectURL:  app.Config.Discord.RedirectURL,
		ClientID:     app.Config.Discord.ClientID,
		ClientSecret: app.Config.Discord.ClientSecret,
	}
}
