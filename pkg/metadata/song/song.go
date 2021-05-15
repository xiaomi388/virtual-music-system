package song

type ID string

type Song struct {
	ID         ID     `json:"id"`
	Name       string `json:"name"`
	ArtistName string `json:"artist_name"`
}
