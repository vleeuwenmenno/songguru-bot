package songwhip

import (
	"errors"
	"regexp"
	"songwhip_bot/models"
	"songwhip_bot/modules/logging"
	songwhipapi "songwhip_bot/modules/songwhip_api"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func ProcessMessage(app *models.App, s *discordgo.Session, m *discordgo.MessageCreate) {
	link, err := ExtractLink(m.Content)
	if err != nil {
		logging.PrintLog(err.Error(), err)
		return
	}

	for _, service := range app.Services.StreamingServices {
		if IsValidLink(*link, service) {
			info, err := songwhipapi.GetInfo(*link)

			if err != nil {
				logging.PrintLog(err.Error(), err)
				return
			}

			s.ChannelMessageSend(m.ChannelID, info.URL)
		}
	}
}

// Checks if a given link is one that we have a service for configured
func IsValidLink(link string, service models.Service) bool {
	for _, url := range service.Urls {
		if strings.HasPrefix(link, url) {
			return true
		}
	}
	return false
}

func ExtractLink(messageContent string) (*string, error) {
	regex := regexp.MustCompile(`(http|ftp|https):\/\/([\w\-_]+(?:(?:\.[\w\-_]+)+))([\w\-\.,@?^=%&:/~\+#]*[\w\-\@?^=%&/~\+#])?`)
	url := regex.FindString(messageContent)

	if url == "" {
		return nil, errors.New("invalid link")
	}

	return &url, nil
}
