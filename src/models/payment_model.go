package models

type PaymentRequest struct {
	Amount  int    `json:"amount,omitempty" validate:"required;numeric"`
	Remarks string `json:"remarks,omitempty"`
}

type Payment struct {
	PaymentId     string `gorm:"primaryKey;type:uuid;default:uuid_generate_v4();" json:"payment_id"`
	UserId        string `json:"user_id,omitempty"`
	AmountPayment int    `json:"amount_payment"`
	Remarks       string `json:"remarks,omitempty"`
	Status        string `json:"status,omitempty"`
	BalanceBefore int    `json:"balance_before"`
	BalanceAfter  int    `json:"balance_after"`
	CreatedDate   string `json:"created_date,omitempty"`
	UpdatedDate   string `json:"updated_date,omitempty"`
}

type Payments []Payment
