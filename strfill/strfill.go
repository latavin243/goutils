package strfill

import (
	"bytes"
	"errors"
	"html/template"
	"strings"
	"time"

	"golang.org/x/exp/constraints"
)

var funcMap = template.FuncMap{
	"yyyyMM": yyyyMM,
	"join":   strings.Join,
}

// a function StrFill("my_table_{yyyyMM(.Date)}", Record{20060102}) should return "my_table_200601"

func StrFill(tmpl string, ref interface{}) (string, error) {
	buf := &bytes.Buffer{}
	t, err := template.New("").Funcs(funcMap).Parse(tmpl)
	if err != nil {
		return "", err
	}
	err = t.Execute(buf, ref)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}

func yyyyMM(rawDt interface{}) string {
	switch rawDt := rawDt.(type) {
	case int:
		dt, _ := parseDt(rawDt)
		return dt.Format("200601")
	case int8:
		dt, _ := parseDt(rawDt)
		return dt.Format("200601")
	case int16:
		dt, _ := parseDt(rawDt)
		return dt.Format("200601")
	case int32:
		dt, _ := parseDt(rawDt)
		return dt.Format("200601")
	case int64:
		dt, _ := parseDt(rawDt)
		return dt.Format("200601")
	case uint:
		dt, _ := parseDt(rawDt)
		return dt.Format("200601")
	case uint8:
		dt, _ := parseDt(rawDt)
		return dt.Format("200601")
	case uint16:
		dt, _ := parseDt(rawDt)
		return dt.Format("200601")
	case uint32:
		dt, _ := parseDt(rawDt)
		return dt.Format("200601")
	case uint64:
		dt, _ := parseDt(rawDt)
		return dt.Format("200601")
	default:
		return ""
	}
}

func parseDt[T constraints.Integer](rawDt T) (time.Time, error) {
	dt := int(rawDt)
	if dt < 10000000 || dt > 99999999 {
		return time.Now(), errors.New("invalid date")
	}
	return time.Date(dt/10000, time.Month((dt/100)%100), dt%100, 0, 0, 0, 0, time.UTC), nil
}
