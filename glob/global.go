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
	Cfg               config.System
	DB                *gorm.DB
	Redis             *redis.Client
	Ctx               = context.Background()
	FileTypeMap       sync.Map
	IdGenerator       *sonyflake.Sonyflake
	StartTime         = "2022-05-21 00:00:01"
	VideoAddr         = "./public/video"
	CoverAddr         = "./public/cover"
	MaxFileSize       = int64(10 << 20)
	MaxTitleLen       = 140
	MaxCommentLen     = 300
	FeedNum           = 30
	MaxUsernameLength = 32
	PasswordPattern   = `^[_0-9A-Za-z]{6,32}$`
	AutoCreateDB      = true
	WhitelistVideo    = map[string]bool{
		".mp4":  true,
		".avi":  true,
		".wmv":  true,
		".mpeg": true,
		".mov":  true,
		".flv":  true,
		".rmvb": true,
		".3gb":  true,
		".vob":  true,
		".m4v":  true,
	}
)

var (
	LikeExpire          = 10 * time.Minute
	VideoCommentsExpire = 10 * time.Minute
	CommentExpire       = 10 * time.Minute
	FollowExpire        = 10 * time.Minute
	UserInfoExpire      = 10 * time.Minute
	VideoExpire         = 10 * time.Minute
	PublishExpire       = 10 * time.Minute
	EmptyExpire         = 10 * time.Minute
	ExpireTimeJitter    = 10 * time.Minute
)
