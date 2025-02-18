package payment

import (
	"context"
	"encoding/json"
	"go-wallet/src/interfaces"
	"go-wallet/src/libs"
	"go-wallet/src/models"
	"sync"
	"time"

	"github.com/google/uuid"
)

var (
	mutex    sync.RWMutex
	firstRun = true
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

func (re *payment_service) WorkerPayment() {
	for {
		ctx := context.Background()
		redisKey := "payment:pending:*"
		var redisData models.Payment
		dataRedis, err := re.payment_repo.GetRedisPayment(ctx, redisKey)
		if dataRedis != nil && err == nil {
			dbByte, err := json.Marshal(dataRedis)
			if err != nil {

			}
			err = json.Unmarshal(dbByte, &redisData)
			if err != nil {

			}

		} else {
			time.Sleep(1 * time.Second)
			continue
		}
		mutex.Lock()
		// StaticParam = *staticParam
		userData, _ := re.user_repo.FindById(redisData.UserId)

		trx := re.user_repo.InitiateTransaction()

		trx = re.user_repo.ExecTrxUpdateBalance(trx, userData.UserId, redisData.BalanceAfter)

		if trx.Error != nil {
			trx.Rollback()
		} else {
			timeNow, _ := libs.TimeNow()
			redisData.Status = "SUCCESS"
			redisData.BalanceBefore = userData.Balance
			redisData.BalanceAfter = userData.Balance - redisData.AmountPayment
			redisData.CreatedDate = timeNow
			redisKey = "payment:success:" + redisData.UserId
			err = re.payment_repo.SetRedisPayment(ctx, redisKey, &redisData, 0)
			if err != nil {
				trx.Rollback()
			} else {
				err := re.payment_repo.DelRedisPayment(ctx, "payment:pending:"+redisData.UserId)
				if err != nil {
					trx.Rollback()
				}
			}
		}
		mutex.Unlock()

		// firstRun = false
		// time.Sleep(1 * time.Second)
	}
}
