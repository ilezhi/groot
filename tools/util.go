package tools

import (
	"strings"
	"reflect"
	"crypto/md5"
	"encoding/hex"
)

func StructToMap(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	data := make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		key := t.Field(i).Tag.Get("json")
		val := v.Field(i).Interface()
		data[key] = val
	}

	return data
}

func GetAvatar(email string) string {
	h := md5.New()
	h.Write([]byte(strings.ToLower(email)))
	return "//gravatar.com/avatar/" + hex.EncodeToString(h.Sum(nil)) + "?size=48"
}
