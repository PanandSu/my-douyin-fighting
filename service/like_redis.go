package service

import (
	"fmt"
	"github.com/redis/go-redis/v9"
	rand2 "math/rand"
	gb "my-douyin-fighting/glob"
	"my-douyin-fighting/model"
)

type likeREdis struct{}

func (r *likeREdis) GetLikeStatus(uid, vid uint) (bool, error) {
	key := fmt.Sprintf("like:%d", uid)
	lua := redis.NewScript("")
	keys := []string{key}
	args := []any{
		vid,
		gb.VideoExpire.Seconds() + gb.ExpireTimeJitter.Seconds()*rand2.Float64(),
	}
	lua.Run(gb.Ctx, cache, keys, args)
}
func (r *likeREdis) AddLikeVids(uid uint, likes []model.Like) error {
	key := fmt.Sprintf("like:%d", uid)
	cache.TxPipelined(gb.Ctx, func(pipe redis.Pipeliner) error {
		pipe.ZAdd(gb.Ctx, key, redis.Z{
			Score:  2,
			Member: "",
		})
		for _, like := range likes {
			if like.IsLike {
				pipe.ZAdd(gb.Ctx, key, redis.Z{
					Score:  1,
					Member: like.VideoId,
				})
			} else {
				pipe.ZAdd(gb.Ctx, key, redis.Z{
					Score:  0,
					Member: like.VideoId,
				})
			}
		}
		expired := gb.LikeExpire + gb.ExpireTimeJitter
		pipe.Expire(gb.Ctx, key, expired)
	})
}
func (r *likeREdis) Like(uid, vid, aid uint) error {
	ch := make(chan error, 2)
	defer close(ch)
	go func() {
		key := fmt.Sprintf("like:%d", uid)
		lua := redis.NewScript("")
		keys := []string{key}
		args := []any{
			gb.LikeExpire.Seconds() + gb.ExpireTimeJitter.Seconds()*rand2.Float64(),
		}
		_, err := lua.Run(gb.Ctx, cache, keys, args).Bool()
		ch <- err
	}()
	go func() {
		key := fmt.Sprintf("user:%d", uid)
		lua := redis.NewScript("")
		keys := []string{key}
		args := []any{
			gb.LikeExpire.Seconds() + gb.ExpireTimeJitter.Seconds()*rand2.Float64(),
		}
		_, err := lua.Run(gb.Ctx, cache, keys, args).Bool()
		ch <- err
	}()
	go func() {
		key := fmt.Sprintf("user:%d", aid)
		lua := redis.NewScript("")
		keys := []string{key}
		args := []any{
			gb.UserInfoExpire.Seconds() + gb.ExpireTimeJitter.Seconds()*rand2.Float64(),
		}
		_, err := lua.Run(gb.Ctx, cache, keys, args).Bool()
		ch <- err
	}()
	go func() {
		key := fmt.Sprintf("video:%d", vid)
		lua := redis.NewScript("")
		keys := []string{key}
		args := []any{
			gb.VideoExpire.Seconds() + gb.ExpireTimeJitter.Seconds()*rand2.Float64(),
		}
		_, err := lua.Run(gb.Ctx, cache, keys, args).Bool()
		ch <- err
	}()
	for i := range 4 {
		<-ch
		fmt.Println(i)
	}
}
func (r *likeREdis) Unlike(uid, vid, aid uint) error {
	ch := make(chan error, 2)
	defer close(ch)
	go func() {
		key := fmt.Sprintf("like:%d", uid)
		lua := redis.NewScript("")
		keys := []string{key}
		args := []any{
			uid,
			gb.LikeExpire.Seconds() + gb.ExpireTimeJitter.Seconds()*rand2.Float64(),
		}
		_, err := lua.Run(gb.Ctx, cache, keys, args).Bool()
		ch <- err
	}()
	go func() {
		key := fmt.Sprintf("user:%d", uid)
		lua := redis.NewScript("")
		keys := []string{key}
		args := []any{
			gb.UserInfoExpire.Seconds() + gb.ExpireTimeJitter.Seconds()*rand2.Float64(),
		}
		_, err := lua.Run(gb.Ctx, cache, keys, args).Bool()
		ch <- err
	}()
	go func() {
		key := fmt.Sprintf("user:%d", aid)
		lua := redis.NewScript("")
		keys := []string{key}
		args := []any{
			gb.UserInfoExpire.Seconds() + gb.ExpireTimeJitter.Seconds()*rand2.Float64(),
		}
		_, err := lua.Run(gb.Ctx, cache, keys, args).Bool()
		ch <- err
	}()
	go func() {
		key := fmt.Sprintf("video:%d", vid)
		lua := redis.NewScript("")
		keys := []string{key}
		args := []any{
			gb.VideoExpire.Seconds() + gb.ExpireTimeJitter.Seconds()*rand2.Float64(),
		}
		_, err := lua.Run(gb.Ctx, cache, keys, args).Bool()
		ch <- err
	}()
	for i := range 4 {
		<-ch
		fmt.Println(i)
	}
}
func (r *likeREdis) GetLikeVids(uid uint) ([]uint, error) {

}
func (r *likeREdis) GetLikeCount(vid uint) (int, error) {

}
func (r *likeREdis) AddLikeCount(vid uint, likeCount int) error {

}
func (r *likeREdis) GetLikeCountList(vids []uint) ([]int, error) {

}
func (r *likeREdis) AddLikeCountList(vs VideoLikeAPI) {

}
