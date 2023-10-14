package notify

import (
	"errors"
	"time"

	"github.com/bwmarrin/discordgo"
	"gorm.io/gorm"

	"songguru_bot/models"
	dbModels "songguru_bot/modules/db/models"
	"songguru_bot/modules/logging"
)

func EnsureNotifyGuildOwner(s *discordgo.Session, app *models.App, g *discordgo.Guild, notifyID NotifyID) error {
	// Check if a notification has already been sent for this combination
	notifyRegistry := dbModels.NoticeRegistry{}
	err := app.DB.Where("notice_id = ? AND guild_id = ? AND owner_id = ?", string(notifyID), g.ID, g.OwnerID).First(&notifyRegistry).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	// If a notification has already been sent, do not send it again
	if notifyRegistry.ID != 0 {
		return nil
	}

	channel, err := s.UserChannelCreate(g.OwnerID)
	if err != nil {
		return err
	}

	_, err = s.ChannelMessageSend(channel.ID, GetNotifyMessageFromConfig(app, notifyID))
	if err != nil {
		return err
	}

	// Save the notification registry entry to the database
	notifyRegistry = dbModels.NoticeRegistry{
		NoticeID: string(notifyID),
		GuildID:  g.ID,
		OwnerID:  g.OwnerID,
		SentAt:   time.Now(),
	}
	err = app.DB.Create(&notifyRegistry).Error
	logging.PrintLog("-    : guild owner %s notified from guild %s notification reason %s", g.OwnerID, g.ID, notifyID)
	if err != nil {
		return err
	}

	return nil
}

func GetNotifyMessageFromConfig(app *models.App, notifyID NotifyID) string {
	var message = ""
	for _, option := range app.Config.NotifyMessages {
		if notifyID == NotifyID(option.ID) {
			message = option.Message
		}
	}

	if message == "" {
		message = "unknown NotifyID, can't send notification message"
	}

	return message
}
