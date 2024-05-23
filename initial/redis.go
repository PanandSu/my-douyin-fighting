package initial

import (
	"github.com/redis/go-redis/v9"
	"my-douyin-fighting/glob"
	"strconv"
)

func Redis() {
	redisConfig := gb.Cfg.RedisConfig
	rdb := redis.NewClient(&redis.Options{
		Addr:     redisConfig.Host + strconv.Itoa(redisConfig.Port),
		Password: redisConfig.Password,
	})
	rdb.Ping(gb.Ctx)
	gb.Redis = rdb
}
