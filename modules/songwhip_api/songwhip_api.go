package songguruapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"songguru_bot/modules/songwhip_api/models"
	"strings"
)

func GetInfo(link string) (*models.SongwhipInfo, error) {
	client := &http.Client{}

	data := map[string]string{
		"url": link,
	}
	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", "https://songwhip.com/", bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("request failed with status: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if strings.Contains(string(body), `"message": "[ApiError] 400"`) {
		return nil, fmt.Errorf("expected Name to be populated, but it was empty")
	}

	var info models.SongwhipInfo
	err = json.Unmarshal(body, &info)
	if err != nil {
		return nil, err
	}

	return &info, nil
}
