package redis

import (
	"time"

	"github.com/go-redis/redis"
)

type RedisClient interface {
	Set(key string, value interface{}, expiration time.Duration) *redis.StatusCmd
	Get(key string) *redis.StringCmd
	TxPipeline() redis.Pipeliner
	GetSet(key string, value interface{}) *redis.StringCmd
	Del(keys ...string) *redis.IntCmd
	Expire(key string, expiration time.Duration) *redis.BoolCmd
	Close() error
	Decr(key string) *redis.IntCmd
	DecrBy(key string, decrement int64) *redis.IntCmd
	Incr(key string) *redis.IntCmd
	IncrBy(key string, value int64) *redis.IntCmd
	HIncrBy(key, field string, incr int64) *redis.IntCmd
	TTL(key string) *redis.DurationCmd
	HMGet(key string, fields ...string) *redis.SliceCmd
	HMSet(key string, fields map[string]interface{}) *redis.StatusCmd
	HGetAll(key string) *redis.StringStringMapCmd
	HDel(key string, fields ...string) *redis.IntCmd
	SetNX(key string, value interface{}, expiration time.Duration) *redis.BoolCmd
}

func NewRedisClient(redisNodes []string) RedisClient {
	if len(redisNodes) == 1 {
		return newSingletonRedisClient(redisNodes[0])
	} else {
		return newClusterRedisClient(redisNodes)
	}
}
