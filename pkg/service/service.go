package service

import (
	"fmt"
)

func ListKeys() []string {
	return []string{"config", "/dev/config", "/prod/config"}
}

func Get(key string) (string, error) {
	switch key {

	case "/dev/config":
		return `name: "dev"`, nil
	case "/prod/config":
		return `name: "prod"`, nil
	case "config":
		return `name: "config"`, nil
	default:
		return ``, nil
	}
}

func Put(key, value string) error {
	fmt.Println(key, value)
	return nil
}
