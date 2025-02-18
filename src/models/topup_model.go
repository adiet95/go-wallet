package models

type TopUpRequest struct {
	Amount  int    `json:"amount,omitempty" validate:"required;numeric"`
	Remarks string `json:"remarks,omitempty"`
}

type TopUp struct {
	TopUpId       string `gorm:"primaryKey;type:uuid;default:uuid_generate_v4();" json:"top_up_id,omitempty"`
	UserId        string `json:"user_id,omitempty"`
	AmountTopUp   int    `json:"amount_top_up,omitempty"`
	Remarks       string `json:"remarks,omitempty"`
	Status        string `json:"status,omitempty"`
	BalanceBefore int    `json:"balance_before,omitempty"`
	BalanceAfter  int    `json:"balance_after,omitempty"`
	CreatedDate   string `json:"created_date,omitempty"`
}

type TopUps []TopUp
