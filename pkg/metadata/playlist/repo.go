package playlist

import "github.com/xiaomi388/virtual-music-system/pkg/metadata/song"

// Repository contains CRUD interfaces for Playlist entities.
type Repository interface {
	GetPlaylist(id song.ID) (Playlist, error)
	GetPlaylistsByQuery(q string, limit int, offset int) (map[song.ID]Playlist, int, error)
}
