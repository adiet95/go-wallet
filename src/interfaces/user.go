package interfaces

import (
	"go-wallet/src/libs"
	"go-wallet/src/models"
	"go-wallet/src/models/entity"

	"gorm.io/gorm"
)

type UserRepo interface {
	UpdateUser(data *entity.User, phoneNumber string) (resp *models.UserResponse, err error)
	FindByPhone(phoneNumber string) (resp *models.UserResponse, err error)
	FindByName(name string) (resp models.UsersResponses, err error)
	FindById(id string) (resp *models.UserResponse, err error)
	InitiateTransaction() *gorm.DB
	ExecTrxUpdateBalance(tx *gorm.DB, userId string, balance int) *gorm.DB
	CommitTrx(tx *gorm.DB, query string, args ...interface{}) error
}

type UserService interface {
	FindPhone(phoneNumber string) *libs.Response
	SearchName(name string) *libs.Response
	GetById(id string) *libs.Response
	UpdateProfile(body *models.UpdateUserRequest, phoneNumber string) *libs.Response
}
