package netease

import (
	"net/http"
	"testing"
)

func TestSongRepository_GetSong(t *testing.T) {
	r := SongRepository{
		BaseURL: "http://127.0.0.1:3000",
		Client:  http.Client{},
	}
	song, err := r.GetSong("347230")
	if err != nil {
		t.Fatalf(err.Error())
	}
	if song.ID != "347230" {
		t.Fatal()
	}
}

func TestSongRepository_GetSongsByQuery(t *testing.T) {
	r := SongRepository{
		BaseURL: "http://127.0.0.1:3000",
		Client:  http.Client{},
	}
	songs, err := r.GetSongsByQuery("海阔天空", 1, 0)
	if err != nil {
		t.Fatalf(err.Error())
	}

	if _, ok := songs["347230"]; !ok {
		t.Fatal()
	}
}
