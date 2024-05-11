package service

import (
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"math"
	rand2 "math/rand"
	gb "my-douyin-fighting/glob"
	"my-douyin-fighting/model"
	"time"
)

type followRedis struct{}

func (r *followRedis) GetFollowStatus(fid, cid uint) (bool, error) {
	// 定义 key
	key := fmt.Sprintf("follower:%d", fid)
	lua := redis.NewScript("")
	keys := []string{key}
	values := []any{cid, gb.FollowExpire.Seconds() + math.Floor(rand2.Float64()*gb.ExpireTimeJitter.Seconds())}
	result, err := lua.Run(gb.Ctx, gb.Redis, keys, values).Bool()
	if err == nil {
		return result, nil
	} else if errors.Is(err, redis.Nil) {
		return false, errors.New("not found in cache")
	} else {
		return false, err
	}
}

func (r *followRedis) AddFollowIDList(fid uint, celebrityList []model.Follow) error {
	// 定义 key
	key := fmt.Sprintf("follower:%d", fid)
	// 使用 pipeline
	_, err := gb.Redis.TxPipelined(gb.Ctx, func(pipe redis.Pipeliner) error {
		// 初始化
		pipe.ZAdd(gb.Ctx, key, redis.Z{Score: 2, Member: ""})
		// 增加点赞关系
		for _, each := range celebrityList {
			if each.IsFollow {
				pipe.ZAdd(gb.Ctx, key, redis.Z{Score: 1, Member: each.CelebrityId})
			} else {
				pipe.ZAdd(gb.Ctx, key, redis.Z{Score: 0, Member: each.CelebrityId})
			}
		}
		// 设置过期时间
		pipe.Expire(gb.Ctx, key, gb.FollowExpire+time.Duration(rand2.Float64()*gb.ExpireTimeJitter.Seconds())*time.Second)
		return nil
	})
	return err
}

func (r *followRedis) Follow(fid, cid uint) error {
	// 设置管道
	ch := make(chan error, 2)
	defer close(ch)

	// 更新 followerRelationRedis 缓存
	go func() {
		// 定义 key
		key := fmt.Sprintf("follower:%d", fid)
		lua := redis.NewScript(`
				if redis.call("Exists", KEYS[1]) > 0 then
					redis.call("ZAdd", KEYS[1], 1, ARGV[1])
					redis.call("Expire", KEYS[1], ARGV[2])
					return true
				end
				return false
			`)
		keys := []string{key}
		values := []any{cid, gb.FollowExpire.Seconds() + math.Floor(rand2.Float64()*gb.ExpireTimeJitter.Seconds())}
		_, err := lua.Run(gb.Ctx, gb.Redis, keys, values).Bool()
		ch <- err
	}()

	// 更新 celebrityRelationRedis 缓存
	go func() {
		// 定义 key
		key := fmt.Sprintf("follower:%d", cid)
		lua := redis.NewScript(`
				if redis.call("Exists", KEYS[1]) > 0 then
					redis.call("ZAdd", KEYS[1], 1, ARGV[1])
					redis.call("Expire", KEYS[1], ARGV[2])
					return true
				end
				return false
			`)
		keys := []string{key}
		values := []any{fid, gb.FollowExpire.Seconds() + math.Floor(rand2.Float64()*gb.ExpireTimeJitter.Seconds())}
		_, err := lua.Run(gb.Ctx, gb.Redis, keys, values).Bool()
		ch <- err
	}()

	// 更新 followerRedis 缓存
	go func() {
		// 定义 key
		key := fmt.Sprintf("follower:%d", fid)
		lua := redis.NewScript(`
				if redis.call("Exists", KEYS[1]) > 0 then
					redis.call("HIncrBy", KEYS[1], "follow_count", 1)
					redis.call("Expire", KEYS[1], ARGV[1])
					return true
				end
				return false
			`)
		keys := []string{key}
		values := []any{gb.UserInfoExpire.Seconds() + math.Floor(rand2.Float64()*gb.ExpireTimeJitter.Seconds())}
		_, err := lua.Run(gb.Ctx, gb.Redis, keys, values).Bool()
		ch <- err
	}()

	// 更新 celebrityRedis 缓存
	go func() {
		// 定义 key
		key := fmt.Sprintf("celebrity:%d", cid)
		lua := redis.NewScript(`
				if redis.call("Exists", KEYS[1]) > 0 then
					redis.call("HIncrBy", KEYS[1], "follower_count", 1)
					redis.call("Expire", KEYS[1], ARGV[1])
					return true
				end
				return false
			`)
		keys := []string{key}
		values := []any{gb.UserInfoExpire.Seconds() + math.Floor(rand2.Float64()*gb.ExpireTimeJitter.Seconds())}
		_, err := lua.Run(gb.Ctx, gb.Redis, keys, values).Bool()
		ch <- err
	}()

	var err error
	for i := 0; i < 4; i++ {
		errTmp := <-ch
		if errTmp != nil && !errors.Is(errTmp, redis.Nil) {
			err = errTmp
		}
	}
	return err
}

