package order

import (
	"encoding/json"

	"github.com/labstack/echo/v4"

	"go-wallet/src/interfaces"
	"go-wallet/src/libs"
	"go-wallet/src/models"
)

type order_ctrl struct {
	svc interfaces.PaymentService
}

func NewCtrl(reps interfaces.PaymentService) *order_ctrl {
	return &order_ctrl{svc: reps}
}

func (re *order_ctrl) PostPayment(c echo.Context) error {
	claim_user := c.Get("user_id")
	if claim_user == "" {
		return libs.New("claim user is not exist", 400, true).Send(c)
	}

	var data models.PaymentRequest
	err := json.NewDecoder(c.Request().Body).Decode(&data)
	if err != nil {
		return libs.New(err.Error(), 400, true).Send(c)
	}
	return re.svc.PostPayment(&data, claim_user.(string)).Send(c)
}
