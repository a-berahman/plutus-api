package models

type TransactionCreateRequest struct {
	UserID uint   `json:"userId" validate:"required"`
	Amount int64  `json:"amount" validate:"required"`
	Type   string `json:"type" validate:"required,oneof=credit debit"`
}

type TransactionUpdateRequest struct {
	Amount int64  `json:"amount" validate:"required"`
	Type   string `json:"type" validate:"required,oneof=credit debit"`
}

type TransactionResponse struct {
	ID     uint   `json:"id,omitempty"`
	UserID uint   `json:"userId,omitempty"`
	Amount int64  `json:"amount,omitempty"`
	Type   string `json:"type,omitempty"`
}
