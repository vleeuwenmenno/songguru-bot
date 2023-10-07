package commands

import (
	"errors"
	"songwhip_bot/models"

	"github.com/bwmarrin/discordgo"
)

type CommandFunc func(session *discordgo.Session, event *discordgo.InteractionCreate, app *models.App) error

var commands = map[string]CommandFunc{
	"settings":  settings,
	"changelog": changelog,
}

func EvaluateCommand(command string) CommandFunc {
	val, ok := commands[command]

	if ok {
		return val
	}

	return func(session *discordgo.Session, event *discordgo.InteractionCreate, app *models.App) error {
		return errors.New("command not found")
	}
}
