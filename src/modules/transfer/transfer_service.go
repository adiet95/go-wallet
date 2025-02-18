package transfer

import (
	"context"
	"encoding/json"
	"fmt"
	"go-wallet/src/constant"
	"go-wallet/src/interfaces"
	"go-wallet/src/libs"
	"go-wallet/src/models"
	"sync"
	"time"

	"github.com/gofrs/uuid/v5"
)

var (
	mutex sync.RWMutex
)

type transfer_service struct {
	redis_repo interfaces.RedisRepo
	user_repo  interfaces.UserRepo
}

func NewService(reps interfaces.RedisRepo, user_repo interfaces.UserRepo) *transfer_service {
	return &transfer_service{
		redis_repo: reps,
		user_repo:  user_repo,
	}
}

func (re *transfer_service) GetAllStatusTransfer(userId string) *libs.Response {
	redisKey := "transfer:*:" + userId
	data, err := re.redis_repo.GetRedis(context.Background(), redisKey)
	if err != nil {
		return libs.New(err.Error(), 400, true)
	}

	return libs.New(data, 200, false)
}

func (re *transfer_service) GetPendingStatusTransfer(userId string) *libs.Response {
	redisKey := "transfer:pending:" + userId
	data, err := re.redis_repo.GetRedis(context.Background(), redisKey)
	if err != nil {
		return libs.New(err.Error(), 400, true)
	}

	return libs.New(data, 200, false)
}

func (re *transfer_service) PostTransfer(data *models.TransferRequest, userId string) *libs.Response {
	userData, err := re.user_repo.FindById(userId)
	if err != nil {
		return libs.New(err.Error(), 400, true)
	}

	timeNow, err := libs.TimeNow()
	if err != nil {
		return libs.New(err.Error(), 400, true)
	}
	uuidID, _ := uuid.NewV4()

	if (userData.Balance - data.Amount) < 0 {
		return libs.New("Balance is not enough", 400, true)
	}

	dataEntity := &models.Transfer{
		TransferId:     uuidID.String(),
		UserId:         userId,
		TargetUser:     data.TargetUser,
		AmountTransfer: data.Amount,
		BalanceBefore:  userData.Balance,
		BalanceAfter:   userData.Balance - data.Amount,
		Remarks:        data.Remarks,
		Status:         "PENDING",
		CreatedDate:    timeNow,
	}

	redisKey := constant.DefaultKeyRedis + ":transfer:pending:" + dataEntity.TransferId + ":" + userId
	ObjectRedis, err := libs.StructToMap(dataEntity)
	if err != nil {
		return libs.New(err.Error(), 400, true)
	}
	err = re.redis_repo.SetRedis(context.Background(), redisKey, ObjectRedis, 0)
	if err != nil {
		return libs.New(err.Error(), 400, true)
	}
	result := &models.Transfer{
		TransferId:     dataEntity.TransferId,
		AmountTransfer: dataEntity.AmountTransfer,
		Remarks:        data.Remarks,
		BalanceBefore:  dataEntity.BalanceBefore,
		BalanceAfter:   dataEntity.BalanceAfter,
		CreatedDate:    timeNow,
	}
	return libs.New(result, 200, false)
}

func (re *transfer_service) WorkerTransfer() {
	for {
		ctx := context.Background()
		redisKey := constant.DefaultKeyRedis + ":transfer:pending:*:*"
		foundKey, _ := re.redis_repo.SearchKey(ctx, redisKey)
		if foundKey != "" {
			var redisData models.Transfer
			dataRedis, err := re.redis_repo.GetRedis(ctx, foundKey)
			if dataRedis != "" && err == nil {
				err = json.Unmarshal([]byte(dataRedis), &redisData)
				if err != nil {
					fmt.Println("error unmarshal :", err.Error())
				}
			} else {
				time.Sleep(1 * time.Second)
				continue
			}

			uuidID, _ := uuid.FromString(redisData.UserId)
			isValid := uuidID.IsNil()
			if !isValid && redisData.UserId != "" {
				mutex.Lock()

				userData, _ := re.user_repo.FindById(redisData.UserId)

				trx := re.user_repo.InitiateTransaction()

				trx = re.user_repo.ExecTrxUpdateBalance(trx, userData.UserId, redisData.AmountTransfer, "transfer")

				if trx.Error != nil {
					trx.Rollback()
				} else {
					//Add Balance to target user
					err := re.user_repo.ExecTrxTransferBalance(redisData.TargetUser, redisData.AmountTransfer, "topup")
					if err != nil {
						trx.Rollback()
					}

					redisKeySuccess := constant.DefaultKeyRedis + ":transfer:success:" + redisData.TransferId + ":" + redisData.UserId
					timeNow, _ := libs.TimeNow()
					redisData.Status = "SUCCESS"
					redisData.BalanceBefore = userData.Balance
					redisData.BalanceAfter = userData.Balance - redisData.AmountTransfer
					redisData.CreatedDate = timeNow

					ObjectRedis, err := libs.StructToMap(redisData)
					if err != nil {
						trx.Rollback()
					}
					err = re.redis_repo.SetRedis(ctx, redisKeySuccess, ObjectRedis, 0)
					if err != nil {
						trx.Rollback()
					} else {
						err := re.redis_repo.DelRedis(ctx, foundKey)
						if err != nil {
							trx.Rollback()
						}
						err = re.user_repo.CommitTrx(trx)
						if err != nil {
							trx.Rollback()
						}
					}
				}
				mutex.Unlock()
			}
		}
	}
}
