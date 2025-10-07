package service

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strconv"
	"time"

	"transfer-system/dao"
	"transfer-system/enums"
	"transfer-system/models"
)

// TransactionService handles business logic for transactions
type TransactionService struct {
	accountDAO     *dao.AccountDAO
	transactionDAO *dao.TransactionDAO
	db             *sql.DB
}

// NewTransactionService creates a new TransactionService instance
func NewTransactionService(db *sql.DB) *TransactionService {
	return &TransactionService{
		accountDAO:     dao.NewAccountDAO(db),
		transactionDAO: dao.NewTransactionDAO(db),
		db:             db,
	}
}

// Transfer performs a fund transfer between two accounts with ACID guarantees
func (s *TransactionService) Transfer(req *models.TransferRequest) (*models.TransferResponse, error) {
	// Validate request
	amount, err := s.validateTransferRequest(req)
	if err != nil {
		return nil, err
	}

	// Start database transaction
	tx, err := s.db.Begin()
	if err != nil {
		return nil, fmt.Errorf("failed to start transaction: %w", err)
	}
	defer tx.Rollback() // Rollback if not committed

	// Retrieve source account
	sourceAccount, err := s.accountDAO.GetAccountByIDForUpdate(tx, req.SourceAccountID)
	if err != nil {
		
		return nil, fmt.Errorf("source account error: %w", err)
	}

	// Retrieve destination account
	destinationAccount, err := s.accountDAO.GetAccountByIDForUpdate(tx, req.DestinationAccountID)
	if err != nil {
		return nil, fmt.Errorf("destination account error: %w", err)
	}

	// Check if source account has sufficient balance
	if sourceAccount.Balance < amount {
		return nil, errors.New("insufficient balance in source account")
	}

	// Calculate new balances
	newSourceBalance := sourceAccount.Balance - amount
	newDestinationBalance := destinationAccount.Balance + amount

	// Update source account balance
	if err := s.accountDAO.UpdateBalance(tx, req.SourceAccountID, newSourceBalance); err != nil {
		return nil, fmt.Errorf("failed to update source account: %w", err)
	}

	// Update destination account balance
	if err := s.accountDAO.UpdateBalance(tx, req.DestinationAccountID, newDestinationBalance); err != nil {
		return nil, fmt.Errorf("failed to update destination account: %w", err)
	}

	// Create transaction record
	transaction := &models.Transaction{
		SourceAccountID:      req.SourceAccountID,
		DestinationAccountID: req.DestinationAccountID,
		Amount:               amount,
		Status:               enums.Success.String(),
		CreatedAt:            time.Now(),
	}

	transactionID, err := s.transactionDAO.CreateTransaction(tx, transaction)
	if err != nil {
		log.Print("Failed to create transaction record in DB.")
		return nil, fmt.Errorf("failed to create transaction record: %w", err)
	}

	// Commit transaction
	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %w", err)
	}

	log.Printf("Successfully transfer %.2f dollars from account ID %d to %d", amount, sourceAccount.AccountID, destinationAccount.AccountID)
	return &models.TransferResponse{
		TransactionID:        transactionID,
		SourceAccountID:      req.SourceAccountID,
		DestinationAccountID: req.DestinationAccountID,
		Amount:               req.Amount,
		Status:               enums.Success.String(),
	}, nil
}

// validateTransferRequest validates the transfer request and returns the parsed amount
func (s *TransactionService) validateTransferRequest(req *models.TransferRequest) (float64, error) {
	if req.SourceAccountID == req.DestinationAccountID {
		return 0, errors.New("source and destination accounts cannot be the same")
	}

	amount, err := strconv.ParseFloat(req.Amount, 64)
	if amount <= 0 || err != nil {
		return 0, errors.New("invalid amount")
	}
	return amount, nil
}
