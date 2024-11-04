package driver

import (
	"github.com/redis/go-redis/v9"
	"github.com/ritik-patel05/image-player/config"
)

var cacheClient *redis.Client

func init() {
	cacheClient = newRedisClient()
}

func newRedisClient() *redis.Client {
	imageRedisConfig := config.GetConfig().Redis.ImageRedis

	rdb := redis.NewClient(&redis.Options{
		Addr:     imageRedisConfig.Host[0],
		Password: "",
		DB:       0,
	})

	return rdb
}

func GetCacheClient() *redis.Client {
	return cacheClient
}
