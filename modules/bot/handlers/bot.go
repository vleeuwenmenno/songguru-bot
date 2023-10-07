package handlers

import "songwhip_bot/models"

type Bot interface {
	GetApp() *models.App
}
