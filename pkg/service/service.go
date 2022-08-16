package service

import (
	"github.com/maxence-charriere/go-app/v9/pkg/errors"
)

var (
	memory = map[string]string{
		"config":       `name: "config"`,
		"/dev/config":  `name: "dev"`,
		"/prod/config": `name: "prod"`,
	}
)

func ListKeys() []string {
	keys := make([]string, len(memory))
	for key := range memory {
		keys = append(keys, key)
	}
	return keys
}

func Get(key string) (string, error) {
	value, ok := memory[key]
	if !ok {
		return "", errors.New("not found")
	}

	return value, nil
}

func Put(key, value string) error {
	memory[key] = value
	return nil
}
