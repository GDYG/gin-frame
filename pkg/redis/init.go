package redis

import (
	"github.com/redis/go-redis/v9"
)

var rdb *redis.Client

func InitRedis() *redis.Client {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "39.96.5.17:6379",
		Password: "",
		DB:       0,
	})

	return rdb
}
