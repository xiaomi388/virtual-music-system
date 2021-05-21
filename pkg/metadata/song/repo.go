// Package song contains codes related to domain model Song
package song

// SearchType represents kinds of metadata available for being used to search.
type SearchType int

// 搜索类型: 默认为 1 即单曲 , 取值意义 : 1: 单曲, 10: 专辑, 100: 歌手,
//		1000: 歌单, 1002: 用户, 1004: MV, 1006: 歌词, 1009: 电台, 1014: 视频, 1018:综合
const (
	SearchTypeSong     SearchType = 1
	SearchTypeAlbum    SearchType = 10
	SearchTypeArtist   SearchType = 100
	SearchTypePlaylist SearchType = 1000
	SearchTypeUser     SearchType = 1002
	SearchTypeMV       SearchType = 1004
	SearchTypeLyric    SearchType = 1009
	SearchTypeDjRadio  SearchType = 1009
	SearchTypeVideo    SearchType = 1014
	SearchTypeMultiple SearchType = 1018
)

// Repository provides CRUD interfaces for Song entities.
type Repository interface {
	GetSong(id ID) (Song, error)
	GetSongsByQuery(q string, limit int, offset int) (map[ID]Song, int, error)
	GetSongsByPlaylistID(pid string, limit int, offset int) (map[ID]Song, int, error)
	SearchByType(t SearchType, q string, limit, offset int) ([]byte, error)
}
