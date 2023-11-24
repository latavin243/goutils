package timeutil

import (
	"time"
)

// RoundToDay rounds time to day 0:00
func RoundToDay(refT time.Time) time.Time {
	return time.Date(refT.Year(), refT.Month(), refT.Day(), 0, 0, 0, 0, refT.Location())
}

// RoundToMonday rounds time to Monday 0:00
func RoundToMonday(t time.Time) time.Time {
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

// RoundToWeekFirst rounds dt to Monday dt
func RoundToWeekFirst(dt uint32) uint32 {
	dtTime := DtToTime(dt, time.UTC)
	gapDayOfMonday := int(dtTime.Weekday())
	if gapDayOfMonday == 0 {
		gapDayOfMonday = 7
	}

	daysUntilMonday := gapDayOfMonday - 1
	monday := dtTime.AddDate(0, 0, daysUntilMonday*-1)
	return TimeToDt(time.Date(monday.Year(), monday.Month(), monday.Day(), 0, 0, 0, 0, time.UTC))
}

// RoundToMonthBeginDt rounds dt to month 1st dt
func RoundToMonthBeginDt(dt uint32) uint32 {
	dtTime := DtToTime(dt, time.UTC)
	currentYear, currentMonth, _ := dtTime.Date()
	monthBegin := time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, time.UTC)
	return TimeToDt(monthBegin)
}

// DtToTime converts dt (yyyyMMdd) uint32 to time.Time
func DtToTime(dt uint32, loc *time.Location) time.Time {
	return time.Date(int(dt/10000), time.Month((dt/100)%100), int(dt%100), 0, 0, 0, 0, loc)
}

// TimeToDt converts time.Time to dt (yyyyMMdd) uint32
func TimeToDt(t time.Time) uint32 {
	return uint32(t.Year()*10000 + int(t.Month())*100 + t.Day())
}

// DtsBetween returns all dts between startDt and endDt
func DtsBetween(startDt, endDt uint32, includeEnd bool) []uint32 {
	startT := DtToTime(startDt, time.UTC)
	endT := DtToTime(endDt, time.UTC)
	if includeEnd {
		endT = endT.AddDate(0, 0, 1)
	}
	dts := make([]uint32, 0)
	for t := startT; t.Before(endT); t = t.AddDate(0, 0, 1) {
		dts = append(dts, TimeToDt(t))
	}
	return dts
}

// MonthDays returns day number of month
func MonthDays(refT time.Time) uint32 {
	roundT := time.Date(refT.Year(), refT.Month(), 1, 0, 0, 0, 0, refT.Location())
	return uint32(roundT.AddDate(0, 1, 0).Sub(roundT).Hours() / 24)
}

// GetDtGapDays returns gap days between startDt and endDt
func GetDtGapDays(startDt, endDt uint32) (gapDays uint32) {
	startT := DtToTime(startDt, time.UTC)
	endT := DtToTime(endDt, time.UTC)
	return uint32(endT.Sub(startT).Hours() / 24)
}

// TsToDt converts timestamp (second) to dt (yyyyMMdd)
func TsToDt(ts uint64, loc *time.Location) uint32 {
	return TimeToDt(TsToTime(ts, loc))
}

// TsToTime converts timestamp (second) to time
func TsToTime(ts uint64, loc *time.Location) time.Time {
	return time.Unix(int64(ts), 0).In(loc)
}

// DtToTs converts timestamp (second) to time
func DtToTs(dt uint32, loc *time.Location) uint64 {
	return uint64(DtToTime(dt, loc).Unix())
}
