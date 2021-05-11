// metadata contains songs metadata including albums, artists, ...
package metadata

import (
	"github.com/xiaomi388/virtual-music-system/metadata/playlist"
	"github.com/xiaomi388/virtual-music-system/metadata/song"
)

type Service struct {
	SongRepo     song.Repository
	PlayListRepo playlist.Repository
}

func (s *Service) GetAlbumsByQuery() {

}

func (s *Service) GetArtistsByQuery() {

}

func (s *Service) GetPlaylistsByQuery(q string, limit int, offset int) ([]playlist.PlayList, int, error) {
	playlistsMap, total, err := s.PlayListRepo.GetPlaylistsByQuery(q, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	playLists := make([]playlist.PlayList, 0, len(playlistsMap))
	for _, p := range playlistsMap {
		playLists = append(playLists, p)
	}
	return playLists, total, nil
}

func (s *Service) GetSongsByPlaylistId(q string, limit int, offset int) ([]song.Song, int, error) {
	songMap, total, err := s.SongRepo.GetSongsByPlayListId(q, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	songs := make([]song.Song, 0, len(songMap))
	for _, s := range songMap {
		songs = append(songs, s)
	}
	return songs, total, nil
}

func (s *Service) GetSongsByQuery(q string, limit int, offset int) ([]song.Song, int, error) {
	songMap, total, err := s.SongRepo.GetSongsByQuery(q, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	songs := make([]song.Song, 0, len(songMap))
	for _, s := range songMap {
		songs = append(songs, s)
	}
	return songs, total, nil
}

func (s *Service) GetPlayListById(pid string) (playlist.PlayList, error) {
	playList, err := s.PlayListRepo.GetPlaylist(song.ID(pid))
	if err != nil {
		return playList, err
	}

	return playList, nil
}
