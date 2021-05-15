package netease

import (
	"github.com/xiaomi388/virtual-music-system/pkg/metadata/album"
	"github.com/xiaomi388/virtual-music-system/pkg/metadata/exception"
)

type AlbumRepository struct {
}

func (r *AlbumRepository) GetAlbum(id album.ID) (album.Album, error) {
	return album.Album{}, new(exception.NotImplementedError)
}

func (r *AlbumRepository) GetAlbumsByQuery(q string, limit int, offset int) (map[album.ID]album.Album, error) {
	return nil, new(exception.NotImplementedError)
}
