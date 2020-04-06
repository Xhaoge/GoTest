package redis

import (
	"strings"

	"github.com/go-redis/redis"
	"rr-factory.gloryholiday.com/yuetu/golang-core/logger"
)

func newClusterRedisClient(addrs []string) *redis.ClusterClient {
	logger.InfoNt(logger.Message("Connecting redis cluster: %s", strings.Join(addrs, ",")))
	result := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: addrs,
	})

	return result
}
