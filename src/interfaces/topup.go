package interfaces

import (
	"go-wallet/src/libs"
	"go-wallet/src/models"
)

type TopUpService interface {
	PostTopUp(data *models.TopUpRequest, userId string) *libs.Response
}
