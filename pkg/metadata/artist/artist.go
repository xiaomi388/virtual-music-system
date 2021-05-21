package artist

import "github.com/xiaomi388/virtual-music-system/pkg/metadata/song"

// ID is the type of Artist.ID
type ID string

// Artist represents an artist entity containing the basic info of
// this artist and the ids of songs she created.
type Artist struct {
	ID    ID
	Songs map[song.ID]struct{}
}
