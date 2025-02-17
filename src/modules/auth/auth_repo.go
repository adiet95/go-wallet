package auth

import (
	"errors"

	"go-wallet/src/models"
	"go-wallet/src/models/entity"

	"gorm.io/gorm"
)

type auth_repo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) *auth_repo {
	return &auth_repo{db}
}

func (re *auth_repo) FindByPhone(phoneNumber string) (resp *models.UserResponse, err error) {
	var data *entity.User

	res := re.db.Where("phone_number = ?", phoneNumber).First(&data)
	if res.Error != nil {
		return nil, errors.New("failed to find data")
	}
	if res.RowsAffected == 0 {
		return nil, errors.New("Phone number not found")
	}

	resp = &models.UserResponse{
		UserId:      data.UserId,
		FirstName:   data.FirstName.String,
		LastName:    data.LastName.String,
		Address:     data.Address.String,
		PhoneNumber: data.PhoneNumber.String,
		Role:        data.Role,
		CreatedDate: data.CreatedDate,
	}

	return resp, nil
}

func (re *auth_repo) RegisterPhone(data *entity.User) (resp *models.UserResponse, err error) {
	var datas *entity.Users

	res := re.db.Model(&datas).Where("phone_number = ?", data.PhoneNumber).First(&data)
	if res.Error != nil {
		return nil, errors.New("failed to find data")
	}
	if res.RowsAffected > 0 {
		return nil, errors.New("Phone number already registered")
	}

	r := res.Create(data)
	if r.Error != nil {
		return nil, errors.New("failed to save data")
	}
	resp = &models.UserResponse{
		UserId:      data.UserId,
		FirstName:   data.FirstName.String,
		LastName:    data.LastName.String,
		Address:     data.Address.String,
		PhoneNumber: data.PhoneNumber.String,
		CreatedDate: data.CreatedDate,
	}

	return resp, nil
}
