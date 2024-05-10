package service

import (
	"fmt"
	"github.com/redis/go-redis/v9"
	rand2 "math/rand"
	gb "my-douyin-fighting/glob"
	"my-douyin-fighting/model"
	"time"
)

var cache = gb.Redis

func AddCommentRedis(c *model.Comment) error {
	com := fmt.Sprintf("CommentsOfVideo:%d", c.Id)
	video := fmt.Sprintf("Video:%d", c.VideoId)
	vid_com := fmt.Sprintf("Comment:%d", c.Id)
	cache.Exists(gb.Ctx)
	lua1 := redis.NewScript("")
	keys1 := []string{vid_com}
	args1 := []any{
		c.CreatedAt.String(),
		c.Id,
		gb.VideoCommentsExpire.Seconds() + gb.ExpireTimeJitter.Seconds()*rand2.Float64(),
	}
	lua1.Run(gb.Ctx, cache, keys1, args1)
	keys2 := []string{video}
	args2 := []any{
		gb.VideoExpire.Seconds() + gb.ExpireTimeJitter.Seconds()*rand2.Float64(),
	}
	lua2 := redis.NewScript("")
	lua2.Run(gb.Ctx, cache, keys2, args2)
	vid := c.VideoId
	uid := c.UserId
	pipe := cache.TxPipeline()
	pipe.Expire(gb.Ctx, com,
		gb.CommentExpire+gb.ExpireTimeJitter,
	)
	pipe.HSet(gb.Ctx, com,
		"comment", c.Content,
		"user_id", uid,
		"video_id", vid,
		"created_at", time.Now().UnixMilli(),
	)
	pipe.Exec(gb.Ctx)
}

func DelCommentRedis(vid, cid uint) error {
	video := fmt.Sprintf("Video:%d", vid)
	com := fmt.Sprintf("Comment:%d", cid)
	vid_com := fmt.Sprintf("CommentsOfVideo:%d", cid)
	lua1 := redis.NewScript("")
	keys1 := []string{vid_com}
	args1 := []any{
		cid,
		gb.VideoCommentsExpire.Seconds() + gb.ExpireTimeJitter.Seconds()*rand2.Float64(),
	}
	lua1.Run(gb.Ctx, cache, keys1, args1)
	lua2 := redis.NewScript("")
	keys2 := []string{video}
	args2 := []any{
		gb.VideoExpire.Seconds() + gb.ExpireTimeJitter.Seconds()*rand2.Float64(),
	}
	lua2.Run(gb.Ctx, cache, keys2, args2)
	cache.Del(gb.Ctx)
}

func GoComment(c *model.Comment) error {
	key := fmt.Sprintf("Comment:%d", c.Id)
	lua := redis.NewScript("")
	keys := []string{key}
	args := []any{
		c.Content,
		c.VideoId,
		c.UserId,
		c.CreatedAt,
		gb.CommentExpire.Seconds() + gb.ExpireTimeJitter.Seconds()*rand2.Float64(),
	}
	lua.Run(gb.Ctx, cache, keys, args)
}

/*
key := fmt.Sprintf(":%d", )
lua := redis.NewScript("")
keys := []string{key}
args := []any{
gb. .Seconds() + gb.ExpireTimeJitter.Seconds()*rand2.Float64(),
}
lua.Run(gb.Ctx, cache, keys, args)
*/
