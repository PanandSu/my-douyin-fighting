package service

import (
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	rand2 "math/rand"
	gb "my-douyin-fighting/glob"
	"my-douyin-fighting/model"
	"strconv"
	"time"
)

func GetUserInfo(uid uint) (*model.User, error) {
	// 定义 key
	key := fmt.Sprintf("user:%d", uid)

	var user model.User
	if result := gb.Redis.Exists(gb.Ctx, key).Val(); result <= 0 {
		return nil, errors.New("not found in cache")
	}
	// 使用 pipeline
	cmds, err := gb.Redis.TxPipelined(gb.Ctx, func(pipe redis.Pipeliner) error {
		pipe.HGetAll(gb.Ctx, key)
		pipe.HGet(gb.Ctx, key, "created_at").Val()
		// 设置过期时间
		pipe.Expire(gb.Ctx, key, gb.UserInfoExpire+time.Duration(rand2.Float64()*gb.ExpireTimeJitter.Seconds())*time.Second)
		return nil
	})
	if err != nil {
		return nil, err
	}
	if err = cmds[0].(*redis.StringCmd).Scan(&user); err != nil {
		return nil, err
	}
	timeUnixMilliStr := cmds[1].(*redis.StringCmd).Val()
	timeUnixMilli, _ := strconv.ParseInt(timeUnixMilliStr, 10, 64)
	user.CreatedAt = time.UnixMilli(timeUnixMilli)
	return &user, nil
}

func AddUserInfo(user *model.User) error {
	// 定义 key
	userRedis := fmt.Sprintf("user:%d", user.Id)

	// 使用 pipeline
	_, err := gb.Redis.TxPipelined(gb.Ctx, func(pipe redis.Pipeliner) error {
		pipe.HSet(gb.Ctx, userRedis, "user_id", user.Id)
		pipe.HSet(gb.Ctx, userRedis, "name", user.Name)
		pipe.HSet(gb.Ctx, userRedis, "password", user.Password)
		pipe.HSet(gb.Ctx, userRedis, "follow_count", user.FollowerCount)
		pipe.HSet(gb.Ctx, userRedis, "follower_count", user.FollowerCount)
		pipe.HSet(gb.Ctx, userRedis, "total_favorited", user.TotalLike)
		pipe.HSet(gb.Ctx, userRedis, "favorite_count", user.LikeCount)
		pipe.HSet(gb.Ctx, userRedis, "created_at", user.CreatedAt.UnixMilli())
		// 设置过期时间
		pipe.Expire(gb.Ctx, userRedis, gb.UserInfoExpire+time.Duration(rand2.Float64()*gb.ExpireTimeJitter.Seconds())*time.Second)
		return nil
	})
	return err
}

func GetUserListByUserIDList(uids []uint) (userList []model.User, notInCache []uint, err error) {
	// 定义 key
	userNum := len(uids)
	userList = make([]model.User, 0, userNum)
	notInCache = make([]uint, 0, userNum)
	for _, each := range uids {
		user, err2 := GetUserInfo(each)
		if err2 != nil && err2.Error() != "not found in cache" {
			return nil, nil, err2
		} else if err2 == nil {
			userList = append(userList, *user)
		} else {
			err = err2
			userList = append(userList, model.User{Id: each})
			notInCache = append(notInCache, each)
		}
	}
	return
}

func AddUserListByUserIDLists(users []model.User) error {
	// 使用 pipeline
	_, err := gb.Redis.TxPipelined(gb.Ctx, func(pipe redis.Pipeliner) error {
		for _, each := range users {
			// 定义 key
			key := fmt.Sprintf("user:%d", each.Id)

			pipe.HSet(gb.Ctx, key, "user_id", each.Id)
			pipe.HSet(gb.Ctx, key, "name", each.Name)
			pipe.HSet(gb.Ctx, key, "password", each.Password)
			pipe.HSet(gb.Ctx, key, "follow_count", each.FollowCount)
			pipe.HSet(gb.Ctx, key, "follower_count", each.FollowerCount)
			pipe.HSet(gb.Ctx, key, "total_favorited", each.TotalLike)
			pipe.HSet(gb.Ctx, key, "favorite_count", each.LikeCount)
			pipe.HSet(gb.Ctx, key, "created_at", each.CreatedAt.UnixMilli())
			// 设置过期时间
			pipe.Expire(gb.Ctx, key, gb.UserInfoExpire+time.Duration(rand2.Float64()*gb.ExpireTimeJitter.Seconds())*time.Second)
		}
		return nil
	})
	return err
}
