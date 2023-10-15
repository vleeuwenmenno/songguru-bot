package api

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"golang.org/x/exp/slices"

	"songguru_bot/models"
	"songguru_bot/modules/api/helpers"
	apimodels "songguru_bot/modules/api/models"
	databaseModels "songguru_bot/modules/db/models"
	sliceshelper "songguru_bot/modules/helpers/slices"
	"songguru_bot/modules/logging"
)

type Output struct {
	Guild          apimodels.DiscordGuild        `json:"guild"`
	GuildSettings  databaseModels.GuildSetting   `json:"guildSettings"`
	MemberSettings *databaseModels.MemberSetting `json:"memberSetting"`
}

func setupRoutes(api *gin.RouterGroup, app *models.App) {
	conf := helpers.GetOAuth2Config(app)
	api.GET("/settings", func(c *gin.Context) {
		userClaims, err := helpers.GetUserClaims(c, app.Config.API.JWTSecret)
		if err != nil {
			c.Status(http.StatusUnprocessableEntity)
			return
		}

		guilds, err := helpers.GetDiscordUserGuilds(conf, &userClaims.Token)
		if err != nil {
			c.Status(http.StatusUnprocessableEntity)
			return
		}

		outputs := []Output{}
		for _, guild := range guilds {
			guildSetting := databaseModels.GuildSetting{}
			result := app.DB.Preload("Guild").Where("ID = ?", guild.ID).First(&guildSetting)

			if result.RowsAffected > 0 {
				memberSettings := databaseModels.MemberSetting{}
				app.DB.Find(&memberSettings, "ID = ? AND GuildRefer = ?", userClaims.DiscordUser.ID, guild.ID).Limit(1)

				if memberSettings.ID != "" {
					output := Output{
						Guild:          guild,
						GuildSettings:  guildSetting,
						MemberSettings: &memberSettings,
					}

					outputs = append(outputs, output)
				} else {
					output := Output{
						Guild:          guild,
						GuildSettings:  guildSetting,
						MemberSettings: nil,
					}

					outputs = append(outputs, output)
				}
			}
		}

		c.JSON(http.StatusOK, outputs)
	})

	api.DELETE("/settings", func(c *gin.Context) {
		//TODO: Implement ability to remove member setting (NOT GUILD SETTING)
	})

	api.POST("/settings", func(c *gin.Context) {
		// Access the userClaims from the context
		userClaims := c.MustGet("userClaims").(apimodels.UserClaims)

		savePreferencesRequest := &apimodels.SavePreferencesRequest{}
		err := c.ShouldBindJSON(&savePreferencesRequest)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		memberPreferences := databaseModels.MemberSetting{}
		if app.DB.Find(&memberPreferences, "ID = ?", userClaims.DiscordUser.ID).RowsAffected != 0 {
			c.JSON(http.StatusBadRequest, map[string]string{"error": " member already has preferences for the supplied guildId, to update use PATCH instead"})
			return
		}

		//TODO: Sanity check, ask discord api if it's true that the current user is in the provided guild id

		guildPreferences := databaseModels.GuildSetting{}
		if app.DB.Find(&guildPreferences, "ID = ?", savePreferencesRequest.GuildID).RowsAffected == 0 {
			c.JSON(http.StatusNotFound, map[string]string{"error": "provided guildId not found"})
			return
		}

		memberPreferences = databaseModels.MemberSetting{
			ID:                  userClaims.DiscordUser.ID,
			MentionOnlyMode:     guildPreferences.MentionOnlyMode,
			SimpleMode:          guildPreferences.SimpleMode,
			KeepOriginalMessage: guildPreferences.KeepOriginalMessage,
			GuildRefer:          savePreferencesRequest.GuildID,
		}
		app.DB.Create(&memberPreferences)

		// TODO: Store preferences for this user
		// TODO: Check if the user is a moderator role user, if so apply also any guild preferences provided.
		c.Status(http.StatusOK)
	})

	auth := api.Group("/auth")
	{
		auth.GET("/", func(c *gin.Context) {
			state := uuid.New().String()
			app.States.Memory = append(app.States.Memory, state)
			logging.PrintLog("spawned state %s for %s", state, c.Request.RemoteAddr)
			c.Redirect(http.StatusTemporaryRedirect, conf.AuthCodeURL(state))
		})

		auth.GET("/callback", func(c *gin.Context) {
			state := c.Query("state")
			if state == "" {
				c.Status(http.StatusBadRequest)
				return
			}

			// Check that the state id exists here, this would mean the login request was made from this service.
			if !slices.Contains(app.States.Memory, state) {
				c.Status(http.StatusUnprocessableEntity)
				return
			}

			// Clean up consumed state id
			app.States.Memory = sliceshelper.Remove(app.States.Memory, state)

			code := c.Query("code")
			token, err := conf.Exchange(context.Background(), code)
			if err != nil {
				c.Redirect(http.StatusTemporaryRedirect, "/api/auth")
				return
			}

			discordUser, err := helpers.GetDiscordUser(conf, token, c)
			if err != nil {
				c.String(http.StatusInternalServerError, err.Error())
				return
			}
			logging.PrintLog("consumed state %s for %s user %s (id: %s)", state, c.Request.RemoteAddr, discordUser.GlobalName, discordUser.ID)

			duration, err := time.ParseDuration(app.Config.API.SessionLifetime)
			if err != nil {
				logging.PrintLog("warning, failed to parse session_lifetime, falling back to 15 minutes")
				duration = time.Minute * 15
			}

			userClaims := apimodels.UserClaims{
				Token:       *token,
				DiscordUser: *discordUser,
				StandardClaims: jwt.StandardClaims{
					IssuedAt:  time.Now().Unix(),
					ExpiresAt: time.Now().Add(duration).Unix(),
				},
			}

			signedAccessToken, err := helpers.NewAccessToken(userClaims, app.Config.API.JWTSecret)
			if err != nil {
				log.Fatal("error creating access token")
				return
			}

			c.SetCookie("jwt", signedAccessToken, 3600, "/", app.Config.WebPortal.Domain, true, true)
			c.Redirect(http.StatusTemporaryRedirect, app.Config.WebPortal.URL)
		})

		auth.GET("/logout", func(c *gin.Context) {
			c.SetCookie("jwt", "", 0, "/", app.Config.WebPortal.Domain, true, true)
			c.Status(http.StatusOK)
		})

		auth.GET("/whoami", func(c *gin.Context) {
			userClaims, err := helpers.GetUserClaims(c, app.Config.API.JWTSecret)
			if err != nil {
				c.Status(http.StatusUnprocessableEntity)
				return
			}

			c.JSON(http.StatusOK, userClaims.DiscordUser)
		})
	}
}
