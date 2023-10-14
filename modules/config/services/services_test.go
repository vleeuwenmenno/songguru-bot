package services

import (
	"testing"

	iohelper "songguru_bot/modules/helpers"
	th "songguru_bot/testing"
)

func TestGetConfig(t *testing.T) {
	th.Setup(t)

	// Specify the path to the test-specific YAML file
	testConfigFile := iohelper.ProjectRoot + "/configs/services.test.yaml"

	// Load the configuration
	services, err := GetStreamingServices(testConfigFile)
	if err != nil {
		t.Fatalf("failed to load config: %v", err)
	}

	// Check that services are loaded
	if len(services.StreamingServices) == 0 {
		t.Fatalf("no services were loaded")
	}

	// Check that two services are loaded
	if len(services.StreamingServices) != 5 {
		t.Fatalf("expected 2 services, but got %d", len(services.StreamingServices))
	}

	// Check the first service
	spotify, ok := services.StreamingServices["spotify"]
	if !ok {
		t.Fatalf("expected to find service 'spotify', but it was not found")
	}

	if spotify.Name != "Spotify" {
		t.Errorf("expected spotify service name to be 'Spotify', but got %q", spotify.Name)
	}

	if spotify.Icon != "<:spotify:860992370954469407>" {
		t.Errorf("expected spotify service icon to be '<:spotify:860992370954469407>', but got %q", spotify.Icon)
	}

	if spotify.Color != 0x1DB954 {
		t.Errorf("expected spotify service color to be '#ff0000', but got %q", spotify.Color)
	}

	expectedUrls := []string{"https://spotify.link", "https://spotify.com", "https://open.spotify.com", "https://www.spotify.com"}
	if !th.Equal(spotify.Urls, expectedUrls) {
		t.Errorf("expected spotify service URLs to be %v, but got %v", expectedUrls, spotify.Urls)
	}
}
