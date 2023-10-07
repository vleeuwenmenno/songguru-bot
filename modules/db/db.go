package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"songwhip_bot/models"
	dbModels "songwhip_bot/modules/db/models"
)

func SetupDB(config *models.Config) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(config.DatabasePath), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&dbModels.Guild{})
	db.AutoMigrate(&dbModels.GuildSetting{})
	db.AutoMigrate(&dbModels.MemberSetting{})

	return db, nil
}
