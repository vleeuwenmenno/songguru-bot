package handlers

import "songguru_bot/models"

type Bot interface {
	GetApp() *models.App
}
