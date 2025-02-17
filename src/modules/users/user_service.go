package users

import (
	"go-wallet/src/interfaces"
	"go-wallet/src/libs"
	"go-wallet/src/models"
	"go-wallet/src/models/entity"
	"strings"
)

type user_service struct {
	user_repo interfaces.UserRepo
}

func NewService(reps interfaces.UserRepo) *user_service {
	return &user_service{reps}
}

func (re *user_service) UpdateProfile(body *models.UpdateUserRequest, phoneNumber string) *libs.Response {
	data := entity.User{}

	if body.FirstName != "" {
		data.FirstName = libs.ToNullString(body.FirstName)
	}
	if body.LastName != "" {
		data.LastName = libs.ToNullString(body.LastName)
	}
	if body.Address != "" {
		data.Address = libs.ToNullString(body.Address)
	}

	resp, err := re.user_repo.UpdateUser(&data, phoneNumber)
	if err != nil {
		if strings.ContainsAny(err.Error(), "not found") {
			return libs.New(err.Error(), 200, false)
		}
		return libs.New(err.Error(), 400, true)
	}
	return libs.New(resp, 202, false)
}

func (re *user_service) FindPhone(phoneNumber string) *libs.Response {
	data, err := re.user_repo.FindByPhone(phoneNumber)
	if err != nil {
		return libs.New(err.Error(), 400, true)
	}
	return libs.New(data, 200, false)
}

func (re *user_service) SearchName(name string) *libs.Response {
	data, err := re.user_repo.FindByName(name)
	if err != nil {
		return libs.New(err.Error(), 400, true)
	}
	return libs.New(data, 200, false)
}

func (re *user_service) GetById(id string) *libs.Response {
	data, err := re.user_repo.FindById(id)
	if err != nil {
		return libs.New(err.Error(), 400, true)
	}
	return libs.New(data, 200, false)
}
