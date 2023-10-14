package musicapi

type MusicAPI interface {
	Search(query string) ([]Song, error)
	GetAlbum(albumID string) (*Album, error)
	GetSong(songID string) (*Song, error)
}
