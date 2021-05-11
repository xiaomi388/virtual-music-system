package netease

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"github.com/xiaomi388/virtual-music-system/metadata/song"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

type SongRepository struct {
	pr      PlaylistRepository
	BaseURL string
	Client  http.Client
}

func (r *SongRepository) GetSong(id song.ID) (song.Song, error) {
	songs, err := r.GetSongs([]song.ID{id})
	if err == nil && len(songs) != 0 {
		return songs[0], err
	}
	return song.Song{}, err
}

func (r *SongRepository) GetSongs(ids []song.ID) ([]song.Song, error) {
	if len(ids) == 0 {
		return nil, nil
	}
	//join ids with ','
	n := len(",") * (len(ids) - 1)
	for i := 0; i < len(ids); i++ {
		n += len(ids[i])
	}
	var b strings.Builder
	b.Grow(n)
	b.WriteString(string(ids[0]))
	for _, s := range ids[1:] {
		b.WriteString(",")
		b.WriteString(string(s))
	}
	idsStr := b.String()
	url := fmt.Sprintf("%s/song/detail?ids=%s", r.BaseURL, idsStr)
	resp, err := r.Client.Get(url)
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
		Songs []struct {
			Name   string `json:"name"`
			ID     int    `json:"id"`
			Artist []struct {
				ID   int    `json:"id"`
				Name string `json:"name"`
			} `json:"ar"`
		} `json:"songs"`
	}{}

	err = json.Unmarshal(respStr, &respObj)
	if err != nil {
		return nil, errors.Wrap(err, "json.Unmarshal")
	}

	if len(respObj.Songs) == 0 {
		return nil, fmt.Errorf("song %+v doesn't exist", ids)
	}

	var ss []song.Song
	for _, s := range respObj.Songs {
		artistName := ""
		if len(s.Artist) != 0 {
			artistName = s.Artist[0].Name
		}
		ss = append(ss, song.Song{
			ID:         song.ID(strconv.Itoa(s.ID)),
			Name:       s.Name,
			ArtistName: artistName,
		})
	}
	return ss, nil
}

func (r *SongRepository) GetSongsByQuery(q string,
	limit int, offset int) (map[song.ID]song.Song, int, error) {

	respStr, err := r.SearchByType(song.SearchTypeSong, q, limit, offset)
	if err != nil {
		return nil, 0, errors.Wrap(err, "resp.Body.Read")
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
			SongCount int `json:"songCount"`
		} `json:"result"`
	}{}

	err = json.Unmarshal(respStr, &respObj)
	if err != nil {
		return nil, 0, errors.Wrap(err, "json.Unmarshal")
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
	return songs, respObj.Result.SongCount, nil
}

func (r *SongRepository) GetSongsByPlayListId(pid string,
	limit int, offset int) (map[song.ID]song.Song, int, error) {
	req, err := http.NewRequest("GET",
		fmt.Sprintf("%s/playlist/detail?id=%s", r.BaseURL, pid), nil)
	if err != nil {
		return nil, 0, errors.Wrap(err, "http.NewRequest")
	}
	resp, err := r.Client.Do(req)
	if err != nil || resp == nil {
		return nil, 0, errors.Wrap(err, "r.Client.Do")
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, 0, fmt.Errorf("resp returns %s", resp.Status)
	}
	respStr, err := ioutil.ReadAll(resp.Body)

	respObj := struct {
		Result struct {
			TrackIds []struct {
				ID int64 `json:"id"`
			} `json:"trackIds"`
		} `json:"playlist"`
	}{}

	err = json.Unmarshal(respStr, &respObj)
	if err != nil {
		return nil, 0, errors.Wrap(err, "json.Unmarshal")
	}
	//all song's ids
	var songIds []int64
	for _, track := range respObj.Result.TrackIds {
		songIds = append(songIds, track.ID)
	}
	var ids []song.ID
	for i := offset * limit; i < offset*limit+limit; i++ {
		if songIds == nil || len(songIds) < i+1 {
			break
		}
		ids = append(ids, song.ID(strconv.Itoa(int(songIds[i]))))
	}

	songs, err := r.GetSongs(ids)
	//why convert slice to map?
	songsMap := map[song.ID]song.Song{}
	for i := range songs {
		songsMap[songs[i].ID] = songs[i]
	}
	return songsMap, len(songIds), nil
}

func (r *SongRepository) SearchByType(t song.SearchType, q string,
	limit, offset int) ([]byte, error) {
	req, err := http.NewRequest("GET",
		fmt.Sprintf("%s/search", r.BaseURL), nil)
	if err != nil {
		return nil, errors.Wrap(err, "http.NewRequest")
	}

	params := req.URL.Query()
	params.Add("type", strconv.Itoa(int(t)))
	params.Add("limit", strconv.Itoa(limit))
	params.Add("offset", strconv.Itoa(offset))
	params.Add("keywords", q)
	req.URL.RawQuery = params.Encode()

	resp, err := r.Client.Do(req)
	if err != nil || resp == nil {
		return nil, errors.Wrap(err, "r.Client.Do")
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("resp returns %s", resp.Status)
	}

	return ioutil.ReadAll(resp.Body)
}
