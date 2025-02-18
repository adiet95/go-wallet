package users

import (
	"database/sql"
	"errors"

	"go-wallet/src/models"
	"go-wallet/src/models/entity"

	"gorm.io/gorm"
)

type user_repo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) *user_repo {
	return &user_repo{db}
}

func (re *user_repo) UpdateUser(data *entity.User, phoneNumber string) (resp *models.UserResponse, err error) {
	res := re.db.Model(&data).Where("LOWER(phone_number) = ?", phoneNumber).Updates(&data)

	if res.Error != nil {
		return nil, errors.New("failed to update data")
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

func (re *user_repo) FindByPhone(phoneNumber string) (resp *models.UserResponse, err error) {
	var data *entity.User
	res := re.db.Model(&data).Where("LOWER(phone_number) = ?", phoneNumber).First(&data)
	if res.Error != nil {
		return nil, errors.New("failed to find data")
	}
	if res.RowsAffected == 0 {
		return nil, errors.New("phoneNumber not found")
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

func (re *user_repo) FindByName(name string) (resp models.UsersResponses, err error) {
	var datas entity.Users

	res := re.db.Where("LOWER(first_name) LIKE ?", "%"+name+"%").Find(&datas)
	if res.Error != nil {
		return nil, errors.New("failed to find data")
	}
	if res.RowsAffected == 0 {
		return nil, errors.New("name not found")
	}

	for _, data := range datas {
		resp = append(resp, models.UserResponse{
			UserId:      data.UserId,
			FirstName:   data.FirstName.String,
			LastName:    data.LastName.String,
			Address:     data.Address.String,
			PhoneNumber: data.PhoneNumber.String,
			CreatedDate: data.CreatedDate,
			Role:        data.Role,
		})
	}

	return resp, nil
}

func (re *user_repo) FindById(id string) (resp *models.UserResponse, err error) {
	var data *entity.User

	res := re.db.Model(&data).Where("user_id = ?", id).First(&data)
	if res.Error != nil {
		return nil, errors.New("failed to find data")
	}
	if res.RowsAffected == 0 {
		return nil, errors.New("name not found")
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

func (re *user_repo) InitiateTransaction() *gorm.DB {
	trx := re.db.Begin(&sql.TxOptions{})
	return trx
}

func (re *user_repo) ExecTrxUpdateBalance(tx *gorm.DB, userId string, balance int) *gorm.DB {
	trx := tx.Where("LOWER(user_id) = ?", userId).Update("balance", balance)
	return trx
}

func (re *user_repo) CommitTrx(tx *gorm.DB, query string, args ...interface{}) error {
	trx := tx.Commit().Error
	return trx
}
