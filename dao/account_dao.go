package dao

import (
	"database/sql"
	"errors"
	"fmt"

	"transfer-system/models"
)

// AccountDAO handles database operations for accounts
type AccountDAO struct {
	db *sql.DB
}

// NewAccountDAO creates a new AccountDAO instance
func NewAccountDAO(db *sql.DB) *AccountDAO {
	return &AccountDAO{db: db}
}

// CreateAccount creates a new account in the database
func (dao *AccountDAO) CreateAccount(account *models.Account) error {
	query := `
		INSERT INTO accounts (account_id, balance, created_at, updated_at)
		VALUES ($1, $2, $3, $4)
	`
	_, err := dao.db.Exec(query, account.AccountID, account.Balance, account.CreatedAt, account.UpdatedAt)
	if err != nil {
		return fmt.Errorf("failed to create account: %w", err)
	}
	return nil
}

// GetAccountByID retrieves an account by its ID
func (dao *AccountDAO) GetAccountByID(accountID uint64) (*models.Account, error) {
	query := `
		SELECT account_id, balance, created_at, updated_at
		FROM accounts
		WHERE account_id = $1
	`
	account := &models.Account{}
	err := dao.db.QueryRow(query, accountID).Scan(
		&account.AccountID,
		&account.Balance,
		&account.CreatedAt,
		&account.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("account not found")
		}
		return nil, fmt.Errorf("failed to get account: %w", err)
	}
	return account, nil
}

// UpdateBalance updates the balance of an account (used within transactions)
func (dao *AccountDAO) UpdateBalance(tx *sql.Tx, accountID uint64, newBalance float64) error {
	query := `
		UPDATE accounts
		SET balance = $1, updated_at = NOW()
		WHERE account_id = $2
	`
	result, err := tx.Exec(query, newBalance, accountID)
	if err != nil {
		return fmt.Errorf("failed to update balance: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return errors.New("account not found")
	}

	return nil
}

// GetAccountByIDForUpdate retrieves an account with row-level locking (SELECT FOR UPDATE)
func (dao *AccountDAO) GetAccountByIDForUpdate(tx *sql.Tx, accountID uint64) (*models.Account, error) {
	query := `
		SELECT account_id, balance, created_at, updated_at
		FROM accounts
		WHERE account_id = $1
		FOR UPDATE
	`
	account := &models.Account{}
	err := tx.QueryRow(query, accountID).Scan(
		&account.AccountID,
		&account.Balance,
		&account.CreatedAt,
		&account.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("account not found")
		}
		return nil, fmt.Errorf("failed to get account: %w", err)
	}
	return account, nil
}
