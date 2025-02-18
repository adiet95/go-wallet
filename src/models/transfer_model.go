package models

type TransferRequest struct {
	Amount     int    `json:"amount,omitempty" validate:"required;numeric"`
	Remarks    string `json:"remarks,omitempty"`
	TargetUser string `json:"target_user,omitempty" validate:"required;"`
}

type Transfer struct {
	TransferId     string `gorm:"primaryKey;type:uuid;default:uuid_generate_v4();" json:"transfer_id,omitempty"`
	UserId         string `json:"user_id,omitempty"`
	TargetUser     string `json:"target_user"`
	AmountTransfer int    `json:"amount_transfer"`
	Remarks        string `json:"remarks,omitempty"`
	Status         string `json:"status,omitempty"`
	BalanceBefore  int    `json:"balance_before"`
	BalanceAfter   int    `json:"balance_after"`
	CreatedDate    string `json:"created_date,omitempty"`
	UpdatedDate    string `json:"updated_date,omitempty"`
}

type Transfers []Transfer
