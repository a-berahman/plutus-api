package models

type TransactionCreateRequest struct {
	UserID   uint   `json:"userId" validate:"required"`
	Amount   string `json:"amount" validate:"required"`
	Currency string `json:"currency" validate:"required,oneof=usd eur"`
	Type     string `json:"type" validate:"required,oneof=credit debit"`
}

type TransactionUpdateRequest struct {
	Amount   string `json:"amount"`
	Currency string `json:"currency" validate:"omitempty,oneof=usd eur"`
	Type     string `json:"type" validate:"omitempty,oneof=credit debit"`
}

type TransactionResponse struct {
	ID       uint   `json:"id,omitempty"`
	UserID   uint   `json:"userId,omitempty"`
	Amount   string `json:"amount,omitempty"`
	Currency string `json:"currency,omitempty"`
	Type     string `json:"type,omitempty"`
}
