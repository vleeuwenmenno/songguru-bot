package models

import "gorm.io/gorm"

type App struct {
	Config   *Config
	Services *Services
	States   *States
	DB       *gorm.DB
}
