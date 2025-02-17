package interfaces

import (
	"go-wallet/src/libs"
	"go-wallet/src/models"
	"go-wallet/src/models/entity"
)

type UserRepo interface {
	UpdateUser(data *entity.User, phoneNumber string) (resp *models.UserResponse, err error)
	FindByPhone(phoneNumber string) (resp *models.UserResponse, err error)
	FindByName(name string) (resp models.UsersResponses, err error)
	FindById(id string) (resp *models.UserResponse, err error)
}

type UserService interface {
	FindPhone(phoneNumber string) *libs.Response
	SearchName(name string) *libs.Response
	GetById(id string) *libs.Response
	UpdateProfile(body *models.UpdateUserRequest, phoneNumber string) *libs.Response
}
