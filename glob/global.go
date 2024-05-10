package gb

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
	Cfg            config.System
	DB             *gorm.DB
	Redis          *redis.Client
	Ctx            = context.Background()
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
var (
	FavoriteExpire      = 10 * time.Minute
	VideoCommentsExpire = 10 * time.Minute
	CommentExpire       = 10 * time.Minute
	FollowExpire        = 10 * time.Minute
	UserInfoExpire      = 10 * time.Minute
	VideoExpire         = 10 * time.Minute
	PublishExpire       = 10 * time.Minute
	EmptyExpire         = 10 * time.Minute
	ExpireTimeJitter    = 10 * time.Minute
)
