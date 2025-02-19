package topup

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

type topup_service struct {
	user_repo  interfaces.UserRepo
	redis_repo interfaces.RedisRepo
}

func NewService(user_repo interfaces.UserRepo, redis_repo interfaces.RedisRepo) *topup_service {
	return &topup_service{
		user_repo:  user_repo,
		redis_repo: redis_repo,
	}
}

func (re *topup_service) PostTopUp(data *models.TopUpRequest, userId string) *libs.Response {
	userData, err := re.user_repo.FindById(userId)
	if err != nil {
		return libs.New(err.Error(), 400, true)
	}

	timeNow, err := libs.TimeNow()
	if err != nil {
		return libs.New(err.Error(), 400, true)
	}

	uuidID, _ := uuid.NewV4()
	dataEntity := &models.TopUp{
		TopUpId:       uuidID.String(),
		UserId:        userId,
		AmountTopUp:   data.Amount,
		BalanceBefore: userData.Balance,
		BalanceAfter:  userData.Balance + data.Amount,
		Remarks:       data.Remarks,
		Status:        "PENDING",
		CreatedDate:   timeNow,
	}

	redisKey := constant.DefaultKeyRedis + ":topup:pending:" + dataEntity.TopUpId + ":" + userId

	ObjectRedis, err := libs.StructToMap(dataEntity)
	if err != nil {
		return libs.New(err.Error(), 400, true)
	}
	err = re.redis_repo.SetRedis(context.Background(), redisKey, ObjectRedis, 0)
	if err != nil {
		return libs.New(err.Error(), 400, true)
	}
	result := &models.TopUp{
		TopUpId:       dataEntity.TopUpId,
		AmountTopUp:   dataEntity.AmountTopUp,
		BalanceBefore: dataEntity.BalanceBefore,
		BalanceAfter:  dataEntity.BalanceAfter,
		CreatedDate:   timeNow,
	}
	return libs.New(result, 200, false)
}

func (re *topup_service) WorkerTopUp() {
	for {
		ctx := context.Background()
		redisKey := constant.DefaultKeyRedis + ":topup:pending:*:*"
		foundKey, _ := re.redis_repo.SearchKey(ctx, redisKey)
		if foundKey != "" {
			var redisData models.TopUp
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

				trx = re.user_repo.ExecTrxUpdateBalance(trx, userData.UserId, redisData.AmountTopUp, "topup")

				if trx.Error != nil {
					trx.Rollback()
				} else {
					redisKeySuccess := constant.DefaultKeyRedis + ":topup:success:" + redisData.TopUpId + ":" + redisData.UserId
					timeNow, _ := libs.TimeNow()
					redisData.Status = "SUCCESS"
					redisData.BalanceBefore = userData.Balance
					redisData.BalanceAfter = userData.Balance + redisData.AmountTopUp
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
