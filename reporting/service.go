package reporting

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Service interface {
	AddDailyRecord(input InputDailyReport, userID string) (DailyReport, error)
	GetRecords(userID string, start time.Time, end time.Time) ([]DailyReport, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) AddDailyRecord(input InputDailyReport, userID string) (DailyReport, error) {
	dailyReport := DailyReport{
		ID:          primitive.NewObjectID(),
		BodyWeight:  input.BodyWeight,
		Counter:     input.Counter,
		RequestTime: time.Now(),
		UserID:      userID,
	}
	dailyReport, err := s.repository.CreateReport(dailyReport)
	if err != nil {
		return dailyReport, err
	}
	return dailyReport, nil
}

func (s *service) GetRecords(userID string, start time.Time, end time.Time) ([]DailyReport, error) {
	reports, err := s.repository.GetReportByUserID(userID, start, end)
	if err != nil {
		return reports, err
	}
	return reports, nil
}
