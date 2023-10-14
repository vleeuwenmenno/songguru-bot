package songwhip

import (
	"errors"
	"testing"

	"github.com/h2non/gock"

	musicapi "songguru_bot/modules/music_api"
	th "songguru_bot/testing"
)

func localSetup() {
	defer gock.Off()

	timeoutErr := errors.New("dial tcp i/o timeout")
	gock.New("https://songwhip.com").Post("/").Reply(200).SetError(timeoutErr)
}

func TestSearch(t *testing.T) {
	th.Setup(t)
	localSetup()

	var musicAPI musicapi.MusicAPI = &SongwhipAdapter{}
	songs, err := musicAPI.Search("https://open.spotify.com/track/5xrtzzzikpG3BLbo4q1Yul?si=8848d83c6fec4073")

	if err != nil {
		t.Error("Search result is nil")
	}

	if len(songs) != 1 {
		t.Errorf("Unexpected result, songs length was %d but expected 1", len(songs))
	}
}

func TestFetchSong(t *testing.T) {
	th.Setup(t)
	var musicAPI musicapi.MusicAPI = &SongwhipAdapter{}
	songID := "12345"
	song, err := musicAPI.GetSong(songID)

	if err != nil {
		t.Error("Song failed to fetch result is nil")
	}

	if song.ID != songID {
		t.Errorf("Song ID does not match, expected '%s' but got '%s'", songID, song.ID)
	}
}

func TestFetchAlbum(t *testing.T) {
	th.Setup(t)
	var musicAPI musicapi.MusicAPI = &SongwhipAdapter{}
	albumID := "12345"
	album, err := musicAPI.GetAlbum(albumID)

	if err != nil {
		t.Error("Album failed to fetch result is nil")
	}

	if album.ID != albumID {
		t.Errorf("Album ID does not match, expected '%s' but got '%s'", albumID, album.ID)
	}
}
