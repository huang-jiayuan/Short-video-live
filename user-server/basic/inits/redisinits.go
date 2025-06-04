package inits

import (
	"github.com/go-redis/redis/v8"
	"user-server/basic/config"
)

var (
	RedisClient *redis.Client
)

func RedisInit() {
	na := config.Appconf.Redis
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     na.Host,
		Password: na.Password, // no password set
		DB:       na.Db,       // use default DB
	})
}
