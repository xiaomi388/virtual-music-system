package album

import "github.com/xiaomi388/virtual-music-system/metadata/song"

type ID string

type Album struct {
	ID    ID
	Name  string
	Songs map[song.ID]struct{}
}
