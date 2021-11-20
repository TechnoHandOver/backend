package parser

import (
	"reflect"
)

func GetOrDefault(value interface{}, defaultValue interface{}) interface{} {
	reflectValue := reflect.ValueOf(value)
	if reflectValue.Kind() != reflect.Ptr {
		panic("Function work only with pointers!")
	}

	if !reflectValue.IsNil() {
		return reflectValue.Elem().Interface()
	}
	return defaultValue
}
