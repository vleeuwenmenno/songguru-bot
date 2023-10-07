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
