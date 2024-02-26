package reflectutil

import "reflect"

func IsString(input interface{}) bool {
	return reflect.TypeOf(input).Kind() == reflect.String
}

func IsPtr(input interface{}) bool {
	return reflect.TypeOf(input).Kind() == reflect.Ptr
}

func IsNil(input interface{}) bool {
	ret := input == nil
	if !ret {
		vi := reflect.ValueOf(input)
		switch vi.Kind() {
		case reflect.Slice, reflect.Map, reflect.Chan, reflect.Interface, reflect.Func, reflect.Ptr:
			return vi.IsNil()
		default:
			// do nothing
		}
	}
	return ret
}

func IsNumber(input interface{}) bool {
	return isNumberKind(reflect.TypeOf(input).Kind())
}

func isNumberKind(input reflect.Kind) bool {
	switch input {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Float32, reflect.Float64:
		return true
	default:
		return false
	}
}

func isStringKind(input reflect.Kind) bool {
	return input == reflect.String
}

func SameType(ref interface{}, others ...interface{}) bool {
	for _, other := range others {
		if reflect.TypeOf(ref) != reflect.TypeOf(other) {
			return false
		}
	}
	return true
}

func IsPtrToStruct(input interface{}) bool {
	v := reflect.ValueOf(input)
	return v.Kind() == reflect.Ptr && v.Elem().Kind() == reflect.Struct
}
