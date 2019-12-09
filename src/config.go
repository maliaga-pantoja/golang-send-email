package src

import (
	"errors"
	"os"
)

func GetEnv (key string) (string) {
	envKey := os.Getenv(key)

	if envKey == "" {
		envKey = defaultValue
	}
	return envKey
}
