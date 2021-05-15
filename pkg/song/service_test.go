package song

import (
	"github.com/xiaomi388/virtual-music-system/pkg/song/driver"
	"github.com/xiaomi388/virtual-music-system/pkg/song/model"
	"testing"
)

func TestService_GetSongByNameArtist(t *testing.T) {
	s := Service{
		LocalRepo:   &driver.FileRepository{RootDir: "/tmp/filerepo"},
		RemoteRepos: []model.RemoteRepository{&driver.YoutubeRepository{}},
	}
	song, err := s.GetSongByNameArtist("海阔天空", "黄家驹")
	if err != nil {
		t.Fatalf(err.Error())
		return
	}
	exp := "/tmp/filerepo/黄家驹/海阔天空.mp3"
	if song.FilePath != exp {
		t.Fatalf("expect: %s, actual: %s", exp, song.FilePath)
		return
	}
}
