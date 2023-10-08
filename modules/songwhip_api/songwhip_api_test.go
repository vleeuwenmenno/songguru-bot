package songguruapi

import (
	"errors"
	th "songguru_bot/testing"
	"strings"
	"testing"

	"github.com/h2non/gock"
)

func TestGetInfo(t *testing.T) {
	th.Setup(t)
	defer gock.Off()
	jsonFile := th.LoadJSON("modules/songguru_api/data/success.json")

	gock.New("https://songguru.com").
		Post("").
		Reply(200).
		JSON(jsonFile)

	link := "https://open.spotify.com/track/3qXXI66oCX4veelxLACqZ3?si=4cca301a271d4550"
	info, err := GetInfo(link)

	if err != nil {
		t.Errorf("Error occurred while getting Songwhip info: %v", err)
		return
	}

	// Perform assertions on the retrieved Songwhip info
	// For example, you can check if certain fields are populated correctly
	if info.Name == "" {
		t.Errorf("Expected Name to be populated, but it was empty")
	}
}

func TestGetInfo_MalformedResponse(t *testing.T) {
	th.Setup(t)
	defer gock.Off()
	jsonFile := th.LoadJSON("modules/songguru_api/data/malformed.json")

	gock.New("https://songguru.com").
		Post("").
		Reply(200).
		JSON(jsonFile)

	link := "https://open.spotify.com/track/3qXXI66oCX4veelxLACqZ3?si=4cca301a271d4550"
	_, err := GetInfo(link)

	if err == nil {
		t.Errorf("Expected malformed response to throw an error, but got %v", err)
		return
	}
}

func TestGetInfo_NetworkError(t *testing.T) {
	th.Setup(t)
	defer gock.Off()

	timeoutErr := errors.New("dial tcp i/o timeout")
	gock.New("https://songguru.com").Post("/").Reply(200).SetError(timeoutErr)

	// Simulate a network error by not setting up a response
	link := "https://open.spotify.com/track/3qXXI66oCX4veelxLACqZ3?si=4cca301a271d4550"
	_, err := GetInfo(link)

	if err == nil {
		t.Errorf("Expected an error, but got nil")
		return
	}

	if !strings.Contains(err.Error(), timeoutErr.Error()) {
		t.Errorf("Expected error message %v, but got %v", timeoutErr, err)
	}
}
