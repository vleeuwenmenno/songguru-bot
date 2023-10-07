package songwhip

import (
	"errors"
	"regexp"
	"songwhip_bot/models"
	dbModels "songwhip_bot/modules/db/models"
	"songwhip_bot/modules/logging"
	songwhipapi "songwhip_bot/modules/songwhip_api"
	"strings"

	"github.com/bwmarrin/discordgo"
	"gorm.io/gorm"
)

func ProcessMessage(app *models.App, session *discordgo.Session, event *discordgo.MessageCreate) {
	link, err := extractLink(event.Content)
	if err != nil {
		logging.PrintLog(err.Error(), err)
		return
	}

	for _, service := range app.Services.StreamingServices {
		if isValidLink(*link, service) {
			info, err := songwhipapi.GetInfo(*link)

			if err != nil {
				logging.PrintLog(err.Error(), err)
				return
			}

			deleteMessageMaybe(app.DB, session, event.Message)

			session.ChannelMessageSend(event.ChannelID, info.URL)
		}
	}
}

func deleteMessageMaybe(db *gorm.DB, session *discordgo.Session, message *discordgo.Message) {
	// Check if we should delete the original message based on the guild
	guildSettings := dbModels.GuildSetting{}
	err := db.Where("ID = ?", message.GuildID).First(&guildSettings)

	if err.Error != nil {
		logging.PrintLog("Error getting guild settings for guild %s, error: %s", message.GuildID, err.Error.Error())
		return
	}
	keepOriginalMessage := guildSettings.KeepOriginalMessage

	// Check if members on this guild are allowed to override this
	if guildSettings.AllowOverrideKeepOriginalMessage {
		memberSettings := dbModels.MemberSetting{}
		err := db.Where("ID = ?", message.Author.ID).First(&memberSettings)

		if err.Error == nil {
			// Check if the member requesting this has chosen to delete the message or not
			keepOriginalMessage = memberSettings.KeepOriginalMessage
		}
	}

	if !keepOriginalMessage {
		session.ChannelMessageDelete(message.ChannelID, message.ID)
	}
}

// Checks if a given link is one that we have a service for configured
func isValidLink(link string, service models.Service) bool {
	for _, url := range service.Urls {
		if strings.HasPrefix(link, url) {
			return true
		}
	}
	return false
}

func extractLink(messageContent string) (*string, error) {
	regex := regexp.MustCompile(`(http|ftp|https):\/\/([\w\-_]+(?:(?:\.[\w\-_]+)+))([\w\-\.,@?^=%&:/~\+#]*[\w\-\@?^=%&/~\+#])?`)
	url := regex.FindString(messageContent)

	if url == "" {
		return nil, errors.New("invalid link")
	}

	return &url, nil
}
