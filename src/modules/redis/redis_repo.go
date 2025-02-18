package redisrepo

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"

	"go-wallet/src/constant"

	"github.com/redis/go-redis/v9"
)

type redis_repo struct {
	rd *redis.Client
}

func NewRepo(rd *redis.Client) *redis_repo {
	return &redis_repo{
		rd: rd,
	}
}

func (repo *redis_repo) SetRedis(ctx context.Context, key string, data map[string]interface{}, ttl time.Duration) error {
	bytesData, _ := json.Marshal(&data)
	err := repo.rd.Set(ctx, key, bytesData, ttl).Err()
	if err != nil {
		fmt.Println(err)
		return errors.New(constant.RedisSetError)
	}
	return err
}

func (repo *redis_repo) GetRedis(ctx context.Context, key string) (string, error) {
	data := repo.rd.Get(ctx, key)
	if data.Err() != nil {
		return "", errors.New(constant.RedisGetError)
	}
	return data.Val(), nil
}

func (repo *redis_repo) SearchKey(ctx context.Context, pattern string) (string, error) {
	keys, err := repo.rd.Keys(ctx, pattern).Result()
	if err != nil {
		log.Fatalf("Error scanning keys: %v", err)
	}
	for _, key := range keys {
		return key, nil
	}

	return "", nil
}

func (repo *redis_repo) DelRedis(ctx context.Context, key string) error {
	data := repo.rd.GetDel(ctx, key)
	if data.Err() != nil {
		return errors.New(constant.RedisDeleteError)
	}
	return nil
}
