package netease

import (
	"github.com/xiaomi388/virtual-music-system/pkg/common"
	"github.com/xiaomi388/virtual-music-system/pkg/metadata/artist"
)

// ArtistRepository retrieves artist metadata from netease API.
type ArtistRepository struct {
}

// GetArtist retrieves the metadata of the artist of a specific id from netease API by.
func (r *ArtistRepository) GetArtist(id artist.ID) (artist.Artist, error) {
	return artist.Artist{}, common.ErrNotImpl
}

func (r *ArtistRepository) GetArtistsByQuery(
	q string, limit int, offset int) (map[artist.ID]artist.Artist, error) {
	return nil, common.ErrNotImpl
}
