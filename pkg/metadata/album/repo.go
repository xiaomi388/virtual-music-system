package album

// Repository contains Album entities.
type Repository interface {
	GetAlbum(id ID) (Album, error)
	GetAlbumsByQuery(q string, limit int, offset int) (map[ID]Album, error)
}
