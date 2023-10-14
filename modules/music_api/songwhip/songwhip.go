package songwhip

import (
	"fmt"

	musicapi "songguru_bot/modules/music_api"
	songwhip_api "songguru_bot/modules/songwhip_api"
	songwhip_api_models "songguru_bot/modules/songwhip_api/models"
)

type SongwhipAdapter struct {
}

func (s *SongwhipAdapter) Search(query string) ([]musicapi.Song, error) {
	result, err := s.GetSong(query)

	if err != nil {
		return nil, err
	}

	// Populate single index with result:
	songs := []musicapi.Song{*result}
	return songs, nil
}

func (s *SongwhipAdapter) GetAlbum(albumID string) (*musicapi.Album, error) {
	// Implementation for getting album details from Spotify
	return &musicapi.Album{}, nil
}

func (s *SongwhipAdapter) GetSong(songID string) (*musicapi.Song, error) {
	result, err := songwhip_api.GetInfo(songID)

	if err != nil {
		return nil, err
	}

	song := musicapi.Song{
		ID:         fmt.Sprintf("%d", result.ID),
		Title:      result.Name,
		Artists:    getArtists(result),
		URL:        result.URL,
		PictureURL: result.Image,
	}
	return &song, nil
}

func getArtists(info *songwhip_api_models.SongwhipInfo) []musicapi.Artist {
	artists := []musicapi.Artist{}
	for _, artist := range info.Artists {
		artists = append(artists, musicapi.Artist{
			ID:         fmt.Sprintf("%d", artist.ID),
			Title:      artist.Name,
			URL:        artist.URL,
			PictureURL: artist.Image,
		})
	}
	return artists
}
