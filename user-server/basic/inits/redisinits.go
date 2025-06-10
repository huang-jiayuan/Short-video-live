package inits

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
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
func RedisSet(key string, code int) {
	RedisClient.Set(context.Background(), key, code, time.Minute*5)
}
func RedisGet(key string) int {
	i, err := RedisClient.Get(context.Background(), key).Int()
	if err != nil {
		return 0
	}
	return i
}
func RedisIncr(key string) int64 {
	incr := RedisClient.Incr(context.Background(), key).Val()
	return incr
}
func RedisExpire(key string) {
	RedisClient.Expire(context.Background(), key, time.Second*60)
}
