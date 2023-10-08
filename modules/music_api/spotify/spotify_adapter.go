package spotify

import musicapi "songguru_bot/modules/music_api"

type SpotifyAdapter struct {
}

func (s *SpotifyAdapter) Search(query string) ([]musicapi.Song, error) {
	return nil, nil
}

func (s *SpotifyAdapter) GetSong(songID string) (*musicapi.Song, error) {
	return nil, nil
}

func (s *SpotifyAdapter) GetAlbum(albumID string) (*musicapi.Album, error) {
	return nil, nil
}
