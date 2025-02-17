package entity

import "database/sql"

type Transfer struct {
	TransferId     string        `gorm:"primaryKey;type:uuid;default:uuid_generate_v4();" json:"transfer_id,omitempty"`
	UserId         string        `gorm:"type:uuid;not null" json:"user_id,omitempty"`
	AmountTransfer sql.NullInt64 `gorm:"type:integer;not null" json:"amount_transfer,omitempty"`
	BalanceBefore  sql.NullInt64 `gorm:"type:integer;not null" json:"balance_before,omitempty"`
	BalanceAfter   sql.NullInt64 `gorm:"type:integer;not null" json:"balance_after,omitempty"`
	CreatedDate    string        `gorm:"type:timestamp;not null" json:"created_date,omitempty"`
}

type Transfers []Transfer
