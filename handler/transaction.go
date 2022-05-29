package handler

import (
	"net/http"
	"transaction-service-v2/helper"
	"transaction-service-v2/transaction"

	"github.com/gin-gonic/gin"
)

type transactionHandler struct {
	transactionService transaction.Service
}

func NewTransactionHandler(transactionService transaction.Service) *transactionHandler {
	return &transactionHandler{transactionService: transactionService}
}

func (h transactionHandler) CreateTransaction(c *gin.Context) {
	var input transaction.InputTransaction
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.ErrorValidationResponse(err)

		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Invalid Input", http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newTrx, err := h.transactionService.CreateTransaction(input)
	if err != nil {
		errorMessage := gin.H{"errors": err}

		response := helper.APIResponse("Failed to add transaction", http.StatusBadRequest, errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	format := transaction.FormatTransaction(newTrx)
	response := helper.APIResponse("Success adding transaction", http.StatusOK, format)
	c.JSON(http.StatusOK, response)

}

func (h transactionHandler) GetTransactionsUser(c *gin.Context) {
	userID := 1
	transactions, err := h.transactionService.GetTransactions(userID)
	if err != nil {
		errorMessage := gin.H{"errors": err}

		response := helper.APIResponse("Failed to get transactions", http.StatusBadRequest, errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	format := transaction.FormatTransactions(transactions)
	response := helper.APIResponse("Success get list transactions", http.StatusOK, format)
	c.JSON(http.StatusOK, response)
}
