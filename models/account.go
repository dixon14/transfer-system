package models

import (
	"time"
)

// Account represents a bank account
type Account struct {
	AccountID uint64    `json:"account_id" db:"account_id"`
	Balance   float64   `json:"balance" db:"balance"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// CreateAccountRequest represents the request body for account creation
type CreateAccountRequest struct {
	AccountID      uint64 `json:"account_id" binding:"required,gt=0"`
	InitialBalance string `json:"initial_balance" binding:"required"`
}

// AccountResponse represents the response for account queries
type AccountResponse struct {
	AccountID uint64 `json:"account_id"`
	Balance   string `json:"balance"`
}
