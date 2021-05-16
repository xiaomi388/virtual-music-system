package netease

import (
	"github.com/xiaomi388/virtual-music-system/pkg/common"
	"github.com/xiaomi388/virtual-music-system/pkg/metadata/album"
)

type AlbumRepository struct {
}

func (r *AlbumRepository) GetAlbum(id album.ID) (album.Album, error) {
	return album.Album{}, common.ErrNotImpl
}

func (r *AlbumRepository) GetAlbumsByQuery(q string, limit int, offset int) (map[album.ID]album.Album, error) {
	return nil, common.ErrNotImpl
}
