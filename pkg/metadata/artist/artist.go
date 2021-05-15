package artist

import "github.com/xiaomi388/virtual-music-system/pkg/metadata/song"

type ID string

type Artist struct {
	ID    ID
	Songs map[song.ID]struct{}
}
