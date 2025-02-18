package interfaces

import (
	"context"
	"go-wallet/src/libs"
	"go-wallet/src/models"
	"time"
)

type TopUpRepo interface {
	SetRedisTopUp(ctx context.Context, key string, data *models.TopUp, ttl time.Duration) error
	GetRedisTopUp(ctx context.Context, key string) (map[string]string, error)
	DelRedisPayment(ctx context.Context, key string) error
}

type TopUpService interface {
	PostTopUp(data *models.TopUpRequest, userId string) *libs.Response
	GetPendingStatusTopUp(userId string) *libs.Response
	GetAllStatusTopUp(userId string) *libs.Response
}
