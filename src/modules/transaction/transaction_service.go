package transaction

import (
	"context"
	"go-wallet/src/interfaces"
	"go-wallet/src/libs"
)

type transaction_service struct {
	transaction_repo interfaces.TransactionRepo
	user_repo        interfaces.UserRepo
}

func NewService(reps interfaces.TransactionRepo, user_repo interfaces.UserRepo) *transaction_service {
	return &transaction_service{
		transaction_repo: reps,
		user_repo:        user_repo,
	}
}

func (re *transaction_service) GetAllStatusTransaction(userId string) *libs.Response {
	redisKey := "*:*:" + userId
	data, err := re.transaction_repo.GetRedisTransaction(context.Background(), redisKey)
	if err != nil {
		return libs.New(err.Error(), 400, true)
	}

	return libs.New(data, 200, false)
}
