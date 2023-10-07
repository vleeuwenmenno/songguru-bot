package models

import (
	"time"

	"gorm.io/gorm"
)

type Guild struct {
	gorm.Model
	ID                     string
	AdminRoleID            string
	NotificationsChannelID string
	JoinedAt               time.Time
}
