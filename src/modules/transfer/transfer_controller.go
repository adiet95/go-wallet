package transfer

import (
	"encoding/json"

	"github.com/labstack/echo/v4"

	"go-wallet/src/interfaces"
	"go-wallet/src/libs"
	"go-wallet/src/models"
)

type transfer_ctrl struct {
	svc interfaces.TransferService
}

func NewCtrl(reps interfaces.TransferService) *transfer_ctrl {
	return &transfer_ctrl{svc: reps}
}

func (re *transfer_ctrl) PostPayment(c echo.Context) error {
	claim_user := c.Get("user_id")
	if claim_user == "" {
		return libs.New("claim user is not exist", 400, true).Send(c)
	}

	var data models.TransferRequest
	err := json.NewDecoder(c.Request().Body).Decode(&data)
	if err != nil {
		return libs.New(err.Error(), 400, true).Send(c)
	}
	return re.svc.PostTransfer(&data, claim_user.(string)).Send(c)
}
