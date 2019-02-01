package config

import (
	"reflect"
)

type Config struct {
	Name string							`json:"name"`
	Version string					`json:"version"`
	Description string			`json:"desc"`
	Keywords string					`json:"keywords"`
	Admin []string					`json:"admin"`
	DB string								`json:"db"`
	Session_Secret string		`json:"ss"`
	Departments []string    `json:"departments"`
	Localhost string        `json:"localhost"`
}

var config = &Config{
	Name: "groot",
	Version: "0.0.1",
	Description: "实时问答社区",
	Keywords: "go, gorm, mysql, socket, iris, angular",
	Admin: []string{"dm@123.com", "skr@123.com"},
	DB: "mysql:ThisISmysql@tcp(db:3306)/groot?charset=utf8&parseTime=True&loc=Local",
	Departments: []string{"技术部", "人事部", "财务部", "其它"},
	Session_Secret: "racconandgroot",
	Localhost: "0.0.0.0:9000",
}

func Values() *Config {
	return config
}

func (c *Config) Get(key string) interface{} {
	v := reflect.ValueOf(c).Elem()

	for i := 0; i < v.NumField(); i++ {
		fieldInfo := v.Type().Field(i)
		tag := fieldInfo.Tag
		name := tag.Get("json")

		if name == key {
			item := v.Field(i)

			switch item.Kind() {
			case reflect.Slice, reflect.Array:
				return item.Interface()

			default:
				return item.String()
			}
		}
	}

	return ""
}
