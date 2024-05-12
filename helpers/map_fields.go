package helpers

import (
	"reflect"
)

func MapFields(src interface{}, dest interface{}) {
	srcVal := reflect.ValueOf(src)
	destVal := reflect.ValueOf(dest).Elem()

	for i := 0; i < srcVal.NumField(); i++ {
		srcField := srcVal.Field(i)
		srcFieldName := srcVal.Type().Field(i).Name
		destField := destVal.FieldByName(srcFieldName)

		if destField.IsValid() && destField.CanSet() && srcField.Kind() == destField.Kind() {
			destField.Set(srcField)
		}
	}
}
