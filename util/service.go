package util

import (
	"os"
	"time"
)

func CTimeNow() time.Time {
	timezone := os.Getenv("TZ")
	loc, _ := time.LoadLocation(timezone)
	return time.Now().In(loc)

}
