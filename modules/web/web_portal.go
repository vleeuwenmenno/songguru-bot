package web

import (
	"fmt"
	"net/http"
	"songwhip_bot/models"
	dbModels "songwhip_bot/modules/db/models"
	"time"

	"github.com/gin-gonic/gin"
)

type MemberSettings struct {
	SimpleMode          bool `json:"SimpleMode"`
	MentionOnlyMode     bool `json:"MentionOnlyMode"`
	KeepOriginalMessage bool `json:"KeepOriginalMessage"`
}

type ServerSettings struct {
	SimpleMode                       bool `json:"SimpleMode"`
	MentionOnlyMode                  bool `json:"MentionOnlyMode"`
	KeepOriginalMessage              bool `json:"KeepOriginalMessage"`
	AllowOverrideSimpleMode          bool `json:"AllowOverridesimpleMode"`
	AllowOverrideMentionOnlyMode     bool `json:"AllowOverridementionOnlyMode"`
	AllowOverrideKeepOriginalMessage bool `json:"AllowOverridekeepOriginalMessage"`
}

type SaveSettingsRequest struct {
	MemberSettings MemberSettings `json:"memberSettings"`
	ServerSettings ServerSettings `json:"serverSettings"`
}

func StartWebPortal(router *gin.Engine, app *models.App) {
	router.LoadHTMLGlob("www/templates/*")

	router.GET("css/pico.min.css", func(c *gin.Context) {
		c.File("www/css/pico.min.css")
	})

	router.POST("/", func(c *gin.Context) {
		token := c.Query("t")
		webToken := dbModels.SettingsWebToken{}
		rows := app.DB.Find(&webToken, "ID = ?", token).Limit(1).RowsAffected

		if rows == 0 || webToken.ExpiresAt.Before(time.Now()) {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		saveSettingsRequest := &SaveSettingsRequest{}
		err := c.ShouldBindJSON(&saveSettingsRequest)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		serverSettings := dbModels.GuildSetting{}
		if app.DB.Find(&serverSettings, "ID = ?", webToken.GuildID).RowsAffected == 0 {
			c.HTML(http.StatusInternalServerError, "error.html", nil)
			return
		}

		memberSettings := dbModels.MemberSetting{}
		if app.DB.Find(&memberSettings, "ID = ?", webToken.DiscordUserID).RowsAffected == 0 {
			memberSettings := dbModels.MemberSetting{
				ID:                  webToken.DiscordUserID,
				MentionOnlyMode:     serverSettings.MentionOnlyMode,
				SimpleMode:          serverSettings.SimpleMode,
				KeepOriginalMessage: serverSettings.KeepOriginalMessage,
			}
			app.DB.Create(&memberSettings)
		}

		app.DB.Model(&dbModels.MemberSetting{}).Where("ID = ?", memberSettings.ID).Update("SimpleMode", saveSettingsRequest.MemberSettings.SimpleMode)
		app.DB.Model(&dbModels.MemberSetting{}).Where("ID = ?", memberSettings.ID).Update("MentionOnlyMode", saveSettingsRequest.MemberSettings.MentionOnlyMode)
		app.DB.Model(&dbModels.MemberSetting{}).Where("ID = ?", memberSettings.ID).Update("KeepOriginalMessage", saveSettingsRequest.MemberSettings.KeepOriginalMessage)

		if webToken.AdminToken {
			app.DB.Model(&dbModels.GuildSetting{}).Where("ID = ?", webToken.GuildID).Update("SimpleMode", saveSettingsRequest.ServerSettings.SimpleMode)
			app.DB.Model(&dbModels.GuildSetting{}).Where("ID = ?", webToken.GuildID).Update("MentionOnlyMode", saveSettingsRequest.ServerSettings.MentionOnlyMode)
			app.DB.Model(&dbModels.GuildSetting{}).Where("ID = ?", webToken.GuildID).Update("KeepOriginalMessage", saveSettingsRequest.ServerSettings.KeepOriginalMessage)
			app.DB.Model(&dbModels.GuildSetting{}).Where("ID = ?", webToken.GuildID).Update("AllowOverrideSimpleMode", saveSettingsRequest.ServerSettings.AllowOverrideSimpleMode)
			app.DB.Model(&dbModels.GuildSetting{}).Where("ID = ?", webToken.GuildID).Update("AllowOverrideMentionOnlyMode", saveSettingsRequest.ServerSettings.AllowOverrideMentionOnlyMode)
			app.DB.Model(&dbModels.GuildSetting{}).Where("ID = ?", webToken.GuildID).Update("AllowOverrideKeepOriginalMessage", saveSettingsRequest.ServerSettings.AllowOverrideKeepOriginalMessage)
		}

		c.JSON(http.StatusAccepted, gin.H{"status": "success"})
	})

	router.GET("/", func(c *gin.Context) {
		indexRouter(c, app, false)
	})

	router.GET("/heartbeat", func(c *gin.Context) {
		c.AbortWithStatus(http.StatusOK)
	})

	router.GET("/changelogs", func(c *gin.Context) {
		c.HTML(http.StatusOK, "changelogs.html", gin.H{
			"title": "Change Log",
		})
	})

	router.Run(fmt.Sprintf(":%d", app.Config.WebPortal.Port))
}

func indexRouter(c *gin.Context, app *models.App, modified bool) {
	token := c.Query("t")
	webToken := dbModels.SettingsWebToken{}
	rows := app.DB.Find(&webToken, "ID = ?", token).Limit(1).RowsAffected

	if rows == 0 || webToken.ExpiresAt.Before(time.Now()) {
		c.HTML(http.StatusUnauthorized, "error.html", nil)
		return
	}

	serverSettings := dbModels.GuildSetting{}
	if app.DB.Find(&serverSettings, "ID = ?", webToken.GuildID).RowsAffected == 0 {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	memberSettings := dbModels.MemberSetting{}
	if app.DB.Find(&memberSettings, "ID = ?", webToken.DiscordUserID).RowsAffected == 0 {
		memberSettings := dbModels.MemberSetting{
			ID:                  webToken.DiscordUserID,
			MentionOnlyMode:     serverSettings.MentionOnlyMode,
			SimpleMode:          serverSettings.SimpleMode,
			KeepOriginalMessage: serverSettings.KeepOriginalMessage,
		}
		app.DB.Create(&memberSettings)
	}

	c.HTML(http.StatusOK, "index.html", gin.H{
		"title":          "Settings",
		"token":          token,
		"memberSettings": memberSettings,
		"serverSettings": serverSettings,
		"adminToken":     webToken.AdminToken,
		"modified":       modified,
	})
}
