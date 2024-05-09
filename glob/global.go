package glob

import (
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/sony/sonyflake"
	"gorm.io/gorm"
	"my-douyin-fighting/config"
	"sync"
	"time"
)

var (
	Config         config.System
	DB             *gorm.DB
	Redis          *redis.Client
	Context        = context.Background()
	FileTypeMap    sync.Map
	IdGenerator    *sonyflake.Sonyflake
	StartTime      time.Time
	VideoAddr      string
	CoverAddr      string
	MaxFileSize    int64
	MaxTitleLen    = 140
	MaxCommentLen  = 300
	AutoCreateDB   = true
	WhitelistVideo = map[string]bool{
		".mp4":  true,
		".avi":  true,
		".flv":  true,
		".mov":  true,
		".mp3":  true,
		".ogg":  true,
		".wav":  true,
		".aac":  true,
		".aiff": true,
	}
)
