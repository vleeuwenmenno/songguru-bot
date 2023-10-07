package handlers

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func ready(s *discordgo.Session, r *discordgo.Ready) {
	log.Printf("Logged in as: %v#%v", s.State.User.Username, s.State.User.Discriminator)
	log.Printf("Bot is serving in %d servers!", len(s.State.Guilds))

	for _, guild := range s.State.Guilds {
		g, err := s.Guild(guild.ID)
		if err != nil {
			log.Printf("Failed to get guild information: %v", err)
			continue
		}

		log.Printf(" - %s (id: %s)", g.Name, g.ID)
	}
}
