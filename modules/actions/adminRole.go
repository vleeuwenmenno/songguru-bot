package actions

import (
	"songguru_bot/models"
	dbModels "songguru_bot/modules/db/models"
	"songguru_bot/modules/logging"

	"github.com/bwmarrin/discordgo"
)

// Ensures the admin role exists in the guild
// If it doesn't exist, it will be created and assigned to the bot
//
// Returns the ID of the admin role
func EnsureAdminRoleExists(s *discordgo.Session, g *discordgo.Guild, app *models.App) {
	role, err := conditionallyAddAdminRole(s, g, app.Config)
	if err != nil {
		logging.PrintLog("Error adding/getting admin role for/from guild: %s", err)
		return
	}
	app.DB.Model(&dbModels.Guild{}).Where("ID = ?", g.ID).Update("AdminRoleID", role)
}

func EnsureAdminRoleAssigned(s *discordgo.Session, g *discordgo.Guild, app *models.App) {
	role, err := GuildHasAdminRole(g.Roles, app.Config)

	if err != nil {
		logging.PrintLog("Error getting admin role for guild %s (ID: %s)", g.Name, g.ID)
		return
	}

	botGuildMember, err := s.GuildMember(g.ID, s.State.User.ID)

	if err != nil {
		logging.PrintLog("Error getting bot roles for guild %s (ID: %s)", g.Name, g.ID)
		return
	}

	botHasAdminRole := false
	for _, botRole := range botGuildMember.Roles {
		if botRole == role {
			botHasAdminRole = true
			break
		}
	}

	if !botHasAdminRole {
		if s.GuildMemberRoleAdd(g.ID, s.State.User.ID, role) != nil {
			logging.PrintLog("Error assigning admin role to bot for guild %s (ID: %s)", g.Name, g.ID)
			return
		}

		logging.PrintLog("Bot got admin role assigned for guild %s (ID: %s) (RoleID: %s)", g.Name, g.ID, role)
	}
}

// Adds the admin role if it doesn't exist in the guild
//
// Returns the ID of the admin role
func conditionallyAddAdminRole(s *discordgo.Session, g *discordgo.Guild, config *models.Config) (string, error) {
	roles, err := s.GuildRoles(g.ID)

	if err != nil {
		logging.PrintLog("Error getting bot roles for guild %s (ID: %s)", g.Name, g.ID)
		return "", err
	}

	role, err := GuildHasAdminRole(roles, config)

	// Check discord can tell us if guild has a role called SongGuruAdmin
	if err != nil {
		role, err := s.GuildRoleCreate(g.ID, &discordgo.RoleParams{
			Name: config.Discord.AdminRoleName,
		})

		if err != nil {
			logging.PrintLog("Error adding admin role for guild %s (ID: %s)", g.Name, g.ID)
			return "", err
		}

		logging.PrintLog("	- added admin role %s for guild %s (ID: %s)", role.Name, g.Name, g.ID)
	}

	return role, nil
}

func GuildHasAdminRole(roles []*discordgo.Role, config *models.Config) (string, error) {
	for _, role := range roles {
		if role.Name == config.Discord.AdminRoleName {
			return role.ID, nil
		}
	}
	return "", nil
}
