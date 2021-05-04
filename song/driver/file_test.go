package driver

import (
	"os"
	"testing"
)

func TestFileRepository_CopySongFrom(t *testing.T) {
	r := FileRepository{RootDir: "/tmp/filerepo"}
	f, err := os.Create("/tmp/mock.mp3")
	if err != nil {
		t.Fatalf(err.Error())
		return
	}
	defer f.Close()
	song, err := r.CopySongFrom("/tmp/mock.mp3", "name", "artist")
	if err != nil {
		t.Fatalf(err.Error())
		return
	}
	exp := "/tmp/filerepo/artist/name.mp3"
	if song.FilePath != exp {
		t.Fatalf("expect: %s, actual: %s", exp, song.FilePath)
		return
	}
}
