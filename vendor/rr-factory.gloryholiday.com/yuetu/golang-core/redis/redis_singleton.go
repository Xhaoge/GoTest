package redis

import (
	"github.com/go-redis/redis"
	"rr-factory.gloryholiday.com/yuetu/golang-core/logger"
)

func newSingletonRedisClient(addr string) *redis.Client {
	logger.InfoNt(logger.Message("Connecting to redis singleton: %s", addr))
	result := redis.NewClient(&redis.Options{
		Addr: addr,
	})

	return result
}
