package models

import (
	"gorm.io/gorm"
)

type TransactionType string

const (
	Credit TransactionType = "credit"
	Debit  TransactionType = "debit"
)

type TransactionCurrency string

const (
	USD TransactionCurrency = "usd"
	EUR TransactionCurrency = "eur"
)

type Transaction struct {
	gorm.Model
	UserID   uint
	Amount   int64
	Currency TransactionCurrency
	Type     TransactionType
}
