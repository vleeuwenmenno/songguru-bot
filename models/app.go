package models

import "gorm.io/gorm"

type App struct {
	Config   *Config
	Services *Services
	DB       *gorm.DB
}
