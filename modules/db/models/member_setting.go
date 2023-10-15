package databaseModels

import (
	"gorm.io/gorm"
)

/*
GuildSetting

[SimpleMode] - If enabled we only return the link from Songwhip.
[MentionOnlyMode] - If enabled we only respond if the bot was mentioned.
[KeepOriginalMessage] - If enabled we keep the original message otherwise we delete it.
*/
type MemberSetting struct {
	gorm.Model
	ID                  string `gorm:"primaryKey"`
	SimpleMode          bool
	MentionOnlyMode     bool
	KeepOriginalMessage bool
	GuildRefer          string
	Guild               Guild `gorm:"foreignKey:GuildRefer"`
}
