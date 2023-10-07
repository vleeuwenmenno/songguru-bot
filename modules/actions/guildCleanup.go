package actions

import (
	"errors"
	"songwhip_bot/models"
	dbModels "songwhip_bot/modules/db/models"
	"songwhip_bot/modules/logging"

	"github.com/bwmarrin/discordgo"
)

func GuildsCleanup(app *models.App, discordGuilds []*discordgo.Guild) {
	db := app.DB

	// Look through all guilds in the database and check if they are still in the variable guilds. If not we should drop that guild.
	allGuilds := []dbModels.Guild{}
	db.Find(&allGuilds)

	for _, guild := range allGuilds {
		_, err := findGuildInGuilds(guild.ID, discordGuilds)

		if err != nil {
			// Guild no longer exists
			db.Delete(&dbModels.Guild{}, guild.ID)
			logging.PrintLog("purged guild %s as it no longer exists", guild.ID)
		}
	}
}

func findGuildInGuilds(guildId string, guilds []*discordgo.Guild) (*discordgo.Guild, error) {
	for _, guild := range guilds {
		if guild.ID == guildId {
			return guild, nil
		}
	}
	return nil, errors.New("guild not found")
}
