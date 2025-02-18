package transfer

import (
	"context"
	"encoding/json"
	"fmt"
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
	uuidID, _ := uuid.NewRandom()

	dataEntity := &models.Transfer{
		TransferId:     uuidID.String(),
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

func (re *transfer_service) WorkerTransfer() {
	for {
		ctx := context.Background()
		redisKey := "topup:pending:*"
		var redisData models.Transfer
		dataRedis, err := re.transfer_repo.GetRedisTransfer(ctx, redisKey)
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

		fmt.Println(redisData, "<<<REDIS DATA")
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
			redisData.BalanceAfter = userData.Balance - redisData.AmountTransfer
			redisData.CreatedDate = timeNow
			redisKey = "topup:success:" + redisData.UserId
			err = re.transfer_repo.SetRedisTransfer(ctx, redisKey, &redisData, 0)
			if err != nil {
				trx.Rollback()
			} else {
				err := re.transfer_repo.DelRedisPayment(ctx, "topup:pending:"+redisData.UserId)
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
