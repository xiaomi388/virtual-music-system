package intf

import "github.com/xiaomi388/virtual-music-system/song"
import "github.com/gin-gonic/gin"

type HTTPService struct {
	Service *song.Service
	GE      *gin.Engine
}

func (s *HTTPService) GetSongByNameArtist(c *gin.Context) {
	name := c.Query("name")
	artist := c.Query("artist")
	song, err := s.Service.GetSongByNameArtist(name, artist)
	if err != nil {
		_ = c.AbortWithError(400, err)
		return
	}
	c.File(song.FilePath)
}

// not threading safe
func (s *HTTPService) Register() {
	s.GE.GET("/v1/song", s.GetSongByNameArtist)
}
