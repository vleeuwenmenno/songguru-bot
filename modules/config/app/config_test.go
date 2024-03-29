package config

import (
	"fmt"
	"strings"
	"testing"

	iohelper "songguru_bot/modules/helpers"
	th "songguru_bot/testing"
)

func TestGetConfig(t *testing.T) {
	th.Setup(t)

	// Specify the path to the test-specific YAML file
	testConfigFile := iohelper.ProjectRoot + "/configs/config.test.yaml"

	// Load the configuration
	config, err := GetConfig(testConfigFile)
	if err != nil {
		t.Fatalf("failed to load config: %v", err)
	}

	// Assert the value of BotToken
	expectedBotToken := "TOKEN_HERE"
	if config.Discord.BotToken != expectedBotToken {
		t.Errorf("unexpected bot token: got %s, want %s", config.Discord.BotToken, expectedBotToken)
	}

	// Assert that we have the correct number of intents
	if len(config.Intents) != 3 {
		t.Errorf("unexpected intents: got %d, want 3", len(config.Intents))
	}

	// Assert that we have the correct intents
	expectedIntents := []string{"Guilds", "GuildMessages", "GuildMembers"}
	if !th.Equal(config.Intents, expectedIntents) {
		t.Errorf("expected intents to be %v, but got %v", expectedIntents, config.Intents)
	}
}

func TestGetConfigFailure(t *testing.T) {
	th.Setup(t)

	// Specify the path to the test-specific YAML file
	testConfigFile := iohelper.ProjectRoot + "/blabla/config.test.yaml"

	// Load the configuration
	_, err := GetConfig(testConfigFile)
	if err == nil {
		t.Errorf("expected error, but got nil")
	}

	expectedErrorMessage := fmt.Sprintf("failed to open config file: open %s/blabla/config.test.yaml", iohelper.ProjectRoot)
	if !strings.Contains(err.Error(), expectedErrorMessage) {
		t.Errorf("expected error message to contain %q, but got %q", expectedErrorMessage, err.Error())
	}
}
