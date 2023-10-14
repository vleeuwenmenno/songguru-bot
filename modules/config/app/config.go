package config

import (
	"fmt"
	"io"
	"os"

	"github.com/bwmarrin/discordgo"
	"gopkg.in/yaml.v3"

	"songguru_bot/models"
	"songguru_bot/modules/logging"
)

var intentMap = map[string]discordgo.Intent{
	"Guilds":                      discordgo.IntentGuilds,
	"GuildMembers":                discordgo.IntentGuildMembers,
	"GuildBans":                   discordgo.IntentGuildBans,
	"GuildEmojis":                 discordgo.IntentGuildEmojis,
	"GuildIntegrations":           discordgo.IntentGuildIntegrations,
	"GuildWebhooks":               discordgo.IntentGuildWebhooks,
	"GuildInvites":                discordgo.IntentGuildInvites,
	"GuildVoiceStates":            discordgo.IntentGuildVoiceStates,
	"GuildPresences":              discordgo.IntentGuildPresences,
	"GuildMessages":               discordgo.IntentGuildMessages,
	"GuildMessageReactions":       discordgo.IntentGuildMessageReactions,
	"GuildMessageTyping":          discordgo.IntentGuildMessageTyping,
	"DirectMessages":              discordgo.IntentDirectMessages,
	"DirectMessageReactions":      discordgo.IntentDirectMessageReactions,
	"DirectMessageTyping":         discordgo.IntentDirectMessageTyping,
	"MessageContent":              discordgo.IntentMessageContent,
	"GuildScheduledEvents":        discordgo.IntentGuildScheduledEvents,
	"AutoModerationConfiguration": discordgo.IntentAutoModerationConfiguration,
	"AutoModerationExecution":     discordgo.IntentAutoModerationExecution,
}

func GetConfig(configFiles ...string) (*models.Config, error) {
	var configFile string

	if len(configFiles) > 0 {
		configFile = configFiles[0]
	} else {
		configFile = "configs/config.yaml"
	}

	file, err := os.Open(configFile)
	if err != nil {
		return nil, fmt.Errorf("failed to open config file: %w", err)
	}
	defer file.Close()

	config := &models.Config{}
	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(config)
	if err != nil && err != io.EOF {
		return nil, fmt.Errorf("failed to decode config: %w", err)
	}

	return config, nil
}

func GetIntents(config *models.Config) discordgo.Intent {
	var intents discordgo.Intent
	for _, intentName := range config.Intents {
		intent, ok := intentMap[intentName]
		if !ok {
			panic(fmt.Sprintf("Unknown intent: %s", intentName))
		}
		logging.PrintLog("Adding intent: %s", intentName)
		intents |= intent
	}

	return intents
}
