package driver

import (
	"github.com/xiaomi388/virtual-music-system/pkg/song/model"
	"os"
	"path/filepath"
)

// FileRepository implements model.LocalRepository.
// It stores songs in the local file system.
type FileRepository struct {
	RootDir string
}

// GetSongByNameArtist get a song from the local file system
// by the name and artist of the song.
func (r *FileRepository) GetSongByNameArtist(name string, artist string) (model.Song, error) {
	fp := filepath.Join(r.RootDir, artist, name+".mp3")

	if _, err := os.Stat(fp); err == nil {
		return model.Song{Name: name, ID: model.ID(name + "_" + artist), FilePath: fp}, nil
	} else if os.IsNotExist(err) {
		return model.Song{}, nil
	} else {
		return model.Song{}, err
	}
}

// CopySongFrom copies a song file located in filePath to "RepoRootDir/name/artist/name_artist.mp3".
func (r *FileRepository) CopySongFrom(filePath string, name string, artist string) (model.Song, error) {
	dirPath := filepath.Join(r.RootDir, artist)
	_ = os.MkdirAll(dirPath, 0755)

	dst := filepath.Join(dirPath, name+".mp3")
	src := filePath

	err := os.Rename(src, dst)
	if err != nil {
		return model.Song{}, err
	}

	s := model.Song{
		ID:       model.ID(name + "_" + artist),
		Name:     name,
		Artist:   artist,
		FilePath: dst,
	}
	return s, err
}
