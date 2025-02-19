package interfaces

import (
	"context"
	"time"
)

type RedisRepo interface {
	SetRedis(ctx context.Context, key string, data map[string]interface{}, ttl time.Duration) error
	GetRedis(ctx context.Context, key string) (string, error)
	DelRedis(ctx context.Context, key string) error
	SearchKey(ctx context.Context, pattern string) (string, error)
	SearchKeyArr(ctx context.Context, pattern string) ([]string, error)
}
