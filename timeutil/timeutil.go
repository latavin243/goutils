package timeutil

import (
	"time"
)

// RoundToDay rounds time to day 0:00
func RoundToDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

// RoundToWeek rounds time to Monday 0:00
func RoundToWeek(t time.Time) time.Time {
	roundDayT := RoundToDay(t)
	weekday := int(roundDayT.Weekday())
	if weekday == int(time.Sunday) {
		weekday = 7
	}
	return roundDayT.AddDate(0, 0, int(time.Monday)-weekday)
}

// RoundToMonth rounds time to month 1st 0:00
func RoundToMonth(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location())
}

// MonthDays returns day number of month
func MonthDays(refT time.Time) uint32 {
	roundT := time.Date(refT.Year(), refT.Month(), 1, 0, 0, 0, 0, refT.Location())
	return uint32(roundT.AddDate(0, 1, 0).Sub(roundT).Hours() / 24)
}

// TsToTime converts timestamp (second) to time
func TsToTime(ts uint64, loc *time.Location) time.Time {
	return time.Unix(int64(ts), 0).In(loc)
}
