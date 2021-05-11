package intf

import (
	"github.com/gin-gonic/gin"
	"github.com/xiaomi388/virtual-music-system/metadata"
	"github.com/xiaomi388/virtual-music-system/metadata/song"
	"strconv"
)

type HTTPService struct {
	Service *metadata.Service
	GE      *gin.Engine
}

type Pager struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
	Total  int `json:"total"`
}

func (s *HTTPService) GetSongsByQuery(c *gin.Context) {
	q := c.Query("q")
	pid := c.Query("pid")
	limitStr := c.DefaultQuery("limit", "20")
	offsetStr := c.DefaultQuery("offset", "0")
	if q == "" && pid == "" {
		c.AbortWithStatusJSON(400, gin.H{
			"error":   400,
			"message": "q or pid is a must parameter",
		})
		return
	} else if q != "" && pid != "" {
		c.AbortWithStatusJSON(400, gin.H{
			"error":   400,
			"message": "only choose one from q or pid",
		})
		return
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		_ = c.AbortWithError(400, err)
		return
	}

	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		_ = c.AbortWithError(400, err)
		return
	}
	var songs []song.Song
	var total int
	if pid == "" {
		songs, total, err = s.Service.GetSongsByQuery(q, limit, offset)
	} else {
		songs, total, err = s.Service.GetSongsByPlaylistId(pid, limit, offset)
	}
	if err != nil {
		_ = c.AbortWithError(500, err)
		return
	}
	s.SendSuccessWithPager(c, songs, &Pager{Limit: limit, Offset: offset, Total: total})
	return
}

func (s *HTTPService) GetPlayListsByQuery(c *gin.Context) {
	q := c.Query("q")
	limitStr := c.DefaultQuery("limit", "20")
	offsetStr := c.DefaultQuery("offset", "0")
	if q == "" {
		c.AbortWithStatusJSON(400, gin.H{
			"error":   400,
			"message": "q is a must parameter",
		})
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		_ = c.AbortWithError(400, err)
		return
	}

	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		_ = c.AbortWithError(400, err)
		return
	}

	playlists, total, err := s.Service.GetPlaylistsByQuery(q, limit, offset)
	if err != nil {
		_ = c.AbortWithError(400, err)
		return
	}
	s.SendSuccessWithPager(c, playlists, &Pager{Limit: limit, Offset: offset, Total: total})
	return
}

func (s *HTTPService) GetPlayListByID(c *gin.Context) {
	pid := c.Query("id")

	if pid == "" {
		c.AbortWithStatusJSON(400, gin.H{
			"error":   400,
			"message": "pid is a must parameter",
		})
	}
	playList, err := s.Service.GetPlayListById(pid)
	if err != nil {
		_ = c.AbortWithError(400, err)
		return
	}
	s.SendSuccessWithPager(c, playList, nil)
	return
}

func (s *HTTPService) GetSongsByPlayListId(c *gin.Context) {
	pid := c.Query("pid")
	limitStr := c.DefaultQuery("limit", "20")
	offsetStr := c.DefaultQuery("offset", "0")
	if pid == "" {
		c.AbortWithStatusJSON(400, gin.H{
			"error":   400,
			"message": "pid is a must parameter",
		})
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		_ = c.AbortWithError(400, err)
		return
	}

	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		_ = c.AbortWithError(400, err)
		return
	}

	songs, total, err := s.Service.GetSongsByPlaylistId(pid, limit, offset)
	if err != nil {
		_ = c.AbortWithError(400, err)
		return
	}
	s.SendSuccessWithPager(c, songs, &Pager{Limit: limit, Offset: offset, Total: total})
	return
}

func (s *HTTPService) SendSuccessWithPager(ctx *gin.Context, resp interface{}, pager *Pager) {
	type Result struct {
		Result interface{} `json:"result"`
		Pager  *Pager      `json:"pager,omitempty"`
	}
	ctx.JSON(200, Result{resp, pager})
}

// not threading safe
func (s *HTTPService) Register() {
	s.GE.GET("/v1/metadata/songs", s.GetSongsByQuery)
	s.GE.GET("/v1/metadata/playlists", s.GetPlayListsByQuery)
	s.GE.GET("/v1/metadata/playlist", s.GetPlayListByID)
	s.GE.GET("/v1/metadata/playlist/songs", s.GetSongsByPlayListId)
}
