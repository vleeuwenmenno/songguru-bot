package actions

import (
	"songwhip_bot/models"
	dbModels "songwhip_bot/modules/db/models"
	"songwhip_bot/modules/logging"

	"github.com/bwmarrin/discordgo"
	"golang.org/x/exp/slices"
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
	app.DB.Model(&dbModels.Guild{}).Where("ID = ?", g.ID).Update("AdminRoleID", *role)
}

func EnsureAdminRoleAssigned(s *discordgo.Session, g *discordgo.Guild, app *models.App) {
	role := guildHasAdminRole(g.Roles, app.Config)
	botRoles, err := s.GuildMember(g.ID, s.State.User.ID)

	if err != nil {
		logging.PrintLog("Error getting bot roles")
	}

	if !slices.Contains(botRoles.Roles, *role) {
		if s.GuildMemberRoleAdd(g.ID, s.State.User.ID, *role) != nil {
			logging.PrintLog("Error assigning admin role to bot")
		}

		logging.PrintLog("Bot got admin role assigned for guild %s (ID: %s) (RoleID: %s)", g.Name, g.ID, *role)
	}
}

// Adds the admin role if it doesn't exist in the guild
//
// Returns the ID of the admin role
func conditionallyAddAdminRole(s *discordgo.Session, g *discordgo.Guild, config *models.Config) (*string, error) {
	roles, err := s.GuildRoles(g.ID)

	if err != nil {
		logging.PrintLog("Error getting guild roles")
		return nil, err
	}

	role := guildHasAdminRole(roles, config)

	// Check discord can tell us if guild has a role called SongwhipAdmin
	if role == nil {
		role, err := s.GuildRoleCreate(g.ID, &discordgo.RoleParams{
			Name: config.Discord.AdminRoleName,
		})

		if err != nil {
			logging.PrintLog("Error adding admin role")
			return nil, err
		}

		logging.PrintLog("	- added admin role %s for guild %s (ID: %s)", role.Name, g.Name, g.ID)
	}

	return role, nil
}

func guildHasAdminRole(roles []*discordgo.Role, config *models.Config) *string {
	for _, role := range roles {
		if role.Name == config.Discord.AdminRoleName {
			return &role.ID
		}
	}
	return nil
}
