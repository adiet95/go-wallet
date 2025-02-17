package models

type TopUpRequest struct {
	Amount  int    `json:"amount,omitempty" validate:"required"`
	Remarks string `json:"remarks,omitempty"`
}

type TopUpResponse struct {
	TopUpId       string `json:"top_up_id,omitempty"`
	UserId        string `json:"user_id,omitempty"`
	AmountTopUp   int    `json:"amount_top_up,omitempty"`
	BalanceBefore int    `json:"balance_before,omitempty"`
	BalanceAfter  int    `json:"balance_after,omitempty"`
	CreatedDate   string `json:"created_date,omitempty"`
}

type TopUpsResponse []TopUpResponse
