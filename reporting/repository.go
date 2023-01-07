package reporting

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository interface {
	CreateReport(dailyReport DailyReport) (DailyReport, error)
	GetReportByUserID(userID string, start time.Time, end time.Time) ([]DailyReport, error)
}

type repository struct {
	db *mongo.Collection
}

func NewRepository(db *mongo.Collection) *repository {
	return &repository{db: db}
}

func (r *repository) CreateReport(dailyReport DailyReport) (DailyReport, error) {
	_, err := r.db.InsertOne(context.TODO(), dailyReport)
	if err != nil {
		return dailyReport, err
	}
	return dailyReport, nil
}

func (r *repository) GetReportByUserID(userID string, start time.Time, end time.Time) ([]DailyReport, error) {
	filter := bson.D{
		{Key: "userId", Value: userID},
		{Key: "requestTime", Value: bson.D{
			{Key: "$gte", Value: start},
			{Key: "$lte", Value: end},
		}},
	}
	findOptions := options.Find()
	findOptions.SetSort(bson.D{{Key: "requestTime", Value: -1}})
	// fmt.Println(filter)
	reports := []DailyReport{}
	cursor, err := r.db.Find(context.TODO(), filter, findOptions)
	if err != nil {
		return reports, err
	}
	err = cursor.All(context.TODO(), &reports)
	if err != nil {
		fmt.Println(err)
		return reports, err
	}
	return reports, nil
}
