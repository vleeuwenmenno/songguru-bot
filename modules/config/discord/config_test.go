package config

import (
	th "songwhip_bot/testing"
	"strings"
	"testing"
)

func TestGetConfig(t *testing.T) {
	th.Setup(t)

	// Specify the path to the test-specific YAML file
	testConfigFile := "configs/config.test.yaml"

	// Load the configuration
	config, err := GetConfig(testConfigFile)
	if err != nil {
		t.Fatalf("failed to load config: %v", err)
	}

	// Assert the value of BotToken
	expectedBotToken := "MOCK_BOT_TOKEN"
	if config.Discord.BotToken != expectedBotToken {
		t.Errorf("unexpected bot token: got %s, want %s", config.Discord.BotToken, expectedBotToken)
	}

	// Assert that we have the correct number of intents
	if len(config.Intents) != 2 {
		t.Errorf("unexpected intents: got %d, want 2", len(config.Intents))
	}

	// Assert that we have the correct intents
	expectedIntents := []string{"Guilds", "GuildMessages"}
	if !th.Equal(config.Intents, expectedIntents) {
		t.Errorf("expected intents to be %v, but got %v", expectedIntents, config.Intents)
	}
}

func TestGetConfigFailure(t *testing.T) {
	th.Setup(t)

	// Specify the path to the test-specific YAML file
	testConfigFile := "blabla/config.test.yaml"

	// Load the configuration
	_, err := GetConfig(testConfigFile)
	if err == nil {
		t.Errorf("expected error, but got nil")
	}

	expectedErrorMessage := "failed to open config file: open blabla/config.test.yaml"
	if !strings.Contains(err.Error(), expectedErrorMessage) {
		t.Errorf("expected error message to contain %q, but got %q", expectedErrorMessage, err.Error())
	}
}
