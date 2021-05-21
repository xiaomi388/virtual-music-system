package driver

import "testing"

func TestYoutubeRepository_GetSongByNameArtist(t *testing.T) {
	r := YoutubeRepository{}
	fp, err := r.GetSongByNameArtist("十年", "陈奕迅")
	if err != nil {
		t.Fatalf(err.Error())
		return
	}
	exp := "/tmp/十年_陈奕迅.mp3"
	if fp != exp {
		t.Fatalf("expected: %s, actual: %s", exp, fp)
		return
	}
}
