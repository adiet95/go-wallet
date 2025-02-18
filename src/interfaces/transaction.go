package interfaces

import (
	"context"
	"go-wallet/src/libs"
)

type TransactionRepo interface {
	GetRedisTransaction(ctx context.Context, key string) (map[string]string, error)
}

type TransactionService interface {
	GetAllStatusTransaction(userId string) *libs.Response
}
