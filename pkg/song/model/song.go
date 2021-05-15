package model

type ID string

type Song struct {
	ID       ID
	Name     string
	Artist   string
	FilePath string
}
