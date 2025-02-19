package interfaces

import (
	"go-wallet/src/libs"
	"go-wallet/src/models"
)

type PaymentService interface {
	PostPayment(data *models.PaymentRequest, userId string) *libs.Response
}
