package database

import (
	"os"
	"strconv"

	"github.com/redis/go-redis/v9"
)

func RedisClient() *redis.Client {
	toInt, _ := strconv.Atoi(os.Getenv("REDIS_DB"))
	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_URL"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       toInt,
	})
	return rdb
}
