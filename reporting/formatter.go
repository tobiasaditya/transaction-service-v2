package reporting

import (
	"strconv"
	"transaction-service-v2/util"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FormatterReporting struct {
	ID          primitive.ObjectID `json:"id"`
	BodyWeight  string             `json:"bodyWeight"`
	Counter     string             `json:"counter"`
	RequestTime string             `json:"requestTime"`
}

func FormatReporting(dailyReport DailyReport) FormatterReporting {
	format := FormatterReporting{
		ID:          dailyReport.ID,
		BodyWeight:  strconv.FormatFloat(dailyReport.BodyWeight, 'f', 2, 64),
		Counter:     strconv.Itoa(dailyReport.Counter),
		RequestTime: util.FormatTime(dailyReport.RequestTime),
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
