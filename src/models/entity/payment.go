package entity

import "database/sql"

type Payment struct {
	PaymentId     string        `gorm:"primaryKey;type:uuid;default:uuid_generate_v4();" json:"payment_id,omitempty"`
	UserId        string        `gorm:"type:uuid;not null" json:"user_id,omitempty"`
	AmountPayment sql.NullInt64 `gorm:"type:integer;not null" json:"amount_payment,omitempty"`
	BalanceBefore sql.NullInt64 `gorm:"type:integer;not null" json:"balance_before,omitempty"`
	BalanceAfter  sql.NullInt64 `gorm:"type:integer;not null" json:"balance_after,omitempty"`
	CreatedDate   string        `gorm:"type:timestamp;not null" json:"created_date,omitempty"`
}

type Payments []Payment
