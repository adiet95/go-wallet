package interfaces

import (
	"go-wallet/src/libs"
	"go-wallet/src/models"
	"go-wallet/src/models/entity"
)

type AuthRepo interface {
	FindByPhone(phoneNumber string) (resp *models.UserResponse, err error)
	RegisterPhone(data *entity.User) (resp *models.UserResponse, err error)
}

type AuthService interface {
	Login(body models.LoginRequest) *libs.Response
	Register(body *models.RegisterRequest) *libs.Response
}
