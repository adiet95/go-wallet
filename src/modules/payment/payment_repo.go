package payment

import (
	"context"
	"errors"
	"time"

	"go-wallet/src/libs"
	"go-wallet/src/models"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type payment_repo struct {
	rd *redis.Client
}

func NewRepo(db *gorm.DB, rd *redis.Client) *payment_repo {
	return &payment_repo{
		rd: rd,
	}
}

func (repo *payment_repo) SetRedisPayment(ctx context.Context, key string, data *models.Payment, ttl time.Duration) error {
	ObjectRedis, err := libs.StructToMap(data)
	if err != nil {
		return errors.New("failed to found data")
	}
	err = repo.rd.HSet(ctx, key, ObjectRedis, ttl).Err()
	if err != nil {
		return errors.New("failed to found data")
	}
	return err
}

func (repo *payment_repo) GetRedisPayment(ctx context.Context, key string) (map[string]string, error) {

	data := repo.rd.HGetAll(ctx, key)
	if data.Err() != nil {
		return nil, errors.New("failed to found data")
	}
	return data.Val(), nil
}

func (repo *payment_repo) DelRedisPayment(ctx context.Context, key string) error {
	data := repo.rd.Del(ctx, key)
	if data.Err() != nil {
		return errors.New("failed to found data")
	}
	return nil
}

// func (repo *payment_repo) ExecTrxCreate(tx *gorm.DB, data *entity.Payment) *gorm.DB {
// 	trx := tx.Create(data)
// 	return trx
// }
