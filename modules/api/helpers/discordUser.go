package helpers

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"

	apimodels "songguru_bot/modules/api/models"
	"songguru_bot/modules/logging"
)

func GetDiscordUser(conf *oauth2.Config, token *oauth2.Token, c *gin.Context) (*apimodels.DiscordUser, error) {
	client := conf.Client(context.Background(), token)
	res, err := client.Get("https://discord.com/api/users/@me")

	if err != nil || res.StatusCode != http.StatusOK {
		return nil, err
	}

	defer res.Body.Close()
	bytes, err := io.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	var discordUser apimodels.DiscordUser
	json.Unmarshal(bytes, &discordUser)
	return &discordUser, nil
}

// TODO: This is limited to 200, we could someday add more but who the f has 200 guilds!?
func GetDiscordUserGuilds(conf *oauth2.Config, token *oauth2.Token) ([]apimodels.DiscordGuild, error) {
	client := conf.Client(context.Background(), token)
	res, err := client.Get("https://discord.com/api/users/@me/guilds")

	if err != nil || res.StatusCode != http.StatusOK {
		return nil, err
	}

	defer res.Body.Close()
	bytes, err := io.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	var guilds []apimodels.DiscordGuild
	logging.PrintLog(string(bytes))
	err = json.Unmarshal(bytes, &guilds)
	if err != nil {
		return nil, err
	}

	return guilds, nil
}
