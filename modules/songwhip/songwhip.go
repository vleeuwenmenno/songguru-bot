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
	mayContinue, mentionOnlyMode := evaluateMentionOnlyMode(app, session, event.Message)
	shouldDeleteMessage := evaluateDeleteMessage(app.DB, session, event.Message)
	simpleMode := evaluateSimpleMode(app, session, event.Message)

	logging.PrintLog("message for guild %s delete? %v, simple mode? %v, mention only mode? %v, continue? %v", event.Message.GuildID, shouldDeleteMessage, simpleMode, mentionOnlyMode, mayContinue)

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

			if shouldDeleteMessage {
				session.ChannelMessageDelete(event.Message.ChannelID, event.Message.ID)
			}

			if simpleMode {
				session.ChannelMessageSend(event.ChannelID, info.URL)
				return
			}

			// TODO: Add support for embeds
			session.ChannelMessageSend(event.ChannelID, "WIP ... not simple mode: "+info.URL)
			return
		}
	}
}

func evaluateSimpleMode(app *models.App, session *discordgo.Session, message *discordgo.Message) bool {
	db := app.DB

	// Check if mention only mode is on in this guild
	guildSettings := dbModels.GuildSetting{}
	err := db.Where("ID = ?", message.GuildID).First(&guildSettings)

	if err.Error != nil {
		logging.PrintLog("Error getting guild settings for guild %s, error: %s", message.GuildID, err.Error.Error())
		return false
	}
	simpleMode := guildSettings.SimpleMode

	// Check if members on this guild are allowed to override this setting
	if guildSettings.AllowOverrideSimpleMode {
		memberSettings := dbModels.MemberSetting{}
		affected := db.Find(&memberSettings, "ID = ?", message.Author.ID).Limit(1).RowsAffected

		if affected > 0 {
			// Check if the member requesting this has chosen to mention only mode or not
			simpleMode = memberSettings.SimpleMode
		}
	}
	return simpleMode
}

func evaluateMentionOnlyMode(app *models.App, session *discordgo.Session, message *discordgo.Message) (bool, bool) {
	db := app.DB

	// Check if mention only mode is on in this guild
	guildSettings := dbModels.GuildSetting{}
	err := db.Where("ID = ?", message.GuildID).First(&guildSettings)

	if err.Error != nil {
		logging.PrintLog("Error getting guild settings for guild %s, error: %s", message.GuildID, err.Error.Error())
		return false, false
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
		return strings.Contains(message.Content, session.State.User.Mention()), mentionOnlyMode
	}

	return true, mentionOnlyMode
}

func evaluateDeleteMessage(db *gorm.DB, session *discordgo.Session, message *discordgo.Message) bool {
	// Check if we should delete the original message based on the guild
	guildSettings := dbModels.GuildSetting{}
	err := db.Where("ID = ?", message.GuildID).First(&guildSettings)

	if err.Error != nil {
		logging.PrintLog("Error getting guild settings for guild %s, error: %s", message.GuildID, err.Error.Error())
		return false
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

	return !keepOriginalMessage
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
