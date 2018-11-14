package config

import (
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"reflect"
)

type Config struct {
	Name string							`json:"name"`
	Version string					`json:"version"`
	Description string			`json:"desc"`
	Keywords string					`json:"keywords"`
	Admin string						`json:"admin"`
	DB string								`json:"db"`
	Session_Secret string		`json:"ss"`
	Port int								`json:"port"`
	Init bool								`json:"init"`
}

var config = new(Config)

func Values() *Config {
	if !config.Init {
		buf, _ := ioutil.ReadFile("./config.yaml")
		yaml.Unmarshal(buf, config)
		config.Init = true
	}

	return config
}

func (c *Config) Get(key string) interface{} {
	v := reflect.ValueOf(c).Elem()

	for i := 0; i < v.NumField(); i++ {
		fieldInfo := v.Type().Field(i)
		tag := fieldInfo.Tag
		name := tag.Get("json")
		if name == key {
			return v.Field(i).String()
		}
	}

	return ""
}
