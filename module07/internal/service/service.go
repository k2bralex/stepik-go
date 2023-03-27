package service

import (
	"reflect"
)

func StructToMap(item interface{}) map[string]interface{} {

	out := make(map[string]interface{})

	v := reflect.ValueOf(item)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct { // Non-structural return error
		return nil
	}

	typ := v.Type()

	for i := 0; i < v.NumField(); i++ {
		fi := typ.Field(i)
		out[fi.Name] = v.Field(i).Interface()
	}

	return out
}
