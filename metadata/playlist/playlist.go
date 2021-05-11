package playlist

import "github.com/xiaomi388/virtual-music-system/metadata/song"

type PlayList struct {
	ID            song.ID     `json:"id"`
	Name          string      `json:"name"`
	CoverImageUrl string      `json:"cover_image_url"`
	Description   string      `json:"description"`
	TrackCount    int         `json:"track_count"` // how many songs are in the playlist
	Songs         []song.Song `json:"songs,omitempty"`
}
