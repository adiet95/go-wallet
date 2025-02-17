package models

type PaymentRequest struct {
	Amount  int    `json:"amount,omitempty" validate:"required;numeric"`
	Remarks string `json:"remarks,omitempty"`
}

type PaymentResponse struct {
	PaymentId     string `json:"payment_id,omitempty"`
	UserId        string `json:"user_id,omitempty"`
	AmountPayment int    `json:"amount_payment,omitempty"`
	BalanceBefore int    `json:"balance_before,omitempty"`
	BalanceAfter  int    `json:"balance_after,omitempty"`
	CreatedDate   string `json:"created_date,omitempty"`
}

type PaymentsResponse []PaymentResponse
