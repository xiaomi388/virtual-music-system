// Package album contains domain models relating to albums.
package album

import "github.com/xiaomi388/virtual-music-system/pkg/metadata/song"

// ID is the type of Album.ID.
type ID string

// Album contains metadata for a album, and the song ids of an album.
type Album struct {
	ID    ID
	Name  string
	Songs map[song.ID]struct{}
}
