package artist

type Repository interface {
	GetArtist(id ID) (Artist, error)
	GetArtistsByQuery(q string, limit int, offset int) (map[ID]Artist, error)
}
