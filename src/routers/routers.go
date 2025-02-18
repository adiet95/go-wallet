package routers

import (
	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"

	auth "go-wallet/src/modules/auth"
	"go-wallet/src/modules/payment"
	"go-wallet/src/modules/topup"
	"go-wallet/src/modules/transaction"
	"go-wallet/src/modules/transfer"
	"go-wallet/src/modules/users"
)

func New(mainRoute *echo.Echo, db *gorm.DB, rd *redis.Client) (*echo.Echo, error) {
	auth.New(mainRoute, db)
	users.New(mainRoute, db)
	payment.New(mainRoute, db, rd)
	topup.New(mainRoute, db, rd)
	transfer.New(mainRoute, db, rd)
	transaction.New(mainRoute, db, rd)

	return mainRoute, nil
}
