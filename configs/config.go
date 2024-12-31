package configs

import (
	"fmt"
	"gopkg.in/ini.v1"
)

var (
	config *ini.File
	err    error
)

func GetValue(section string, key string) string {
	sec, err := config.GetSection(section)
	if err != nil {
		fmt.Printf("Fail to get section: %v", err)
		return ""
	}
	getKey, err := sec.GetKey(key)
	if err != nil {
		fmt.Printf("Fail to get key: %v", err)
		return ""
	}
	return getKey.Value()
}
func GetConfig() *ini.File {
	return config
}
func InitConfig() {
	config, err = ini.Load("configs/dev/config.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		return
	}
}
