package databaseModels

import (
	"time"

	"gorm.io/gorm"
)

/*
SettingsWebToken
*/
type SettingsWebToken struct {
	gorm.Model
	ID            string `gorm:"primaryKey"`
	GuildID       string
	DiscordUserID string
	AdminToken    bool
	ExpiresAt     time.Time
}
