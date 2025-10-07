package service

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strconv"
	"time"

	"transfer-system/dao"
	"transfer-system/models"
)

// AccountService handles business logic for accounts
type AccountService struct {
	accountDAO *dao.AccountDAO
}

// NewAccountService creates a new AccountService instance
func NewAccountService(db *sql.DB) *AccountService {
	return &AccountService{
		accountDAO: dao.NewAccountDAO(db),
	}
}

// CreateAccount creates a new account with validation
func (s *AccountService) CreateAccount(req *models.CreateAccountRequest) error {
	accountInitialBalance, err := s.validateInitialBalance(req)
	if err != nil {
		return err
	}
	// Check if account already exists
	existingAccount, _ := s.accountDAO.GetAccountByID(req.AccountID)
	if existingAccount != nil {
		log.Printf("Account with ID %d already exists.", req.AccountID)
		return fmt.Errorf("account with ID %d already exists", req.AccountID)
	}

	// Create account
	account := &models.Account{
		AccountID: req.AccountID,
		Balance:   accountInitialBalance,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	return s.accountDAO.CreateAccount(account)
}

// GetAccount retrieves an account by ID
func (s *AccountService) GetAccount(accountID uint64) (*models.AccountResponse, error) {
	account, err := s.accountDAO.GetAccountByID(accountID)
	if err != nil {
		return nil, err
	}

	balanceStr := strconv.FormatFloat(account.Balance, 'f', 5, 64)
	return &models.AccountResponse{
		AccountID: account.AccountID,
		Balance:   balanceStr,
	}, nil
}

func (s *AccountService) validateInitialBalance(req *models.CreateAccountRequest) (float64, error) {
	// Validate initial balance
	accountInitialBalance, err := strconv.ParseFloat(req.InitialBalance, 64)
	if err != nil {
		log.Print("Failed to parse initial balance.")
		return 0, fmt.Errorf("invalid initial_balance, %w", err)
	}

	if accountInitialBalance < 0 {
		log.Print("Negative account initial balance.")
		return 0, errors.New("initial_balance cannot be less than 0")
	}
	return accountInitialBalance, nil
}
