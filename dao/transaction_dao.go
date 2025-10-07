package dao

import (
	"database/sql"

	"transfer-system/models"
)

// TransactionDAO handles database operations for transactions
type TransactionDAO struct {
	db *sql.DB
}

// NewTransactionDAO creates a new TransactionDAO instance
func NewTransactionDAO(db *sql.DB) *TransactionDAO {
	return &TransactionDAO{db: db}
}

// CreateTransaction creates a new transaction record
func (dao *TransactionDAO) CreateTransaction(tx *sql.Tx, transaction *models.Transaction) (int64, error) {
	query := `
		INSERT INTO transactions (source_account_id, destination_account_id, amount, status, created_at)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id
	`
	var id int64
	err := tx.QueryRow(
		query,
		transaction.SourceAccountID,
		transaction.DestinationAccountID,
		transaction.Amount,
		transaction.Status,
		transaction.CreatedAt,
	).Scan(&id)

	if err != nil {
		return 0, err
	}

	return id, nil
}
