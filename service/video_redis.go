package service

import (
	"fmt"
	"github.com/redis/go-redis/v9"
	"math"
	rand2 "math/rand"
	"math/rand/v2"
	gb "my-douyin-fighting/glob"
	"my-douyin-fighting/model"
	"time"
)

type videoRedis struct{}

// GoPublish 将用户发表过的视频写入缓存中
func (r *followRedis) GoPublish(uid uint, listZ ...redis.Z) error {
	//定义 key
	key := fmt.Sprintf("publish:%d", uid)
	pipe := gb.Redis.TxPipeline()
	pipe.ZAdd(gb.Ctx, key, listZ...)
	pipe.Expire(gb.Ctx, key, gb.PublishExpire+time.Duration(rand2.Float64()*gb.ExpireTimeJitter.Seconds())*time.Second)
	_, err := pipe.Exec(gb.Ctx)
	return err
}

// GoVideoList 将视频批量写入缓存
func (r *followRedis) GoVideoList(videoList []model.Video) error {
	pipe := gb.Redis.TxPipeline()
	for _, video := range videoList {
		key := fmt.Sprintf("video:%d", video.VideoId)
		pipe.HSet(gb.Ctx, key, "title", video.Title, "play_name", video.PlayName, "cover_name", video.CoverName,
			"favorite_count", video.FavoriteCount, "comment_count", video.CommentCount, "author_id", video.AuthorId, "created_at", video.CreatedAt.UnixMilli())
		pipe.Expire(gb.Ctx, key, gb.VideoExpire+time.Duration(rand2.Float64()*gb.ExpireTimeJitter.Seconds())*time.Second)
	}
	_, err := pipe.Exec(gb.Ctx)
	return err
}

// GoFeed 确保feed在缓存中
func (r *followRedis) GoFeed() error {
	n, err := gb.Redis.Exists(gb.Ctx, "feed").Result()
	if err != nil {
		return err
	}
	if n <= 0 {
		// "feed"不存在
		var allVideos []model.Video
		if err := gb.DB.Find(&allVideos).Error; err != nil {
			return err
		}
		if len(allVideos) == 0 {
			return nil
		}
		var listZ = make([]redis.Z, 0, len(allVideos))
		for _, video := range allVideos {
			listZ = append(listZ, redis.Z{Score: float64(video.CreatedAt.UnixMilli()) / 1000, Member: video.VideoId})
		}
		return gb.Redis.ZAdd(gb.Ctx, "feed", listZ...).Err()
	}
	return nil
}

// PublishEvent 用户上视频的缓存操作
func (r *followRedis) PublishEvent(video model.Video, listZ ...redis.Z) error {
	publish := fmt.Sprintf("publish:%d", video.AuthorId)
	keyVideo := fmt.Sprintf("video:%d", video.VideoId)
	empty := fmt.Sprintf("empty:%d", video.AuthorId)
	pipe := gb.Redis.TxPipeline()
	pipe.ZAdd(gb.Ctx, "feed", redis.Z{Score: float64(video.CreatedAt.UnixMilli()) / 1000, Member: keyVideo})
	pipe.ZAdd(gb.Ctx, publish, listZ...)
	pipe.Expire(gb.Ctx, publish, gb.PublishExpire+time.Duration(rand.Float64()*gb.ExpireTimeJitter.Seconds())*time.Second)

	pipe.HSet(gb.Ctx, keyVideo, "author_id", video.AuthorId, "play_name", video.PlayName, "cover_name", video.CoverName,
		"favorite_count", video.FavoriteCount, "comment_count", video.CommentCount, "title", video.Title, "created_at", video.CreatedAt.UnixMilli())
	pipe.Expire(gb.Ctx, keyVideo, gb.VideoExpire+time.Duration(rand.Float64()*gb.ExpireTimeJitter.Seconds())*time.Second)
	pipe.Del(gb.Ctx, empty)
	_, err := pipe.Exec(gb.Ctx)
	return err
}

func (r *followRedis) GoCommentsOfVideo(commentList []model.Comment, keyCommentsOfVideo string) error {
	var listZ = make([]redis.Z, 0, len(commentList))
	for _, comment := range commentList {
		listZ = append(listZ, redis.Z{Score: float64(comment.CreatedAt.UnixMilli()) / 1000, Member: comment.Id})
	}
	pipe := gb.Redis.TxPipeline()
	pipe.ZAdd(gb.Ctx, keyCommentsOfVideo, listZ...)
	pipe.Expire(gb.Ctx, keyCommentsOfVideo, gb.VideoCommentsExpire+time.Duration(rand.Float64()*gb.ExpireTimeJitter.Seconds())*time.Second)
	_, err := pipe.Exec(gb.Ctx)
	return err
}

func (r *followRedis) GetCommentCountOfVideo(videoID uint64) (int, error) {
	key := fmt.Sprintf("video:%d", videoID)
	lua := redis.NewScript(`
				local key = KEYS[1]
				local expire_time = ARGV[1]
				if redis.call("Exists", key) > 0 then
					redis.call("Expire", key, expire_time)
					return redis.call("HGet", key, "comment_count")
				end
				return -1
			`)
	keys := []string{key}
	values := []any{gb.VideoCommentsExpire.Seconds() + math.Floor(rand2.Float64()*gb.ExpireTimeJitter.Seconds())}
	numComments, err := lua.Run(gb.Ctx, gb.Redis, keys, values).Int()
	if err != nil {
		return 0, err
	}
	return numComments, nil
}

func (r *followRedis) SetUserPublishEmpty(userID uint64) error {
	key := fmt.Sprintf("empty:%d", userID)
	return gb.Redis.Set(gb.Ctx, key, "1", gb.EmptyExpire+time.Duration(rand.Float64()*gb.ExpireTimeJitter.Seconds())*time.Second).Err()
}
