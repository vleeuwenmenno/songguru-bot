package models

import (
	"time"

	"gorm.io/gorm"
)

type Guild struct {
	gorm.Model
	ID       string
	JoinedAt time.Time
}
