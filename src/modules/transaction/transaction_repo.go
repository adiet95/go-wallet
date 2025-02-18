package transaction

import (
	"context"
	"errors"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type transaction_repo struct {
	rd *redis.Client
}

func NewRepo(db *gorm.DB, rd *redis.Client) *transaction_repo {
	return &transaction_repo{
		rd: rd,
	}
}

func (repo *transaction_repo) GetRedisTransaction(ctx context.Context, key string) (map[string]string, error) {

	data := repo.rd.HGetAll(ctx, key)
	if data.Err() != nil {
		return nil, errors.New("failed to found data")
	}
	return data.Val(), nil
}

// func (repo *transaction_repo) ExecTrxCreate(tx *gorm.DB, data *entity.Payment) *gorm.DB {
// 	trx := tx.Create(data)
// 	return trx
// }
