package interfaces

import (
	"context"
	"go-wallet/src/libs"
	"go-wallet/src/models"
	"time"
)

type PaymentRepo interface {
	SetRedisPayment(ctx context.Context, key string, data *models.Payment, ttl time.Duration) error
	GetRedisPayment(ctx context.Context, key string) (map[string]string, error)
	DelRedisPayment(ctx context.Context, key string) error
}

type PaymentService interface {
	PostPayment(data *models.PaymentRequest, userId string) *libs.Response
	GetPendingPaymentStatus(userId string) *libs.Response
	GetAllPaymentStatus(userId string) *libs.Response
}
