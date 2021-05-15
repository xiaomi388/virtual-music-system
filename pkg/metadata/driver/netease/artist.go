package netease

import (
	"github.com/xiaomi388/virtual-music-system/pkg/metadata/artist"
	"github.com/xiaomi388/virtual-music-system/pkg/metadata/exception"
)

type ArtistRepository struct {
}

func (r *ArtistRepository) GetArtist(id artist.ID) (artist.Artist, error) {
	return artist.Artist{}, new(exception.NotImplementedError)
}

func (r *ArtistRepository) GetArtistsByQuery(
	q string, limit int, offset int) (map[artist.ID]artist.Artist, error) {
	return nil, new(exception.NotImplementedError)
}
