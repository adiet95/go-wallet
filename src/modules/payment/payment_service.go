package order

import (
	"context"
	"go-wallet/src/interfaces"
	"go-wallet/src/libs"
	"go-wallet/src/models"

	"github.com/google/uuid"
)

type payment_service struct {
	payment_repo interfaces.PaymentRepo
	user_repo    interfaces.UserRepo
}

func NewService(reps interfaces.PaymentRepo, user_repo interfaces.UserRepo) *payment_service {
	return &payment_service{
		payment_repo: reps,
		user_repo:    user_repo,
	}
}

func (re *payment_service) GetAllPaymentStatus(userId string) *libs.Response {
	redisKey := "payment:*:" + userId
	data, err := re.payment_repo.GetRedisPayment(context.Background(), redisKey)
	if err != nil {
		return libs.New(err.Error(), 400, true)
	}

	return libs.New(data, 200, false)
}

func (re *payment_service) GetPendingPaymentStatus(userId string) *libs.Response {
	redisKey := "payment:pending:" + userId
	data, err := re.payment_repo.GetRedisPayment(context.Background(), redisKey)
	if err != nil {
		return libs.New(err.Error(), 400, true)
	}

	return libs.New(data, 200, false)
}

func (re *payment_service) PostPayment(data *models.PaymentRequest, userId string) *libs.Response {
	result := models.Payment{}
	userData, err := re.user_repo.FindById(userId)
	if err != nil {
		return libs.New(err.Error(), 400, true)
	}

	timeNow, err := libs.TimeNow()
	if err != nil {
		return libs.New(err.Error(), 400, true)
	}

	if (userData.Balance - data.Amount) < 0 {
		return libs.New("Balance is not enough", 400, true)

	}

	dataEntity := &models.Payment{
		PaymentId:     uuid.New().String(),
		UserId:        userId,
		AmountPayment: data.Amount,
		BalanceBefore: userData.Balance,
		BalanceAfter:  userData.Balance - data.Amount,
		Remarks:       data.Remarks,
		Status:        "PENDING",
		CreatedDate:   timeNow,
	}

	redisKey := "payment:pending:" + userId

	err = re.payment_repo.SetRedisPayment(context.Background(), redisKey, dataEntity, 0)
	if err != nil {
		return libs.New(err.Error(), 400, true)
	}
	return libs.New(result, 200, false)
}
