package metadata

import (
	"github.com/xiaomi388/virtual-music-system/pkg/metadata/playlist"
	"github.com/xiaomi388/virtual-music-system/pkg/metadata/song"
)

// Service is an application service for CRUD of song-related metadata
type Service struct {
	SongRepo     song.Repository
	PlayListRepo playlist.Repository
}

func (s *Service) GetAlbumsByQuery() {

}

func (s *Service) GetArtistsByQuery() {

}

// GetPlaylistsByQuery returns playlist.Playlist entities to user.
func (s *Service) GetPlaylistsByQuery(q string, limit int, offset int) ([]playlist.Playlist, int, error) {
	playlistsMap, total, err := s.PlayListRepo.GetPlaylistsByQuery(q, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	playLists := make([]playlist.Playlist, 0, len(playlistsMap))
	for _, p := range playlistsMap {
		playLists = append(playLists, p)
	}
	return playLists, total, nil
}

// GetSongsByPlaylistID returns a playlist.Playlist entity of a specific id.
func (s *Service) GetSongsByPlaylistID(q string, limit int, offset int) ([]song.Song, int, error) {
	songMap, total, err := s.SongRepo.GetSongsByPlaylistID(q, limit, offset)
	if err != nil {
		return nil, 0, err
	}
	songs := make([]song.Song, 0, len(songMap))
	for _, s := range songMap {
		songs = append(songs, s)
	}
	return songs, total, nil
}

// GetSongsByQuery returns song.Song entities by searching a song keyword.
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

// GetPlaylistByID returns a playlist.Playlist entity of a specific id.
func (s *Service) GetPlaylistByID(pid string) (playlist.Playlist, error) {
	pl, err := s.PlayListRepo.GetPlaylist(song.ID(pid))
	if err != nil {
		return pl, err
	}

	return pl, nil
}
