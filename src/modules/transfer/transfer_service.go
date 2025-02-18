package order

import (
	"context"
	"go-wallet/src/interfaces"
	"go-wallet/src/libs"
	"go-wallet/src/models"

	"github.com/google/uuid"
)

type transfer_service struct {
	transfer_repo interfaces.TransferRepo
	user_repo     interfaces.UserRepo
}

func NewService(reps interfaces.TransferRepo, user_repo interfaces.UserRepo) *transfer_service {
	return &transfer_service{
		transfer_repo: reps,
		user_repo:     user_repo,
	}
}

func (re *transfer_service) GetAllStatusTransfer(userId string) *libs.Response {
	redisKey := "transfer:*:" + userId
	data, err := re.transfer_repo.GetRedisTransfer(context.Background(), redisKey)
	if err != nil {
		return libs.New(err.Error(), 400, true)
	}

	return libs.New(data, 200, false)
}

func (re *transfer_service) GetPendingStatusTransfer(userId string) *libs.Response {
	redisKey := "transfer:pending:" + userId
	data, err := re.transfer_repo.GetRedisTransfer(context.Background(), redisKey)
	if err != nil {
		return libs.New(err.Error(), 400, true)
	}

	return libs.New(data, 200, false)
}

func (re *transfer_service) PostTransfer(data *models.TransferRequest, userId string) *libs.Response {
	result := models.Transfer{}
	userData, err := re.user_repo.FindById(userId)
	if err != nil {
		return libs.New(err.Error(), 400, true)
	}

	timeNow, err := libs.TimeNow()
	if err != nil {
		return libs.New(err.Error(), 400, true)
	}

	dataEntity := &models.Transfer{
		TransferId:     uuid.New().String(),
		UserId:         userId,
		AmountTransfer: data.Amount,
		BalanceBefore:  userData.Balance,
		BalanceAfter:   userData.Balance + data.Amount,
		Remarks:        data.Remarks,
		Status:         "PENDING",
		CreatedDate:    timeNow,
	}

	redisKey := "transfer:pending:" + userId

	err = re.transfer_repo.SetRedisTransfer(context.Background(), redisKey, dataEntity, 0)
	if err != nil {
		return libs.New(err.Error(), 400, true)
	}
	return libs.New(result, 200, false)
}
