package reporting

type InputDailyReport struct {
	BodyWeight float64 `bson:"bodyWeight"`
	Counter    int     `bson:"counter"`
}
