package actions

import (
	"songwhip_bot/modules/logging"

	"github.com/bwmarrin/discordgo"
)

// Ensures the admin role exists in the guild
// If it doesn't exist, it will be created and assigned to the bot
//
// Returns the ID of the admin role
func EnsureAdminRole(s *discordgo.Session, g *discordgo.Guild) (*string, error) {
	roles, err := s.GuildRoles(g.ID)

	if err != nil {
		logging.PrintLog("Error getting guild roles")
		return nil, err
	}

	role := GuildHasAdminRole(roles)

	// Check discord can tell us if guild has a role called SongwhipAdmin
	if role == nil {
		role, err := s.GuildRoleCreate(g.ID, &discordgo.RoleParams{
			Name: "SongwhipAdmin",
		})

		if err != nil {
			logging.PrintLog("Error adding admin role")
			return nil, err
		}

		if s.GuildMemberRoleAdd(g.ID, s.State.User.ID, role.ID) != nil {
			logging.PrintLog("Error assigning admin role to bot")
			return nil, err
		}
		return &role.ID, nil
	}
	return role, nil
}

func GuildHasAdminRole(roles []*discordgo.Role) *string {
	for _, role := range roles {
		if role.Name == "SongwhipAdmin" {
			return &role.ID
		}
	}
	return nil
}
