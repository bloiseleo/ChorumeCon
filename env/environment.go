package env

import (
	"os"
)

func Env(key string, defaultValue string) string {
	value, present := os.LookupEnv(key)
	if present {
		return value
	}
	return defaultValue
}
