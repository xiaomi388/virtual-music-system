// Package playlist contains codes related to domain model Playlist
package playlist

import "github.com/xiaomi388/virtual-music-system/pkg/metadata/song"

// Playlist contains metadata of a playlist, and what songs are in this playlist.
type Playlist struct {
	ID            song.ID     `json:"id"`
	Name          string      `json:"name"`
	CoverImageUrl string      `json:"cover_image_url"`
	Description   string      `json:"description"`
	TrackCount    int         `json:"track_count"` // how many songs are in the playlist
	Songs         []song.Song `json:"songs,omitempty"`
}
