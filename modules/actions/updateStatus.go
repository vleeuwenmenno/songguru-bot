package actions

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func UpdateWatchStatus(s *discordgo.Session) {
	count := len(s.State.Guilds)
	s.UpdateWatchStatus(0, fmt.Sprintf("for music links in %d guilds", count))
}
