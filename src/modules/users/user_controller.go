package users

import (
	"encoding/json"
	"strings"

	"github.com/labstack/echo/v4"

	"go-wallet/src/interfaces"
	"go-wallet/src/libs"
	"go-wallet/src/models"
)

type user_ctrl struct {
	svc interfaces.UserService
}

func NewCtrl(reps interfaces.UserService) *user_ctrl {
	return &user_ctrl{svc: reps}
}

func (re *user_ctrl) UpdateProfile(c echo.Context) error {
	claim_user := c.Get("phone_number")
	if claim_user == "" {
		return libs.New("claim user is not exist", 400, true).Send(c)
	}

	var data models.UpdateUserRequest
	err := json.NewDecoder(c.Request().Body).Decode(&data)
	if err != nil {
		return libs.New(err.Error(), 400, true).Send(c)
	}

	return re.svc.UpdateProfile(&data, claim_user.(string)).Send(c)
}

func (re *user_ctrl) SearchPhone(c echo.Context) error {
	val := c.QueryParam("phone_number")

	return re.svc.FindPhone(val).Send(c)
}

func (re *user_ctrl) SearchName(c echo.Context) error {
	val := c.QueryParam("first_name")
	v := strings.ToLower(val)
	return re.svc.SearchName(v).Send(c)
}

func (re *user_ctrl) SearchId(c echo.Context) error {
	val := c.Param("id")
	return re.svc.GetById(val).Send(c)
}
