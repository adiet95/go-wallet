package payment

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
	mutex    sync.RWMutex
	firstRun = true
)

type payment_service struct {
	redis_repo interfaces.RedisRepo
	user_repo  interfaces.UserRepo
}

func NewService(reps interfaces.RedisRepo, user_repo interfaces.UserRepo) *payment_service {
	return &payment_service{
		redis_repo: reps,
		user_repo:  user_repo,
	}
}

func (re *payment_service) PostPayment(data *models.PaymentRequest, userId string) *libs.Response {
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

	uuidID, _ := uuid.NewV4()
	dataEntity := &models.Payment{
		PaymentId:     uuidID.String(),
		UserId:        userId,
		AmountPayment: data.Amount,
		BalanceBefore: userData.Balance,
		BalanceAfter:  userData.Balance - data.Amount,
		Remarks:       data.Remarks,
		Status:        "PENDING",
		CreatedDate:   timeNow,
	}

	redisKey := constant.DefaultKeyRedis + ":payment:pending:" + dataEntity.PaymentId + ":" + userId
	ObjectRedis, err := libs.StructToMap(dataEntity)
	if err != nil {
		return libs.New(err.Error(), 400, true)
	}
	err = re.redis_repo.SetRedis(context.Background(), redisKey, ObjectRedis, 0)
	if err != nil {
		return libs.New(err.Error(), 400, true)
	}
	result := &models.Payment{
		PaymentId:     dataEntity.PaymentId,
		AmountPayment: dataEntity.AmountPayment,
		Remarks:       data.Remarks,
		BalanceBefore: dataEntity.BalanceBefore,
		BalanceAfter:  dataEntity.BalanceAfter,
		CreatedDate:   timeNow,
	}
	return libs.New(result, 200, false)
}

func (re *payment_service) WorkerPayment() {
	for {
		ctx := context.Background()
		redisKey := constant.DefaultKeyRedis + ":payment:pending:*:*"
		foundKey, _ := re.redis_repo.SearchKey(ctx, redisKey)
		if foundKey != "" {
			var redisData models.Payment
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

				trx = re.user_repo.ExecTrxUpdateBalance(trx, userData.UserId, redisData.AmountPayment, "payment")

				if trx.Error != nil {
					trx.Rollback()
				} else {
					redisKeySuccess := constant.DefaultKeyRedis + ":payment:success:" + redisData.PaymentId + ":" + redisData.UserId
					timeNow, _ := libs.TimeNow()
					redisData.Status = "SUCCESS"
					redisData.BalanceBefore = userData.Balance
					redisData.BalanceAfter = userData.Balance - redisData.AmountPayment
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
					fmt.Println("<<SUCCESS")
				}
				mutex.Unlock()
			}
		}
	}
}
