package models

import "math/big"

type Transaction struct {
	Data BigA `json:"data,omitempty"`
}

type BigA struct{ *big.Int }
