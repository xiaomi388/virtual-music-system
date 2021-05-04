package netease

import (
	"github.com/xiaomi388/virtual-music-system/metadata/exception"
	"github.com/xiaomi388/virtual-music-system/metadata/playlist"
)

type PlaylistRepository struct {
}

func (r *PlaylistRepository) GetPlaylist(id playlist.ID) (playlist.Playlist, error) {
	return playlist.Playlist{}, new(exception.NotImplementedError)
}

func (r *PlaylistRepository) GetPlaylistsByQuery(q string,
	limit int, offset int) (map[playlist.ID]playlist.Playlist, error) {
	return nil, new(exception.NotImplementedError)
}
