package models

type TransferRequest struct {
	Amount  int    `json:"amount,omitempty" validate:"required;numeric"`
	Remarks string `json:"remarks,omitempty"`
}

type TransferResponse struct {
	TransferId     string `json:"transfer_id,omitempty"`
	UserId         string `json:"user_id,omitempty"`
	AmountTransfer int    `json:"amount_transfer,omitempty"`
	BalanceBefore  int    `json:"balance_before,omitempty"`
	BalanceAfter   int    `json:"balance_after,omitempty"`
	CreatedDate    string `json:"created_date,omitempty"`
}

type TransfersResponse []TransferResponse
