package actions

import (
	"errors"

	"github.com/bwmarrin/discordgo"

	"songguru_bot/models"
	dbModels "songguru_bot/modules/db/models"
	"songguru_bot/modules/logging"
)

// Ensures the moderator role exists in the guild
// If it doesn't exist, it will be created
//
// Returns the ID of the moderator role
func EnsureBotRoleExists(role **discordgo.Role, s *discordgo.Session, g *discordgo.Guild, app *models.App) {
	var err error
	*role, err = GetBotRole(s, g, app.Config.Discord.ModeratorRoleName)
	if err != nil {
		ClearGuildBotRoleID(g.ID, app)

		// Bot role not found, create it
		*role, err = CreateBotRole(s, g, app.Config.Discord.ModeratorRoleName)
		if err != nil {
			logging.PrintLog("-    : error creating bot role '%s' for guild %s : %s", app.Config.Discord.ModeratorRoleName, g.ID, err)
			return
		}
	}

	logging.PrintLog("-    : using bot role %s (id: %s) for guild %s", (*role).Name, (*role).ID, g.ID)
	app.DB.Model(&dbModels.Guild{}).Where("ID = ?", g.ID).Update("AdminRoleID", (*role).ID)
}

func GetBotRole(s *discordgo.Session, g *discordgo.Guild, botRoleName string) (*discordgo.Role, error) {
	for _, role := range g.Roles {
		if role.Name == botRoleName {
			return role, nil
		}
	}

	return nil, errors.New("bot role not found in guild")
}

func ClearGuildBotRoleID(guildID string, app *models.App) {
	// Find the guild in the database
	existingGuild := dbModels.Guild{}
	err := app.DB.Find(&existingGuild, "ID = ?", guildID).Limit(1).Error
	if err != nil {
		logging.PrintLog("warning, received a role created event for a unknown guild %s : %s", guildID, err)
		return
	}

	// Update the AdminRoleId column on the guild table
	existingGuild.AdminRoleID = nil
	err = app.DB.Save(&existingGuild).Error
	if err != nil {
		logging.PrintLog("failed to clear role id for guild %s : %s", guildID, err)
		return
	}
}

func CreateBotRole(s *discordgo.Session, g *discordgo.Guild, botRoleName string) (*discordgo.Role, error) {
	role, err := s.GuildRoleCreate(g.ID, &discordgo.RoleParams{
		Name: botRoleName,
	})
	if err != nil {
		return nil, err
	}

	return role, nil
}

func EnsureGuildOwnerHasBotRole(s *discordgo.Session, g *discordgo.Guild, app *models.App) {
	ownerID := g.OwnerID
	member, err := s.GuildMember(g.ID, ownerID)
	if err != nil {
		logging.PrintLog("-    : error getting guild member: %s", err)
		return
	}

	existingGuild := dbModels.Guild{}
	err = app.DB.Find(&existingGuild, "ID = ?", g.ID).Limit(1).Error
	if err != nil {
		logging.PrintLog("-    : error failed to query existing guild from db: %s", err)
		return
	}

	// Check if the owner already has the bot role
	for _, roleID := range member.Roles {
		if roleID == *existingGuild.AdminRoleID {
			logging.PrintLog("-    : bot role found for guild owner %s for guild %s", ownerID, g.ID)
			return
		}
	}

	// Assign the bot role to the owner
	err = s.GuildMemberRoleAdd(g.ID, ownerID, *existingGuild.AdminRoleID)
	if err != nil {
		logging.PrintLog("-    : error assigning bot role to guild owner: %s", err)
		return
	}

	logging.PrintLog("-    : assigned bot role to guild owner %s for guild %s", ownerID, g.ID)
}

func EnsureBotHasBotRole(s *discordgo.Session, g *discordgo.Guild, app *models.App) {
	existingGuild := dbModels.Guild{}
	err := app.DB.Find(&existingGuild, "ID = ?", g.ID).Limit(1).Error
	if err != nil {
		logging.PrintLog("-    : error failed to query existing guild from db: %s", err)
		return
	}

	botID := s.State.User.ID
	member, err := s.GuildMember(g.ID, botID)
	if err != nil {
		logging.PrintLog("-    : error getting bot id member: %s", err)
		return
	}

	// Check if the owner already has the bot role
	for _, roleID := range member.Roles {
		if roleID == *existingGuild.AdminRoleID {
			logging.PrintLog("-    : bot role found for bot %s for guild %s", botID, g.ID)
			return
		}
	}

	// Assign the bot role
	err = s.GuildMemberRoleAdd(g.ID, botID, *existingGuild.AdminRoleID)
	if err != nil {
		logging.PrintLog("-    : error assigning bot role to bot: %s", err)
		return
	}

	logging.PrintLog("-    : assigned bot role to bot %s for guild %s", botID, g.ID)
}
