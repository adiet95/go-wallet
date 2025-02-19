package interfaces

import (
	"go-wallet/src/libs"
)

type TransactionService interface {
	GetAllStatusTransaction(userId string) *libs.Response
}
