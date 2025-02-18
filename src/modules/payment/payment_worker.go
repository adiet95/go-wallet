package payment

import (
	"go-wallet/src/modules/users"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func NewWorker(db *gorm.DB, rd *redis.Client) {
	repo := NewRepo(db, rd)
	userRepo := users.NewRepo(db)
	svc := NewService(repo, userRepo)
	svc.WorkerPayment()
}
