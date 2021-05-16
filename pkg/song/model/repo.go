package model

// RemoteRepository retrieves songs content from a remote source
type RemoteRepository interface {
	GetSongByNameArtist(name string, artist string) (filePath string, err error)
}

// LocalRepository stores Song entities in the local environment
// for quickly accessing by clients
type LocalRepository interface {
	GetSongByNameArtist(name string, artist string) (Song, error)
	CopySongFrom(filePath string, name string, artist string) (Song, error)
}
