package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"songguru_bot/models"
	dbModels "songguru_bot/modules/db/models"
)

func SetupDB(config *models.Config) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(config.DatabasePath), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&dbModels.Guild{})
	db.AutoMigrate(&dbModels.GuildSetting{})
	db.AutoMigrate(&dbModels.MemberSetting{})
	db.AutoMigrate(&dbModels.SettingsWebToken{})
	db.AutoMigrate(&dbModels.NoticeRegistry{})

	return db, nil
}
