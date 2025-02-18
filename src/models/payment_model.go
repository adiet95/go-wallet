package models

type PaymentRequest struct {
	Amount  int    `json:"amount,omitempty" validate:"required;numeric"`
	Remarks string `json:"remarks,omitempty"`
}

type Payment struct {
	PaymentId     string `gorm:"primaryKey;type:uuid;default:uuid_generate_v4();" json:"payment_id,omitempty"`
	UserId        string `json:"user_id,omitempty"`
	AmountPayment int    `json:"amount_payment,omitempty"`
	Remarks       string `json:"remarks,omitempty"`
	Status        string `json:"status,omitempty"`
	BalanceBefore int    `json:"balance_before,omitempty"`
	BalanceAfter  int    `json:"balance_after,omitempty"`
	CreatedDate   string `json:"created_date,omitempty"`
}

type Payments []Payment
