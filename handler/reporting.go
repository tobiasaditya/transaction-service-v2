package handler

import (
	"net/http"
	"time"
	"transaction-service-v2/helper"
	"transaction-service-v2/reporting"
	"transaction-service-v2/user"
	"transaction-service-v2/util"

	"github.com/gin-gonic/gin"
)

type ReportingHandler struct {
	reportingService reporting.Service
}

func NewReportingHandler(reportingService reporting.Service) *ReportingHandler {
	return &ReportingHandler{reportingService: reportingService}
}

func (h ReportingHandler) CreateRecordDaily(c *gin.Context) {
	var input reporting.InputDailyReport
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.ErrorValidationResponse(err)

		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Invalid Input", http.StatusUnprocessableEntity, errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)

	newRecord, err := h.reportingService.AddDailyRecord(input, currentUser.ID.Hex())
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Failed to add record", http.StatusBadRequest, errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	format := reporting.FormatReporting(newRecord)
	response := helper.APIResponse("Success add record", http.StatusOK, format)
	c.JSON(http.StatusOK, response)
}

func (h ReportingHandler) GetReportingsUser(c *gin.Context) {
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
		timeNow := util.CTimeNow()
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
		inputEndDate = util.CTimeNow()
	}

	currentUser := c.MustGet("currentUser").(user.User)
	reportings, err := h.reportingService.GetRecords(currentUser.ID.Hex(), inputStartDate, inputEndDate)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Failed to get reporting", http.StatusBadRequest, errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	// format := transaction.FormatTransactions(transactions)
	format := reporting.FormatReportings(reportings)
	response := helper.APIResponse("Success get list reporting", http.StatusOK, format)
	c.JSON(http.StatusOK, response)
}
