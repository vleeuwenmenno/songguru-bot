package services

import (
	th "songwhip_bot/testing"
	"testing"
)

func TestGetConfig(t *testing.T) {
	th.Setup(t)

	// Specify the path to the test-specific YAML file
	testConfigFile := "configs/services.test.yaml"

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
	if len(services.StreamingServices) != 2 {
		t.Fatalf("expected 2 services, but got %d", len(services.StreamingServices))
	}

	// Check the first service
	serviceA, ok := services.StreamingServices["serviceA"]
	if !ok {
		t.Fatalf("expected to find service 'serviceA', but it was not found")
	}

	if serviceA.Name != "ServiceA" {
		t.Errorf("expected serviceA service name to be 'ServiceA', but got %q", serviceA.Name)
	}

	if serviceA.Icon != "<:serviceA:12345678901234567>" {
		t.Errorf("expected serviceA service icon to be '<:serviceA:12345678901234567>', but got %q", serviceA.Icon)
	}

	if serviceA.Color != "#ff0000" {
		t.Errorf("expected serviceA service color to be '#ff0000', but got %q", serviceA.Color)
	}

	expectedUrls := []string{"https://exampleA.link", "https://exampleA.com"}
	if !th.Equal(serviceA.Urls, expectedUrls) {
		t.Errorf("expected serviceA service URLs to be %v, but got %v", expectedUrls, serviceA.Urls)
	}
}
