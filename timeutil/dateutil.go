package timeutil

import "time"

type dateT = uint32

var defaultLoc = time.UTC // default loc for dt calculation

// === Conversion Functions ===

// DtToTime converts dt (yyyyMMdd dateT) to time.Time
func DtToTime(dt dateT, loc *time.Location) time.Time {
	return time.Date(int(dt/10000), time.Month((dt/100)%100), int(dt%100), 0, 0, 0, 0, loc)
}

// TimeToDt converts time.Time to dt (yyyyMMdd dateT)
func TimeToDt(t time.Time) dateT {
	return dateT(t.Year()*10000 + int(t.Month())*100 + t.Day())
}

// TsToDt converts timestamp (second) to dt (yyyyMMdd)
func TsToDt(ts uint64, loc *time.Location) dateT {
	return TimeToDt(TsToTime(ts, loc))
}

// DtToTs converts timestamp (second) to time
func DtToTs(dt dateT, loc *time.Location) uint64 {
	return uint64(DtToTime(dt, loc).Unix())
}

// === Round Functions ===

// RoundToMonthFirstDt rounds dt to month 1st dt
func RoundToMonthFirstDt(dt dateT) dateT {
	dtTime := DtToTime(dt, defaultLoc)
	currentYear, currentMonth, _ := dtTime.Date()
	monthBegin := time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, defaultLoc)
	return TimeToDt(monthBegin)
}

// RoundToWeekFirstDt rounds dt to Monday dt
func RoundToWeekFirstDt(dt dateT) dateT {
	dtTime := DtToTime(dt, defaultLoc)
	gapDayOfMonday := int(dtTime.Weekday())
	if gapDayOfMonday == 0 {
		gapDayOfMonday = 7
	}

	daysUntilMonday := gapDayOfMonday - 1
	monday := dtTime.AddDate(0, 0, daysUntilMonday*-1)
	return TimeToDt(time.Date(monday.Year(), monday.Month(), monday.Day(), 0, 0, 0, 0, defaultLoc))
}

// === Other Functions ===

// DtsBetween generates all dts between startDt and endDt
func DtsBetween(startDt, endDt dateT, includeEnd bool) []dateT {
	startT := DtToTime(startDt, defaultLoc)
	endT := DtToTime(endDt, defaultLoc)
	if includeEnd {
		endT = endT.AddDate(0, 0, 1)
	}
	dts := make([]dateT, 0)
	for t := startT; t.Before(endT); t = t.AddDate(0, 0, 1) {
		dts = append(dts, TimeToDt(t))
	}
	return dts
}

// DtGapDays calculates gap days between startDt and endDt
func DtGapDays(startDt, endDt dateT) (gapDays uint32) {
	startT := DtToTime(startDt, defaultLoc)
	endT := DtToTime(endDt, defaultLoc)
	return uint32(endT.Sub(startT).Hours() / 24)
}