func (r *followRedis) Unfollow(fid, cid uint) error {
	// 设置管道
	ch := make(chan error, 2)
	defer close(ch)

	// 更新 followerRelationRedis 缓存
	go func() {
		// 定义 key
		key := fmt.Sprintf("follower:%d", fid)
		lua := redis.NewScript(`
				if redis.call("Exists", KEYS[1]) > 0 then
					redis.call("ZAdd", KEYS[1], 0, ARGV[1])
					redis.call("Expire", KEYS[1], ARGV[2])
					return true
				end
				return false
			`)
		keys := []string{key}
		values := []any{cid, gb.FollowExpire.Seconds() + math.Floor(rand2.Float64()*gb.ExpireTimeJitter.Seconds())}
		_, err := lua.Run(gb.Ctx, gb.Redis, keys, values).Bool()
		ch <- err
	}()

	// 更新 celebrityRelationRedis 缓存
	go func() {
		// 定义 key
		key := fmt.Sprintf("celebrity:%d", cid)
		lua := redis.NewScript(`
				if redis.call("Exists", KEYS[1]) > 0 then
					redis.call("ZAdd", KEYS[1], 0, ARGV[1])
					redis.call("Expire", KEYS[1], ARGV[2])
					return true
				end
				return false
			`)
		keys := []string{key}
		values := []any{fid, gb.FollowExpire.Seconds() + math.Floor(rand2.Float64()*gb.ExpireTimeJitter.Seconds())}
		_, err := lua.Run(gb.Ctx, gb.Redis, keys, values).Bool()
		ch <- err
	}()

	// 更新 followerRedis 缓存
	go func() {
		// 定义 key
		key := fmt.Sprintf("follower:%d", fid)
		lua := redis.NewScript(`
				if redis.call("Exists", KEYS[1]) > 0 then
					redis.call("HIncrBy", KEYS[1], "follow_count", -1)
					redis.call("Expire", KEYS[1], ARGV[1])
					return true
				end
				return false
			`)
		keys := []string{key}
		values := []any{gb.UserInfoExpire.Seconds() + math.Floor(rand2.Float64()*gb.ExpireTimeJitter.Seconds())}
		_, err := lua.Run(gb.Ctx, gb.Redis, keys, values).Bool()
		ch <- err
	}()

	// 更新 celebrityRedis 缓存
	go func() {
		// 定义 key
		key := fmt.Sprintf("follower:%d", cid)
		lua := redis.NewScript(`
				if redis.call("Exists", KEYS[1]) > 0 then
					redis.call("HIncrBy", KEYS[1], "follower_count", -1)
					redis.call("Expire", KEYS[1], ARGV[1])
					return true
				end
				return false
			`)
		keys := []string{key}
		values := []any{gb.UserInfoExpire.Seconds() + math.Floor(rand2.Float64()*gb.ExpireTimeJitter.Seconds())}
		_, err := lua.Run(gb.Ctx, gb.Redis, keys, values).Bool()
		ch <- err
	}()

	var err error
	for i := 0; i < 4; i++ {
		errTmp := <-ch
		if errTmp != nil && !errors.Is(errTmp, redis.Nil) {
			err = errTmp
		}
	}
	return err
}

func (r *followRedis) GetFollowIDList(fid uint) ([]uint, error) {
	// 定义 key
	key := fmt.Sprintf("follower:%d", fid)
	lua := redis.NewScript(`
			if redis.call("Exists", KEYS[1]) <= 0 then
				return false
			end
			redis.call("Expire", KEYS[1], ARGV[1])
			return redis.call("ZRangeByScore", KEYS[1], 1, 1)
			`)
	keys := []string{key}
	values := []any{gb.FollowExpire.Seconds() + math.Floor(rand2.Float64()*gb.ExpireTimeJitter.Seconds())}
	result, err := lua.Run(gb.Ctx, gb.Redis, keys, values).Uint64Slice()
	var ints any
	ints = result
	res := ints.([]uint)
	if err == nil {
		return res, nil
	} else if errors.Is(err, redis.Nil) {
		return nil, errors.New("not found in cache")
	} else {
		return nil, err
	}
}

func (r *followRedis) GetFollowerIDList(cid uint) ([]uint, error) {
	// 定义 key
	key := fmt.Sprintf("celebrity:%d", cid)
	lua := redis.NewScript(`
			if redis.call("Exists", KEYS[1]) <= 0 then
				return false
			end
			redis.call("Expire", KEYS[1], ARGV[1])
			return redis.call("ZRangeByScore", KEYS[1], 1, 1)
			`)
	keys := []string{key}
	values := []any{gb.FollowExpire.Seconds() + math.Floor(rand2.Float64()*gb.ExpireTimeJitter.Seconds())}
	result, err := lua.Run(gb.Ctx, gb.Redis, keys, values).Uint64Slice()
	var ints any
	ints = result
	res := ints.([]uint)
	if err == nil {
		return res, nil
	} else if errors.Is(err, redis.Nil) {
		return nil, errors.New("not found in cache")
	} else {
		return nil, err
	}
}

func (r *followRedis) AddFollowerIDList(cid uint, followers []model.Follow) error {
	// 定义 key
	key := fmt.Sprintf("celebrity:%d", cid)
	// 使用 pipeline
	_, err := gb.Redis.TxPipelined(gb.Ctx, func(pipe redis.Pipeliner) error {
		//初始化
		pipe.ZAdd(gb.Ctx, key, redis.Z{Score: 2, Member: ""})
		// 增加点赞关系
		for _, each := range followers {
			if each.IsFollow {
				pipe.ZAdd(gb.Ctx, key, redis.Z{Score: 1, Member: each.FollowerId})
			} else {
				pipe.ZAdd(gb.Ctx, key, redis.Z{Score: 0, Member: each.FollowerId})
			}
		}
		//设置过期时间
		pipe.Expire(gb.Ctx, key, gb.FollowExpire+time.Duration(rand2.Float64()*gb.ExpireTimeJitter.Seconds())*time.Second)
		return nil
	})
	return err
}
