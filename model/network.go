package model

import (
	"gorm.io/gorm"
)

type Network struct {
	gorm.Model
	ExternalId          string
	Name                string
	Status              string
	MinWithdrawalAmount float64
	MaxWithdrawalAmount float64
}

type NetworkIn struct {
	ExternalId          string  `json:"id"`
	Name                string  `json:"name"`
	Status              string  `json:"status"`
	MinWithdrawalAmount float64 `json:"min_withdrawal_amount"`
	MaxWithdrawalAmount float64 `json:"max_withdrawal_amount"`
}
