package util

import (
	"time"
)

func CTimeNow() time.Time {
	// timezone := os.Getenv("TZ")
	timezone := "Asia/Jakarta"
	loc, _ := time.LoadLocation(timezone)
	return time.Now().In(loc)

}

func FormatTime(input time.Time) string {
	// timezone := os.Getenv("TZ")
	timezone := "Asia/Jakarta"
	loc, _ := time.LoadLocation(timezone)
	return input.In(loc).Format("2006-01-02 15:04:05")

}
