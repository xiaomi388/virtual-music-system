package model

type RemoteRepository interface {
	GetSongByNameArtist(name string, artist string) (filePath string, err error)
}

type LocalRepository interface {
	GetSongByNameArtist(name string, artist string) (Song, error)
	CopySongFrom(filePath string, name string, artist string) (Song, error)
}
