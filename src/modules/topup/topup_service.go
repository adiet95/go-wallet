package order

import (
	"context"
	"go-wallet/src/interfaces"
	"go-wallet/src/libs"
	"go-wallet/src/models"

	"github.com/google/uuid"
)

type topup_service struct {
	topup_repo interfaces.TopUpRepo
	user_repo  interfaces.UserRepo
}

func NewService(reps interfaces.TopUpRepo, user_repo interfaces.UserRepo) *topup_service {
	return &topup_service{
		topup_repo: reps,
		user_repo:  user_repo,
	}
}

func (re *topup_service) GetAllStatusTopUp(userId string) *libs.Response {
	redisKey := "topup:*:" + userId
	data, err := re.topup_repo.GetRedisTopUp(context.Background(), redisKey)
	if err != nil {
		return libs.New(err.Error(), 400, true)
	}

	return libs.New(data, 200, false)
}

func (re *topup_service) GetPendingStatusTopUp(userId string) *libs.Response {
	redisKey := "topup:pending:" + userId
	data, err := re.topup_repo.GetRedisTopUp(context.Background(), redisKey)
	if err != nil {
		return libs.New(err.Error(), 400, true)
	}

	return libs.New(data, 200, false)
}

func (re *topup_service) PostTopUp(data *models.TopUpRequest, userId string) *libs.Response {
	result := models.TopUp{}
	userData, err := re.user_repo.FindById(userId)
	if err != nil {
		return libs.New(err.Error(), 400, true)
	}

	timeNow, err := libs.TimeNow()
	if err != nil {
		return libs.New(err.Error(), 400, true)
	}

	dataEntity := &models.TopUp{
		TopUpId:       uuid.New().String(),
		UserId:        userId,
		AmountTopUp:   data.Amount,
		BalanceBefore: userData.Balance,
		BalanceAfter:  userData.Balance + data.Amount,
		Remarks:       data.Remarks,
		Status:        "PENDING",
		CreatedDate:   timeNow,
	}

	redisKey := "topup:pending:" + userId

	err = re.topup_repo.SetRedisTopUp(context.Background(), redisKey, dataEntity, 0)
	if err != nil {
		return libs.New(err.Error(), 400, true)
	}
	return libs.New(result, 200, false)
}
