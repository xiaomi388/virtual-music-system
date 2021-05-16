package song

// ID is the type of Song.ID.
type ID string

// Song contains the metadata of a song.
type Song struct {
	ID         ID     `json:"id"`
	Name       string `json:"name"`
	ArtistName string `json:"artist_name"`
}
