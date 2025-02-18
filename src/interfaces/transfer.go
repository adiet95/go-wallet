package interfaces

import (
	"context"
	"go-wallet/src/libs"
	"go-wallet/src/models"
	"time"
)

type TransferRepo interface {
	SetRedisTransfer(ctx context.Context, key string, data *models.Transfer, ttl time.Duration) error
	GetRedisTransfer(ctx context.Context, key string) (map[string]string, error)
	DelRedisPayment(ctx context.Context, key string) error
}

type TransferService interface {
	PostTransfer(data *models.TransferRequest, userId string) *libs.Response
	GetPendingStatusTransfer(userId string) *libs.Response
	GetAllStatusTransfer(userId string) *libs.Response
}
