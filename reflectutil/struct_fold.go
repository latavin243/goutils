package reflectutil

import (
	"fmt"
	"reflect"
)

// StructFoldSum folds all items and add to res
// ONLY works for struct with all demanded fields of number type
// put tag `fold:"skip"` on fields definition to skip fold (e.g. skip string fields)
// check unittest for example
func StructFoldSum(res interface{}, items ...interface{}) (err error) {
	for _, item := range items {
		err = structFoldAdd(res, item)
		if err != nil {
			return err
		}
	}
	return nil
}

func structFoldAdd(res interface{}, item interface{}) (err error) {
	if !SameType(res, item) {
		return fmt.Errorf("res and item should be the same type, res=%T, item=%T", res, item)
	}

	resValue := reflect.ValueOf(res)
	itemValue := reflect.ValueOf(item)
	if resValue.Kind() != reflect.Ptr || itemValue.Kind() != reflect.Ptr {
		return fmt.Errorf("res and item should be pointer")
	}

	resValue = resValue.Elem()
	itemValue = itemValue.Elem()
	if resValue.Kind() != reflect.Struct || itemValue.Kind() != reflect.Struct {
		return fmt.Errorf("res and item should be pointer to struct")
	}

Loop:
	for i := 0; i < resValue.NumField(); i++ {
		resField := resValue.Field(i)
		itemField := itemValue.Field(i)

		switch resValue.Type().Field(i).Tag.Get("fold") {
		case "skip", "first":
			continue Loop
		case "zero":
			resField.Set(reflect.Zero(resField.Type()))
			continue Loop
		default:
			// do nothing
		}

		if !isNumberValue(resField) || !isNumberValue(itemField) {
			return fmt.Errorf("res and item should be pointer to struct of numbers, resFieldKind=%+v, itemFieldKind=%+v", resField.Kind(), itemField.Kind())
		}

		addRes, err := numberValueAdd(resField, itemField)
		if err != nil {
			return err
		}
		resField.Set(addRes.Convert(resField.Type()))
	}
	return nil
}

func numberValueAdd(lhs, rhs reflect.Value) (reflect.Value, error) {
	if !SameType(lhs, rhs) {
		return reflect.Value{}, fmt.Errorf("inputs should be the same type, lhs=%T, rhs=%T", lhs, rhs)
	}
	if !isNumberValue(lhs) || !isNumberValue(rhs) {
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
