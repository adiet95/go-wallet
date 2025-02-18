package users

import (
	"database/sql"
	"errors"

	"go-wallet/src/libs"
	"go-wallet/src/models"
	"go-wallet/src/models/entity"

	"github.com/gofrs/uuid/v5"
	"gorm.io/gorm"
)

type user_repo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) *user_repo {
	return &user_repo{db}
}

func (re *user_repo) UpdateUser(data *entity.User, userId string) (resp *models.UserResponse, err error) {
	uuidID, _ := uuid.FromString(userId)
	res := re.db.Model(&data).Where("user_id = ?", uuidID).Updates(&data)

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
	uuidID, _ := uuid.FromString(id)

	res := re.db.Model(data).Where("user_id = ?", uuidID).Find(&data)
	if res.Error != nil {
		return nil, errors.New("failed to find data")
	}
	if res.RowsAffected == 0 {
		return nil, errors.New("name not found")
	}
	resp = &models.UserResponse{
		UserId:      data.UserId,
		Balance:     int(data.Balance.Int64),
		FirstName:   data.FirstName.String,
		LastName:    data.LastName.String,
		Address:     data.Address.String,
		PhoneNumber: data.PhoneNumber.String,
		CreatedDate: data.CreatedDate,
		UpdatedDate: data.UpdatedDate,
	}

	return resp, nil
}

func (re *user_repo) InitiateTransaction() *gorm.DB {
	trx := re.db.Begin(&sql.TxOptions{})
	return trx
}

func (re *user_repo) ExecTrxUpdateBalance(tx *gorm.DB, userId string, amount int, typeTrx string) *gorm.DB {
	data := &entity.User{}
	var finalAmount int64
	uuidID, _ := uuid.FromString(userId)

	tx.Where("user_id = ?", uuidID).Find(data)

	switch typeTrx {
	case "topup":
		finalAmount = data.Balance.Int64 + int64(amount)
	case "payment":
		finalAmount = data.Balance.Int64 - int64(amount)
	case "transfer":
		finalAmount = data.Balance.Int64 - int64(amount)
	}
	data.Balance = libs.ToNullInt64(finalAmount)
	trx := tx.Where("user_id = ?", uuidID).Updates(data)
	return trx
}

func (re *user_repo) ExecTrxTransferBalance(userId string, amount int, typeTrx string) error {
	data := &entity.User{}
	var finalAmount int64
	uuidID, _ := uuid.FromString(userId)

	re.db.Where("user_id = ?", uuidID).Find(data)

	switch typeTrx {
	case "topup":
		finalAmount = data.Balance.Int64 + int64(amount)
	case "payment":
		finalAmount = data.Balance.Int64 - int64(amount)
	case "transfer":
		finalAmount = data.Balance.Int64 - int64(amount)
	}
	data.Balance = libs.ToNullInt64(finalAmount)
	err := re.db.Where("user_id = ?", uuidID).Updates(data)
	if err.Error != nil {
		return errors.New("Error update balance")
	}
	return nil
}

func (re *user_repo) CommitTrx(tx *gorm.DB) error {
	trx := tx.Commit().Error
	return trx
}
