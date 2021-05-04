package intf

import (
	"github.com/gin-gonic/gin"
	"github.com/xiaomi388/virtual-music-system/metadata"
	"strconv"
)

type HTTPService struct {
	Service *metadata.Service
	GE      *gin.Engine
}

func (s *HTTPService) GetSongsByQuery(c *gin.Context) {
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

	songs, err := s.Service.GetSongsByQuery(q, limit, offset)
	if err != nil {
		_ = c.AbortWithError(400, err)
		return
	}
	c.JSON(200, songs)
	return
}

// not threading safe
func (s *HTTPService) Register() {
	s.GE.GET("/v1/metadata/songs", s.GetSongsByQuery)
}
