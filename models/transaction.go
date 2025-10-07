package models

import (
	"time"
)

// Transaction represents a fund transfer between accounts
type Transaction struct {
	ID                   int64     `json:"id" db:"id"`
	SourceAccountID      uint64    `json:"source_account_id" db:"source_account_id"`
	DestinationAccountID uint64    `json:"destination_account_id" db:"destination_account_id"`
	Amount               float64   `json:"amount" db:"amount"` // storing amount as numeric with 5 precision
	Status               string    `json:"status" db:"status"` // for extensibility, can be async operation.
	CreatedAt            time.Time `json:"created_at" db:"created_at"`
}

// TransferRequest represents the request body for fund transfer
type TransferRequest struct {
	SourceAccountID      uint64 `json:"source_account_id" binding:"required"`
	DestinationAccountID uint64 `json:"destination_account_id" binding:"required"`
	Amount               string `json:"amount" binding:"required"`
}

// TransferResponse represents the response for successful transfer
type TransferResponse struct {
	TransactionID        int64  `json:"transaction_id"`
	SourceAccountID      uint64 `json:"source_account_id"`
	DestinationAccountID uint64 `json:"destination_account_id"`
	Amount               string `json:"amount"`
	Status               string `json:"status"`
}
