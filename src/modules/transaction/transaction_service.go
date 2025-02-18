package transaction

import (
	"context"
	"go-wallet/src/interfaces"
	"go-wallet/src/libs"
	"go-wallet/src/models"

	"github.com/google/uuid"
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
	redisKey := "transaction:*:" + userId
	data, err := re.transaction_repo.GetRedisTransaction(context.Background(), redisKey)
	if err != nil {
		return libs.New(err.Error(), 400, true)
	}

	return libs.New(data, 200, false)
}

func (re *transaction_service) GetPendingStatusTransaction(userId string) *libs.Response {
	redisKey := "transaction:pending:" + userId
	data, err := re.transaction_repo.GetRedisTransaction(context.Background(), redisKey)
	if err != nil {
		return libs.New(err.Error(), 400, true)
	}

	return libs.New(data, 200, false)
}

func (re *transaction_service) PostTransaction(data *models.TransactionRequest, userId string) *libs.Response {
	result := models.Transaction{}
	userData, err := re.user_repo.FindById(userId)
	if err != nil {
		return libs.New(err.Error(), 400, true)
	}

	timeNow, err := libs.TimeNow()
	if err != nil {
		return libs.New(err.Error(), 400, true)
	}

	dataEntity := &models.Transaction{
		TransactionId:     uuid.New().String(),
		UserId:            userId,
		AmountTransaction: data.Amount,
		BalanceBefore:     userData.Balance,
		BalanceAfter:      userData.Balance + data.Amount,
		Remarks:           data.Remarks,
		Status:            "PENDING",
		CreatedDate:       timeNow,
	}

	redisKey := "transaction:pending:" + userId

	err = re.transaction_repo.SetRedisTransaction(context.Background(), redisKey, dataEntity, 0)
	if err != nil {
		return libs.New(err.Error(), 400, true)
	}
	return libs.New(result, 200, false)
}
