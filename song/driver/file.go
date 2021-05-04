package driver

import (
	"github.com/xiaomi388/virtual-music-system/song/model"
	"os"
	"path/filepath"
)

type FileRepository struct {
	RootDir string
}

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
