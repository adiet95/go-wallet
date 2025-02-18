package topup

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
	mutex sync.RWMutex
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

func (re *topup_service) WorkerTopUp() {
	for {
		ctx := context.Background()
		redisKey := "topup:pending:*"
		var redisData models.TopUp
		dataRedis, err := re.topup_repo.GetRedisTopUp(ctx, redisKey)
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
			redisData.BalanceAfter = userData.Balance - redisData.AmountTopUp
			redisData.CreatedDate = timeNow
			redisKey = "topup:success:" + redisData.UserId
			err = re.topup_repo.SetRedisTopUp(ctx, redisKey, &redisData, 0)
			if err != nil {
				trx.Rollback()
			} else {
				err := re.topup_repo.DelRedisPayment(ctx, "topup:pending:"+redisData.UserId)
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
