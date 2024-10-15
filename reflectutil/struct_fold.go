package reflectutil

import (
	"fmt"
	"reflect"
)

const (
	// common fold process
	foldTagSkip  = "skip" // default skip
	foldTagFirst = "first"

	// fold process for number
	foldTagNumZero = "num/zero"
	foldTagNumSum  = "num/sum"

	// fold process for string
	foldTagStrEmpty    = "str/empty"
	foldTagStrCatenate = "str/cat"
	foldTagStrJoin     = "str/join" // e.g. "str/join/-" to join with "-"
)

// StructFold folds all items
// check unittest for example
func StructFold(res interface{}, items ...interface{}) (err error) {
	if !IsStructPtr(res) {
		return fmt.Errorf("res should be pointer to struct")
	}
	// TODO check res tag
	if err = checkStructFoldTag(res); err != nil {
		return err
	}

	for _, item := range items {
		if !IsStructPtr(item) {
			return fmt.Errorf("res and item should be pointer to struct")
		}
		if !SameType(res, item) {
			return fmt.Errorf("res and item should be the same type, res=%T, item=%T", res, item)
		}

		err = structFold(res, item)
		if err != nil {
			return err
		}
	}
	return nil
}

func checkStructFoldTag(input interface{}) error {
	v := reflect.ValueOf(input).Elem()
	for i := 0; i < v.NumField(); i++ {
		field := v.Type().Field(i)
		tag := field.Tag.Get("fold")
		switch tag {
		case foldTagNumZero, foldTagNumSum:
			if !isNumberKind(field.Type.Kind()) {
				return fmt.Errorf("field %s should be number type", field.Type.Name())
			}
		case foldTagStrEmpty, foldTagStrCatenate, foldTagStrJoin:
			if !isStringKind(field.Type.Kind()) {
				return fmt.Errorf("field %s should be string type", field.Type.Name())
			}
		default:
			// do nothing
		}
	}
	return nil
}

func structFold(res interface{}, item interface{}) (err error) {
	resValue := reflect.ValueOf(res).Elem()
	itemValue := reflect.ValueOf(item).Elem()
Loop:
	for i := 0; i < resValue.NumField(); i++ {
		resField := resValue.Field(i)
		itemField := itemValue.Field(i)

		switch resValue.Type().Field(i).Tag.Get("fold") {
		case foldTagSkip, foldTagFirst:
			continue Loop
		case foldTagNumZero:
			resField.Set(reflect.Zero(resField.Type()))
			continue Loop
		case foldTagNumSum:
			sumRes, err := foldNumSum(resField, itemField)
			if err != nil {
				return err
			}
			resField.Set(sumRes.Convert(resField.Type()))
		default:
			continue Loop
		}
	}
	return nil
}

func foldNumSum(lhs, rhs reflect.Value) (reflect.Value, error) {
	if !SameType(lhs, rhs) {
		return reflect.Value{}, fmt.Errorf("inputs should be the same type, lhs=%T, rhs=%T", lhs, rhs)
	}
	if !isNumberKind(lhs.Kind()) || !isNumberKind(rhs.Kind()) {
		return reflect.Value{}, fmt.Errorf("inputs should be numbers, lhs=%T, rhs=%T", lhs, rhs)
	}

	switch lhs.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return reflect.ValueOf(lhs.Int() + rhs.Int()), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return reflect.ValueOf(lhs.Uint() + rhs.Uint()), nil
	case reflect.Float32, reflect.Float64:
		return reflect.ValueOf(lhs.Float() + rhs.Float()), nil
	default:
		return reflect.Value{}, fmt.Errorf("inputs should be number")
	}
}
