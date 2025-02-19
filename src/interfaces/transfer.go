package interfaces

import (
	"go-wallet/src/libs"
	"go-wallet/src/models"
)

type TransferService interface {
	PostTransfer(data *models.TransferRequest, userId string) *libs.Response
}
