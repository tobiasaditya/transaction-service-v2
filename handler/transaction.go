package handler

import (
	"net/http"
	"time"
	"transaction-service-v2/helper"
	"transaction-service-v2/transaction"
	"transaction-service-v2/user"

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
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Failed to add transaction", http.StatusBadRequest, errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	format := transaction.FormatTransaction(newTrx)
	response := helper.APIResponse("Success adding transaction", http.StatusOK, format)
	c.JSON(http.StatusOK, response)

}

func (h transactionHandler) GetTransactionsUser(c *gin.Context) {
	startDate := c.Query("startDate")
	endDate := c.Query("endDate")
	var inputStartDate time.Time
	var inputEndDate time.Time

	if len(startDate) > 0 {
		resultParse, err := time.Parse("2006-01-02", startDate)
		if err != nil {
			errors := helper.ErrorValidationResponse(err)

			errorMessage := gin.H{"errors": errors}

			response := helper.APIResponse("Invalid Input", http.StatusUnprocessableEntity, errorMessage)
			c.JSON(http.StatusUnprocessableEntity, response)
			return
		}
		inputStartDate = resultParse
	} else {
		//If no input, make default date as start of month
		timeNow := time.Now()
		inputStartDate = time.Date(timeNow.Year(), timeNow.Month(), 1, 0, 0, 0, 0, timeNow.Location())

	}

	if len(endDate) > 0 {
		resultParse, err := time.Parse("2006-01-02", endDate)
		if err != nil {
			errors := helper.ErrorValidationResponse(err)

			errorMessage := gin.H{"errors": errors}

			response := helper.APIResponse("Invalid Input", http.StatusUnprocessableEntity, errorMessage)
			c.JSON(http.StatusUnprocessableEntity, response)
			return
		}
		inputEndDate = resultParse
	} else {
		//If no input, make default date from today
		inputEndDate = time.Now()
	}

	currentUser := c.MustGet("currentUser").(user.User)
	transactions, err := h.transactionService.GetTransactions(currentUser.ID.Hex(), inputStartDate, inputEndDate)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Failed to get transactions", http.StatusBadRequest, errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	format := transaction.FormatTransactions(transactions)
	response := helper.APIResponse("Success get list transactions", http.StatusOK, format)
	c.JSON(http.StatusOK, response)
}
