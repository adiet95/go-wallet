package transaction

import (
	"context"
	"encoding/json"
	"go-wallet/src/constant"
	"go-wallet/src/interfaces"
	"go-wallet/src/libs"
)

type transaction_service struct {
	redis_repo interfaces.RedisRepo
	user_repo  interfaces.UserRepo
}

func NewService(reps interfaces.RedisRepo, user_repo interfaces.UserRepo) *transaction_service {
	return &transaction_service{
		redis_repo: reps,
		user_repo:  user_repo,
	}
}

func (re *transaction_service) GetAllStatusTransaction(userId string) *libs.Response {
	ctx := context.Background()
	datas := []interface{}{}

	redisKey := constant.DefaultKeyRedis + ":*:*:*:" + userId
	foundKey, _ := re.redis_repo.SearchKeyArr(ctx, redisKey)
	if len(foundKey) != 0 {
		for _, valueKey := range foundKey {
			result := make(map[string]interface{})

			dataRedis, err := re.redis_repo.GetRedis(ctx, valueKey)
			if dataRedis != "" && err == nil {
				err = json.Unmarshal([]byte(dataRedis), &result)
				if err != nil {
					return libs.New(constant.RedisGetError, 400, true)
				}
			}
			datas = append(datas, result)
		}
	}
	return libs.New(datas, 200, false)
}

func (re *transaction_service) AdminGetAllStatusTransaction() *libs.Response {
	ctx := context.Background()
	datas := []interface{}{}

	redisKey := constant.DefaultKeyRedis + ":*:*:*:*"
	foundKey, _ := re.redis_repo.SearchKeyArr(ctx, redisKey)
	if len(foundKey) != 0 {
		for _, valueKey := range foundKey {
			result := make(map[string]interface{})

			dataRedis, err := re.redis_repo.GetRedis(ctx, valueKey)
			if dataRedis != "" && err == nil {
				err = json.Unmarshal([]byte(dataRedis), &result)
				if err != nil {
					return libs.New(constant.RedisGetError, 400, true)
				}
			}
			datas = append(datas, result)
		}
	}
	return libs.New(datas, 200, false)
}
