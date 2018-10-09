package tools

import (
	"strings"
	"reflect"
)

func StructToMap(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	data := make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		key := strings.ToLower(t.Field(i).Name)
		val := v.Field(i).Interface()
		data[key] = val
	}

	return data
}
