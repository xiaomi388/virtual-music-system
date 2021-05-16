package driver

import (
	"fmt"
	"github.com/pkg/errors"
	"os/exec"
)

// YoutubeRepository implements model.RemoteRepository.
// It retrieve songs from youtube.
type YoutubeRepository struct {
}

// GetSongByNameArtist get a song from youtube by a song's name and artist name.
func (r *YoutubeRepository) GetSongByNameArtist(name string, artist string) (filePath string, err error) {
	fileName := fmt.Sprintf("/tmp/%s_%s", name, artist)
	filePath = fileName + ".mp3"

	cmd := exec.Command("youtube-dl", "--extract-audio", "--audio-format", "mp3",
		"--ignore-errors", "--output", fileName+".%(ext)s",
		fmt.Sprintf("ytsearch1: %s %s", name, artist))
	err = cmd.Run()
	if err != nil {
		err = errors.Wrapf(err, "run cmd:%s failed", cmd.String())
		return
	}
	return
}
