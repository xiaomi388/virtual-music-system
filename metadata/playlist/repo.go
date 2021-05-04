package playlist

type Repository interface {
	GetPlaylist(id ID) (Playlist, error)
	GetPlaylistsByQuery(q string, limit int, offset int) (map[ID]Playlist, error)
}
