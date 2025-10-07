package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"transfer-system/models"
	"transfer-system/service"
)

// AccountController handles HTTP requests for account operations
type AccountController struct {
	accountService *service.AccountService
}

// NewAccountController creates a new AccountController instance
func NewAccountController(accountService *service.AccountService) *AccountController {
	return &AccountController{
		accountService: accountService,
	}
}

// CreateAccount handles accounts creation
func (c *AccountController) CreateAccount(ctx *gin.Context) {
	var req models.CreateAccountRequest

	// Bind and validate request
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error:   err.Error(),
			Message: "invalid request",
		})
		return
	}

	// Create account
	if err := c.accountService.CreateAccount(&req); err != nil {
		ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error:   err.Error(),
			Message: "account creation failed",
		})
		return
	}
	ctx.Status(http.StatusOK)
}

// GetAccount handles GET /accounts/:account_id
func (c *AccountController) GetAccount(ctx *gin.Context) {
	accountIDStr := ctx.Param("account_id")

	if accountIDStr == "" {
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error:   "account_id is required",
			Message: "invalid request",
		})
		return
	}

	accountID, err := strconv.ParseUint(accountIDStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error:   "invalid account_id",
			Message: "invalid request",
		})
		return
	}
	// Get account
	account, err := c.accountService.GetAccount(accountID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error:   err.Error(),
			Message: "account not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, account)
}
