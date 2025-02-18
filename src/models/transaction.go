package models

type Transaction struct {
	TransactionId string `gorm:"primaryKey;type:uuid;default:uuid_generate_v4();" json:"transaction_id,omitempty"`
	UserId        string `json:"user_id,omitempty"`
	TransferId    string `json:"transfer_id,omitempty"`
	TopUpId       string `json:"top_up_id,omitempty"`
	PaymentId     string `json:"payment_id,omitempty"`
}

type Transactions []Transaction
