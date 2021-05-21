package intf

import "github.com/xiaomi388/virtual-music-system/pkg/song"
import "github.com/gin-gonic/gin"

// HTTPService is a HTTP adapter to Service
type HTTPService struct {
	Service *song.Service
	GE      *gin.Engine
}

// GetSongByNameArtist ports the application service Service.GetSongByNameArtist by using HTTP
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

// Register creates routes from url paths to handlers
func (s *HTTPService) Register() {
	s.GE.GET("/v1/song", s.GetSongByNameArtist)
}
