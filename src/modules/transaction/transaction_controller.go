package transaction

import (
	"github.com/labstack/echo/v4"

	"go-wallet/src/interfaces"
	"go-wallet/src/libs"
)

type transaction_ctrl struct {
	svc interfaces.TransactionService
}

func NewCtrl(reps interfaces.TransactionService) *transaction_ctrl {
	return &transaction_ctrl{svc: reps}
}

func (re *transaction_ctrl) GetAllTransaction(c echo.Context) error {
	claim_user := c.Get("user_id")
	if claim_user == "" {
		return libs.New("claim user is not exist", 400, true).Send(c)
	}
	return re.svc.GetAllStatusTransaction(claim_user.(string)).Send(c)
}

func (re *transaction_ctrl) AdminGetAllTransaction(c echo.Context) error {
	return re.svc.AdminGetAllStatusTransaction().Send(c)
}
