package topup

import (
	"encoding/json"

	"github.com/labstack/echo/v4"

	"go-wallet/src/interfaces"
	"go-wallet/src/libs"
	"go-wallet/src/models"
)

type order_ctrl struct {
	svc interfaces.TopUpService
}

func NewCtrl(reps interfaces.TopUpService) *order_ctrl {
	return &order_ctrl{svc: reps}
}

func (re *order_ctrl) PostTopUp(c echo.Context) error {
	claim_user := c.Get("user_id")
	if claim_user == "" {
		return libs.New("claim user is not exist", 400, true).Send(c)
	}

	var data models.TopUpRequest
	err := json.NewDecoder(c.Request().Body).Decode(&data)
	if err != nil {
		return libs.New(err.Error(), 400, true).Send(c)
	}
	return re.svc.PostTopUp(&data, claim_user.(string)).Send(c)
}
