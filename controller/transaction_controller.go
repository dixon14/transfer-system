package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"transfer-system/models"
	"transfer-system/service"
)

// TransactionController handles HTTP requests for transaction operations
type TransactionController struct {
	transactionService *service.TransactionService
}

// NewTransactionController creates a new TransactionController instance
func NewTransactionController(transactionService *service.TransactionService) *TransactionController {
	return &TransactionController{
		transactionService: transactionService,
	}
}

// Transfer handles POST /transactions
func (c *TransactionController) Transfer(ctx *gin.Context) {
	var req models.TransferRequest

	// Bind and validate request
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error:   err.Error(),
			Message: "invalid transaction request",
		})
		return
	}

	// Perform transfer
	response, err := c.transactionService.Transfer(&req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error:   err.Error(),
			Message: "transaction failed",
		})
		return
	}

	ctx.JSON(http.StatusOK, response)
}
