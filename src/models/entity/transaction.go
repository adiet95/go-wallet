package entity

type Transaction struct {
	PaymentData  Payments  `json:"payment_data,omitempty"`
	TopUpData    TopUps    `json:"topup_data,omitempty"`
	TransferData Transfers `json:"transfer_data,omitempty"`
}

type Transactions []Transaction
