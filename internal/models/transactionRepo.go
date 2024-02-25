package models

import (
	"gorm.io/gorm"
)

type TransactionType string

const (
	Credit TransactionType = "credit"
	Debit  TransactionType = "debit"
)

type Transaction struct {
	gorm.Model
	UserID uint
	Amount int64
	Type   TransactionType
}
