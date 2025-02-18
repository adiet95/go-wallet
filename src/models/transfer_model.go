package models

type TransferRequest struct {
	Amount  int    `json:"amount,omitempty" validate:"required;numeric"`
	Remarks string `json:"remarks,omitempty"`
}

type Transfer struct {
	TransferId     string `gorm:"primaryKey;type:uuid;default:uuid_generate_v4();" json:"transfer_id,omitempty"`
	UserId         string `json:"user_id,omitempty"`
	TargetUserId   string `json:"target_user_id,omitempty"`
	AmountTransfer int    `json:"amount_transfer,omitempty"`
	Remarks        string `json:"remarks,omitempty"`
	Status         string `json:"status,omitempty"`
	BalanceBefore  int    `json:"balance_before,omitempty"`
	BalanceAfter   int    `json:"balance_after,omitempty"`
	CreatedDate    string `json:"created_date,omitempty"`
	UpdatedDate    string `json:"updated_date,omitempty"`
}

type Transfers []Transfer
