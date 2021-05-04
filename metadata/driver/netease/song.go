package netease

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"github.com/xiaomi388/virtual-music-system/metadata/song"
	"io/ioutil"
	"net/http"
	"strconv"
)

type SongRepository struct {
	BaseURL string
	Client  http.Client
}

func (r *SongRepository) GetSong(id song.ID) (song.Song, error) {
	url := fmt.Sprintf("%s/song/detail?ids=%s", r.BaseURL, id)
	resp, err := r.Client.Get(url)
	if err != nil {
		return song.Song{}, errors.Wrap(err, "r.Client.Do")
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return song.Song{}, fmt.Errorf("resp returns %s", resp.Status)
	}

	respStr, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return song.Song{}, errors.Wrap(err, "resp.Body.Read")
	}

	respObj := struct {
		Songs []struct {
			Name string `json:"name"`
			ID   int    `json:"id"`
		} `json:"songs"`
	}{}

	err = json.Unmarshal(respStr, &respObj)
	if err != nil {
		return song.Song{}, errors.Wrap(err, "json.Unmarshal")
	}

	if len(respObj.Songs) == 0 {
		return song.Song{}, fmt.Errorf("song %s doesn't exist", id)
	}

	s := song.Song{
		ID:   song.ID(strconv.Itoa(respObj.Songs[0].ID)),
		Name: respObj.Songs[0].Name,
	}
	return s, nil
}

func (r *SongRepository) GetSongsByQuery(q string,
	limit int, offset int) (map[song.ID]song.Song, error) {

	req, err := http.NewRequest("GET",
		fmt.Sprintf("%s/search", r.BaseURL), nil)
	if err != nil {
		return nil, errors.Wrap(err, "http.NewRequest")
	}

	params := req.URL.Query()
	params.Add("type", "1")
	params.Add("limit", strconv.Itoa(limit))
	params.Add("offset", strconv.Itoa(offset))
	params.Add("keywords", q)
	req.URL.RawQuery = params.Encode()

	resp, err := r.Client.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "r.Client.Do")
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("resp returns %s", resp.Status)
	}

	respStr, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "resp.Body.Read")
	}

	respObj := struct {
		Result struct {
			Songs []struct {
				Name    string `json:"name"`
				ID      int    `json:"id"`
				Artists []struct {
					Name string `json:"name"`
				} `json:"artists"`
			} `json:"songs"`
		} `json:"result"`
	}{}

	err = json.Unmarshal(respStr, &respObj)
	if err != nil {
		return nil, errors.Wrap(err, "json.Unmarshal")
	}

	songs := map[song.ID]song.Song{}
	for _, s := range respObj.Result.Songs {
		artistName := ""
		if len(s.Artists) > 0 {
			artistName = s.Artists[0].Name
		}
		songs[song.ID(strconv.Itoa(s.ID))] = song.Song{
			ID:         song.ID(strconv.Itoa(s.ID)),
			Name:       s.Name,
			ArtistName: artistName,
		}
	}
	return songs, nil
}
