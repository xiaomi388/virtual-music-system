package netease

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"github.com/xiaomi388/virtual-music-system/pkg/metadata/playlist"
	"github.com/xiaomi388/virtual-music-system/pkg/metadata/song"
	"io/ioutil"
	"net/http"
	"strconv"
)

// PlaylistRepository retrieves playlist metadata from netease API.
type PlaylistRepository struct {
	Sr      song.Repository
	BaseURL string
	Client  http.Client
}

// GetPlaylist retrieves playlist metadata by a song's id from netease API.
func (r *PlaylistRepository) GetPlaylist(pid song.ID) (playlist.Playlist, error) {
	req, err := http.NewRequest("GET",
		fmt.Sprintf("%s/playlist/detail?id=%s", r.BaseURL, pid), nil)
	if err != nil {
		return playlist.Playlist{}, errors.Wrap(err, "http.NewRequest")
	}
	resp, err := r.Client.Do(req)
	if err != nil || resp == nil {
		return playlist.Playlist{}, errors.Wrap(err, "r.Client.Do")
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return playlist.Playlist{}, fmt.Errorf("resp returns %s", resp.Status)
	}
	respStr, err := ioutil.ReadAll(resp.Body)

	respObj := struct {
		Playlist struct {
			ID          int    `json:"id"`
			Name        string `json:"name"`
			CoverImgURL string `json:"coverImgUrl"`
			Creator     struct {
				Nickname string `json:"nickname"`
				UserID   int    `json:"userId"`
			} `json:"creator"`
			TrackCount  int    `json:"trackCount"`
			UserID      int    `json:"userId"`
			PlayCount   int    `json:"playCount"`
			BookCount   int    `json:"bookCount"`
			Description string `json:"description"`
			HighQuality bool   `json:"highQuality"`
			Tracks      []struct {
				Name    string `json:"name"`
				ID      int    `json:"id"`
				Artists []struct {
					ID   int    `json:"id"`
					Name string `json:"name"`
				} `json:"ar"`
			} `json:"tracks"`
		} `json:"playlist"`
	}{}

	err = json.Unmarshal(respStr, &respObj)
	if err != nil {
		return playlist.Playlist{}, errors.Wrap(err, "json.Unmarshal")
	}
	var songs []song.Song
	for _, track := range respObj.Playlist.Tracks {
		artistName := ""
		if len(track.Artists) > 0 {
			artistName = track.Artists[0].Name
		}
		songs = append(songs, song.Song{
			ID:         song.ID(strconv.Itoa(track.ID)),
			Name:       track.Name,
			ArtistName: artistName,
		})
	}
	playList := playlist.Playlist{
		ID:            song.ID(strconv.Itoa(respObj.Playlist.ID)),
		Name:          respObj.Playlist.Name,
		CoverImageUrl: respObj.Playlist.CoverImgURL,
		Description:   respObj.Playlist.Description,
		TrackCount:    respObj.Playlist.TrackCount,
		Songs:         songs,
	}

	return playList, nil
}

// GetPlaylistsByQuery retrieves playlists metadata from netease by searching a keyword.
func (r *PlaylistRepository) GetPlaylistsByQuery(q string,
	limit int, offset int) (map[song.ID]playlist.Playlist, int, error) {

	respStr, err := r.Sr.SearchByType(song.SearchTypePlaylist, q, limit, offset)
	if err != nil {
		return nil, 0, errors.Wrap(err, "resp.Body.Read")
	}

	var respObj struct {
		Result struct {
			Playlists []struct {
				ID          int    `json:"id"`
				Name        string `json:"name"`
				CoverImgURL string `json:"coverImgUrl"`
				Creator     struct {
					Nickname string `json:"nickname"`
					UserID   int    `json:"userId"`
				} `json:"creator"`
				TrackCount  int    `json:"trackCount"`
				UserID      int    `json:"userId"`
				PlayCount   int    `json:"playCount"`
				BookCount   int    `json:"bookCount"`
				Description string `json:"description"`
				HighQuality bool   `json:"highQuality"`
			} `json:"playlists"`
			PlaylistCount int `json:"playlistCount"`
		} `json:"result"`
		Code int `json:"code"`
	}

	err = json.Unmarshal(respStr, &respObj)
	if err != nil {
		return nil, 0, errors.Wrap(err, "json.Unmarshal")
	}

	playlists := map[song.ID]playlist.Playlist{}
	for _, p := range respObj.Result.Playlists {
		playlists[song.ID(strconv.Itoa(p.ID))] = playlist.Playlist{
			ID:            song.ID(strconv.Itoa(p.ID)),
			Name:          p.Name,
			CoverImageUrl: p.CoverImgURL,
			Description:   p.Description,
			TrackCount:    p.TrackCount,
		}
	}
	return playlists, 0, nil
}
