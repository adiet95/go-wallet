package auth

import (
	"go-wallet/src/interfaces"
	"go-wallet/src/libs"
	"go-wallet/src/models"
	"go-wallet/src/models/entity"
)

type auth_service struct {
	repo interfaces.AuthRepo
}
type token_response struct {
	Tokens       string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func NewService(reps interfaces.AuthRepo) *auth_service {
	return &auth_service{reps}
}

func (u auth_service) Login(body models.LoginRequest) *libs.Response {
	user, err := u.repo.FindByPhone(body.PhoneNumber)
	if err != nil {
		return libs.New(err.Error(), 400, true)
	}
	if !libs.CheckPass(user.Pin, body.Pin) {
		return libs.New("Phone Number and PIN doesn't match!", 400, true)
	}
	token := libs.NewToken(user.UserId, user.Role)
	theToken, err := token.Create()
	if err != nil {
		return libs.New(err.Error(), 400, true)
	}
	refreshToken := libs.NewRefreshToken(user.UserId, user.Role)
	refToken, err := refreshToken.Create()
	if err != nil {
		return libs.New(err.Error(), 400, true)
	}
	return libs.New(token_response{Tokens: theToken, RefreshToken: refToken}, 200, false)
}

func (u auth_service) Register(body *models.RegisterRequest) *libs.Response {
	hassPass, err := libs.HashPassword(body.Pin)
	if err != nil {
		return libs.New(err.Error(), 400, true)
	}

	timeNow, err := libs.TimeNow()
	if err != nil {
		return libs.New(err.Error(), 400, true)
	}

	data := &entity.User{
		FirstName:   libs.ToNullString(body.FirstName),
		LastName:    libs.ToNullString(body.LastName),
		Address:     libs.ToNullString(body.Address),
		PhoneNumber: libs.ToNullString(body.PhoneNumber),
		Pin:         libs.ToNullString(hassPass),
		CreatedDate: timeNow,
		Role:        "user",
	}

	result, err := u.repo.RegisterPhone(data)
	if err != nil {
		return libs.New(err.Error(), 400, true)
	}
	return libs.New(result, 200, false)
}
