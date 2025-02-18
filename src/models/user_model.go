package models

type RegisterRequest struct {
	FirstName   string `json:"first_name,omitempty" validate:"required;alpha"`
	LastName    string `json:"last_name,omitempty"`
	Address     string `json:"address,omitempty" validate:"required;alphanum"`
	PhoneNumber string `json:"phone_number,omitempty" validate:"required;numeric"`
	Pin         string `json:"pin,omitempty" validate:"required;numeric"`
}

type LoginRequest struct {
	PhoneNumber string `json:"phone_number,omitempty" validate:"required;numeric"`
	Pin         string `json:"pin,omitempty" validate:"required;numeric"`
}

type UpdateUserRequest struct {
	FirstName string `json:"first_name,omitempty" validate:"required;alpha"`
	LastName  string `json:"last_name,omitempty"`
	Address   string `json:"address,omitempty" validate:"required;alphanum"`
}

type UpdateUserResponse struct {
	FirstName   string `json:"first_name,omitempty" validate:"required;alpha"`
	LastName    string `json:"last_name,omitempty"`
	Address     string `json:"address,omitempty" validate:"required;alphanum"`
	PhoneNumber string `json:"phone_number,omitempty" validate:"required;numeric"`
	UpdatedDate string `json:"updated_date,omitempty"`
}

type UserResponse struct {
	UserId      string `json:"user_id,omitempty"`
	FirstName   string `json:"first_name,omitempty"`
	LastName    string `json:"last_name,omitempty"`
	Address     string `json:"address,omitempty"`
	PhoneNumber string `json:"phone_number,omitempty"`
	Pin         string `json:"pin,omitempty"`
	Balance     int    `json:"balance,omitempty"`
	CreatedDate string `json:"created_date,omitempty"`
	UpdatedDate string `json:"updated_date,omitempty"`
	Role        string `json:"role,omitempty"`
}

type UsersResponses []UserResponse
