package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"songwhip_bot/modules/db/models"
)

func SetupDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("database.sqlite"), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&models.Guild{})
	return db, nil
}
