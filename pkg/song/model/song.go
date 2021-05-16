package model

// ID is the id of a Song entity
type ID string

// Song contains a song file metadata
type Song struct {
	ID       ID
	Name     string
	Artist   string
	FilePath string
}
