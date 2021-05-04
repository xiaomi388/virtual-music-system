package song

type Repository interface {
	GetSong(id ID) (Song, error)
	GetSongsByQuery(q string, limit int, offset int) (map[ID]Song, error)
}
