package handlers

import (
	"github.com/bwmarrin/discordgo"

	"songguru_bot/modules/actions"
	"songguru_bot/modules/logging"
)

func NewReadyHandler(b Bot) func(s *discordgo.Session, r *discordgo.Ready) {
	return func(s *discordgo.Session, r *discordgo.Ready) {
		logging.PrintLog("Logged in as: %v#%v", s.State.User.Username, s.State.User.Discriminator)
		logging.PrintLog("Bot is serving in %d servers!", len(s.State.Guilds))

		commands := []*discordgo.ApplicationCommand{
			{
				Name:        "settings",
				Description: "Request a settings edit link",
			},
			{
				Name:        "changelog",
				Description: "View the changelog of the bot",
			},
		}
	
		logging.PrintLog("Registering commands...")
		registeredCommands := make([]*discordgo.ApplicationCommand, len(commands))
		for i, v := range commands {
			cmd, err := s.ApplicationCommandCreate(s.State.User.ID, "", v)
			if err != nil {
				logging.PrintLog("Failed to register command '%v'. Err: %v", v.Name, err)
			}
	
			logging.PrintLog("- registered command '%v'", v.Name)
			registeredCommands[i] = cmd
		}

		actions.GuildsCleanup(b.GetApp(), s.State.Guilds)
	}
}
