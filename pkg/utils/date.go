package utils

import (
	"strconv"
	"strings"
	"time"

	"github.com/snabb/isoweek"
)

const (
	layoutISO = "2006-01-02"
	layoutUS  = "January 2, 2006"
)

func DateStringToInt(date string) uint64 {
	removedDashStr := strings.ReplaceAll(date, "-", "")

	i, err := strconv.ParseUint(removedDashStr, 10, 64)
	if err != nil {
		return 0
	}

	return i
}

func GetWeekFromDateString(date string) (weekNumber int, startDate string) {
	t, _ := time.Parse(layoutISO, date)
	wyear, week := isoweek.FromDate(t.Year(), t.Month(), t.Day())

	start := isoweek.StartTime(wyear, week, time.UTC)

	startDate = start.Format(layoutISO)
	weekNumber = week
	return
}

func GetDateString(t time.Time) string {

	return t.Format(layoutISO)
}
