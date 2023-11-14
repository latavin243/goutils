package reflectutil

import (
	"fmt"
	"reflect"
	"strconv"

	"golang.org/x/exp/constraints"
)

// DirectPrintable includes types which can be directly converted to string
// float is excluded because of precision loss
type DirectPrintable interface {
	constraints.Integer | ~string | ~bool
}

func ToString(input interface{}) (string, error) {
	switch reflect.TypeOf(input).Kind() {
	case reflect.String:
		return input.(string), nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(reflect.ValueOf(input).Int(), 10), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return strconv.FormatUint(reflect.ValueOf(input).Uint(), 10), nil
	case reflect.Bool:
		return strconv.FormatBool(input.(bool)), nil
	default:
		return "", fmt.Errorf("invalid type, kind=%s", reflect.TypeOf(input).Kind().String())
	}
}

func ToStrings(inputs []interface{}) ([]string, error) {
	strs := make([]string, 0, len(inputs))
	for _, input := range inputs {
		str, err := ToString(input)
		if err != nil {
			return strs, err
		}
		strs = append(strs, str)
	}
	return strs, nil
}

func DirectToString[T DirectPrintable](input T) string {
	return fmt.Sprintf("%v", input)
}

func DirectToStrings[T DirectPrintable](inputs []T) []string {
	strs := make([]string, 0, len(inputs))
	for _, input := range inputs {
		strs = append(strs, DirectToString(input))
	}
	return strs
}
