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
	mayContinue := evaluateMentionOnlyMode(app, session, event.Message)

	if !mayContinue {
		return
	}

	link, err := extractLink(event.Content)
	if err != nil {
		logging.PrintLog(err.Error(), err)
		return
	}

	for _, service := range app.Services.StreamingServices {
		if isValidLink(*link, service) {
			info, err := songwhipapi.GetInfo(*link)

			if err != nil {
				if strings.Contains(err.Error(), "429 Too Many Requests") {
					session.ChannelMessageSendReply(event.ChannelID, "Songwhip is currently rate limiting requests, please try again later", event.Message.Reference())
				}

				logging.PrintLog(err.Error(), err)
				return
			}

			deleteMessageMaybe(app.DB, session, event.Message)
			session.ChannelMessageSend(event.ChannelID, info.URL)
			return
		}
	}
}

func evaluateMentionOnlyMode(app *models.App, session *discordgo.Session, message *discordgo.Message) bool {
	db := app.DB

	// Check if mention only mode is on in this guild
	guildSettings := dbModels.GuildSetting{}
	err := db.Where("ID = ?", message.GuildID).First(&guildSettings)

	if err.Error != nil {
		logging.PrintLog("Error getting guild settings for guild %s, error: %s", message.GuildID, err.Error.Error())
		return false
	}
	mentionOnlyMode := guildSettings.MentionOnlyMode

	// Check if members on this guild are allowed to override this setting
	if guildSettings.AllowOverrideMentionOnlyMode {
		memberSettings := dbModels.MemberSetting{}
		affected := db.Find(&memberSettings, "ID = ?", message.Author.ID).Limit(1).RowsAffected

		if affected > 0 {
			// Check if the member requesting this has chosen to mention only mode or not
			mentionOnlyMode = memberSettings.MentionOnlyMode
		}
	}

	if mentionOnlyMode {
		return strings.Contains(message.Content, session.State.User.Mention())
	}

	return true
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
		affected := db.Find(&memberSettings, "ID = ?", message.Author.ID).Limit(1).RowsAffected

		if affected > 0 {
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
