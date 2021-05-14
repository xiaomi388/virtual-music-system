package song

import (
	"github.com/xiaomi388/virtual-music-system/song/model"
	"testing"
)

type mockLocalRepo struct{}
type mockRemoteRepo struct{}

func (r *mockLocalRepo) GetSongByNameArtist(name string, artist string) (model.Song, error) {
	return model.Song{}, nil
}

func (r *mockLocalRepo) CopySongFrom(filePath string, name string, artist string) (model.Song, error) {
	song := model.Song{
		ID:       "mock",
		Name:     name,
		Artist:   artist,
		FilePath: "/repo/hello_world.mp3",
	}
	return song, nil
}

func (r *mockRemoteRepo) GetSongByNameArtist(name string, artist string) (string, error) {
	return "/tmp/hello_world.mp3", nil
}

func TestService_GetSongByNameArtist(t *testing.T) {
	s := Service{
		LocalRepo: &mockLocalRepo{},
		RemoteRepos: []model.RemoteRepository{
			&mockRemoteRepo{},
		},
	}
	name, artist := "海阔天空", "黄家驹"
	song, err := s.GetSongByNameArtist("海阔天空", "黄家驹")
	if err != nil {
		t.Fatalf(err.Error())
		return
	}
	exp := model.Song{
		ID:       "mock",
		Name:     name,
		Artist:   artist,
		FilePath: "/repo/hello_world.mp3",
	}
	if song != exp {
		t.Fatalf("expect: %s, actual: %s", exp, song.FilePath)
		return
	}
}
