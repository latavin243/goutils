package number

import (
	"strconv"

	"golang.org/x/exp/constraints"
)

type Num interface {
	constraints.Integer | constraints.Float
}

// number to string

func IntToString[T constraints.Integer](n T) string {
	return strconv.FormatInt(int64(n), 10)
}

func FloatToString[T constraints.Float](n T, precision int) string {
	return strconv.FormatFloat(float64(n), 'f', precision, 64)
}

// string to number

func StringToInt[T constraints.Integer](s string) (T, error) {
	res, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0, err
	}
	return T(res), nil
}

func StringToFloat[T constraints.Float](s string) (T, error) {
	n, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0, err
	}
	return T(n), nil
}
