package song

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/xiaomi388/virtual-music-system/song/model"
)

type Service struct {
	LocalRepo   model.LocalRepository
	RemoteRepos []model.RemoteRepository
}

func (s *Service) GetSongByNameArtist(name string, artist string) (model.Song, error) {
	song, err := s.LocalRepo.GetSongByNameArtist(name, artist)
	if err != nil {
		return model.Song{}, errors.Wrap(err, "GetSongByName")
	}

	if song.ID != "" {
		return song, nil
	}

	// song not found
	// FIXME: only trigger download one time if
	// many users request a same song at the same time
	var fp string
	for _, rr := range s.RemoteRepos {
		fp, err = rr.GetSongByNameArtist(name, artist)
		if err != nil {
			continue
		}
		if fp != "" {
			break
		}
	}
	if err != nil {
		return model.Song{}, errors.Wrap(err, "GetSongByNameArtist")
	}
	if fp == "" {
		return model.Song{}, fmt.Errorf("song not found")
	}
	song, err = s.LocalRepo.CopySongFrom(fp, name, artist)
	if err != nil {
		return model.Song{}, errors.Wrap(err, "CopySongFrom")
	}
	return song, nil
}
