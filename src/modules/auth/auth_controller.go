package auth

import (
	"encoding/json"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"

	"go-wallet/src/interfaces"
	"go-wallet/src/libs"
	"go-wallet/src/models"
)

type user_ctrl struct {
	repo     interfaces.AuthService
	Validate *validator.Validate
}

func NewCtrl(reps interfaces.AuthService, Validate *validator.Validate) *user_ctrl {
	return &user_ctrl{
		repo:     reps,
		Validate: Validate,
	}
}

func (u user_ctrl) SignIn(c echo.Context) error {
	var data models.LoginRequest

	err := json.NewDecoder(c.Request().Body).Decode(&data)
	if err != nil {
		return libs.New(err.Error(), 401, true).Send(c)
	}

	err = u.Validate.Struct(data)
	if err != nil {
		return libs.New(err.Error(), 400, true).Send(c)
	}

	return u.repo.Login(data).Send(c)
}

func (u user_ctrl) Register(c echo.Context) error {
	var data *models.RegisterRequest

	err := json.NewDecoder(c.Request().Body).Decode(&data)
	if err != nil {
		return libs.New(err.Error(), 401, true).Send(c)
	}

	err = u.Validate.Struct(data)
	if err != nil {
		return libs.New(err.Error(), 400, true).Send(c)
	}

	return u.repo.Register(data).Send(c)
}
