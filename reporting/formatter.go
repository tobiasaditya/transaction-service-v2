package reporting

import (
	"strconv"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FormatterReporting struct {
	ID          primitive.ObjectID `json:"id"`
	BodyWeight  string             `json:"trxType"`
	Counter     string             `json:"amount"`
	RequestTime string             `json:"requestTime"`
}

func FormatReporting(dailyReport DailyReport) FormatterReporting {
	format := FormatterReporting{
		ID:         dailyReport.ID,
		BodyWeight: strconv.FormatFloat(dailyReport.BodyWeight, 'f', 2, 64),
		Counter:    strconv.Itoa(dailyReport.Counter),
	}

	return format
}

func FormatReportings(dailyReports []DailyReport) []FormatterReporting {
	formatRepostings := []FormatterReporting{}

	for _, report := range dailyReports {
		formatRepostings = append(formatRepostings, FormatReporting(report))
	}
	return formatRepostings
}
