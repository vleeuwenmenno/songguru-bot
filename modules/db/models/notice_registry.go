package databaseModels

import (
	"time"

	"gorm.io/gorm"
)

type NoticeRegistry struct {
	gorm.Model
	ID       uint      // Auto-incrementing ID
	GuildID  string    // Guild ID
	OwnerID  string    // Owner ID
	NoticeID string    // Notice ID
	SentAt   time.Time // Timestamp of when the notice was sent
}
