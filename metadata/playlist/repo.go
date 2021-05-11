package playlist

import "github.com/xiaomi388/virtual-music-system/metadata/song"

type Repository interface {
	GetPlaylist(id song.ID) (PlayList, error)
	GetPlaylistsByQuery(q string, limit int, offset int) (map[song.ID]PlayList, int, error)
}
