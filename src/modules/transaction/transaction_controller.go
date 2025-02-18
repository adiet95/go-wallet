package transaction

import (
	"encoding/json"

	"github.com/labstack/echo/v4"

	"go-wallet/src/interfaces"
	"go-wallet/src/libs"
	"go-wallet/src/models"
)

type transaction_ctrl struct {
	svc interfaces.TransactionService
}

func NewCtrl(reps interfaces.TransactionService) *transaction_ctrl {
	return &transaction_ctrl{svc: reps}
}

func (re *transaction_ctrl) PostPayment(c echo.Context) error {
	claim_user := c.Get("user_id")
	if claim_user == "" {
		return libs.New("claim user is not exist", 400, true).Send(c)
	}

	var data models.TransactionRequest
	err := json.NewDecoder(c.Request().Body).Decode(&data)
	if err != nil {
		return libs.New(err.Error(), 400, true).Send(c)
	}
	return re.svc.PostTransaction(&data, claim_user.(string)).Send(c)
}
