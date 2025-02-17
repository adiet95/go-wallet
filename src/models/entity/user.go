package entity

import "database/sql"

type User struct {
	UserId      string         `gorm:"primaryKey;type:uuid;default:uuid_generate_v4();" json:"user_id,omitempty"`
	FirstName   sql.NullString `json:"first_name,omitempty" validate:"required"`
	LastName    sql.NullString `json:"last_name,omitempty"`
	Address     sql.NullString `json:"address,omitempty" validate:"required"`
	PhoneNumber sql.NullString `json:"phone_number,omitempty" validate:"required"`
	Pin         sql.NullString `json:"pin,omitempty" validate:"required"`
	Balance     sql.NullInt64  `json:"balance,omitempty"`
	Role        string         `json:"role,omitempty"`
	CreatedDate string         `json:"created_date,omitempty"`
	UpdatedDate string         `json:"updated_date,omitempty"`
}

type Users []User
