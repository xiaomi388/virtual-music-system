package playlist

import "github.com/xiaomi388/virtual-music-system/metadata/song"

type ID string

type Playlist struct {
	ID    ID
	Songs map[song.ID]struct{}
}
