package processing

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/bwmarrin/discordgo"
	"gorm.io/gorm"

	"songguru_bot/models"
	dbModels "songguru_bot/modules/db/models"
	"songguru_bot/modules/logging"
	songwhip_api "songguru_bot/modules/songwhip_api"
	songguruModels "songguru_bot/modules/songwhip_api/models"
)

func ProcessMessage(app *models.App, session *discordgo.Session, event *discordgo.MessageCreate) {
	mayContinue, mentionOnlyMode := evaluateMentionOnlyMode(app, session, event.Message)
	shouldDeleteMessage := evaluateDeleteMessage(app.DB, session, event.Message)
	simpleMode := evaluateSimpleMode(app, session, event.Message)

	link, err := extractLink(event.Content)
	linkValid := err == nil && link != nil

	logging.PrintLog("message user %s in guild %s, link valid? %v, delete? %v, simple mode? %v, mention only mode? %v, may continue? %v", event.Message.Author.ID, event.Message.GuildID, linkValid, shouldDeleteMessage, simpleMode, mentionOnlyMode, mayContinue)

	if err != nil {
		return
	}

	if !mayContinue {
		return
	}

	for _, service := range app.Services.StreamingServices {
		if isValidLink(*link, service) {
			info, err := songwhip_api.GetInfo(*link)

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

			embed := buildEmbed(info, shouldDeleteMessage, event, link, session, service)
			session.ChannelMessageSendEmbed(event.ChannelID, embed)
			return
		}
	}
}

func buildEmbed(info *songguruModels.SongwhipInfo, shouldDeleteMessage bool, event *discordgo.MessageCreate, link *string, session *discordgo.Session, service models.Service) *discordgo.MessageEmbed {
	newArtists := []string{}
	for _, artist := range info.Artists {
		newArtists = append(newArtists, artist.Name)
	}

	artists := strings.Join(newArtists, ", ")
	desc := fmt.Sprintf("**Track:** %s\n**Artist:** %s\n**Stream it from** %s", info.Name, artists, buildStreamingServices(info))
	displayName := event.Message.Member.Nick

	if displayName == "" {
		displayName = event.Message.Author.Username
	}

	if shouldDeleteMessage {
		contents := event.Message.Content

		contents = strings.Replace(contents, *link, "", -1)
		contents = strings.Replace(contents, session.State.User.Mention(), "", -1)

		if strings.Trim(contents, " ") != "" {
			desc += "\n\n**" + displayName + " says** " + contents + "\n"
		}
	}

	embed := &discordgo.MessageEmbed{
		Title: info.Name + " - " + artists,
		URL:   info.URL,
		Image: &discordgo.MessageEmbedImage{
			URL: info.Image,
		},
		Description: desc,
		Color:       service.Color,
		Footer: &discordgo.MessageEmbedFooter{
			Text:    "Shared by " + displayName,
			IconURL: event.Author.AvatarURL(""),
		},
	}
	return embed
}

func buildStreamingServices(info *songguruModels.SongwhipInfo) string {
	desc := ""

	if info.Links.Spotify {
		desc += "<:spotify:860992370954469407> "
	}

	if info.Links.Deezer {
		desc += "<:deezer:860992333914570772> "
	}

	if info.Links.Itunes {
		desc += "<:applemusic:860995200797507604> "
	}

	if info.Links.YoutubeMusic {
		desc += "<:youtubemusic:860994648888836118> "
	}

	if info.Links.Youtube {
		desc += "<:youtube:860992285483335730> "
	}

	if info.Links.Pandora {
		desc += "<:pandora:860992558519418910> "
	}

	if info.Links.Tidal {
		desc += "<:tidal:860992188434612245> "
	}

	return desc
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
