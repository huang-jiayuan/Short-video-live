package inits

import (
	"github.com/go-redis/redis/v8"
)

var (
	RedisClient *redis.Client
)

func RedisInit() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     "14.103.149.197:6379",
		Password: "91F6E0538A51D156E652FF47755BB44E", // no password set
		DB:       0,                                  // use default DB
	})
}
