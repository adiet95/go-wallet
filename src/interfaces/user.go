package interfaces

import (
	"go-wallet/src/libs"
	"go-wallet/src/models"
	"go-wallet/src/models/entity"

	"gorm.io/gorm"
)

type UserRepo interface {
	UpdateUser(data *entity.User, userId string) (resp *models.UserResponse, err error)
	FindByName(name string) (resp models.UsersResponses, err error)
	FindById(id string) (resp *models.UserResponse, err error)
	InitiateTransaction() *gorm.DB
	ExecTrxUpdateBalance(tx *gorm.DB, userId string, amount int, typeTrx string) *gorm.DB
	ExecTrxTransferBalance(userId string, amount int, typeTrx string) error
	CommitTrx(tx *gorm.DB) error
}

type UserService interface {
	SearchName(name string) *libs.Response
	GetById(id string) *libs.Response
	UpdateProfile(body *models.UpdateUserRequest, phoneNumber string) *libs.Response
}
