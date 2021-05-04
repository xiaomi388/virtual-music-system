// metadata contains songs metadata including albums, artists, ...
package metadata

import "github.com/xiaomi388/virtual-music-system/metadata/song"

type Service struct {
	SongRepo song.Repository
}

func (s *Service) GetAlbumsByQuery() {

}

func (s *Service) GetArtistsByQuery() {

}

func (s *Service) GetPlaylistsByQuery() {

}

func (s *Service) GetSongsByQuery(q string, limit int, offset int) ([]song.Song, error) {
	songMap, err := s.SongRepo.GetSongsByQuery(q, limit, offset)
	if err != nil {
		return nil, err
	}
	songs := make([]song.Song, 0, len(songMap))
	for _, s := range songMap {
		songs = append(songs, s)
	}
	return songs, nil
}
